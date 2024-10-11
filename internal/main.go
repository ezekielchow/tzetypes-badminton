package main

import (
	"common"
	commonmiddlewares "common/middlewares"
	"common/models"
	"common/oapiprivate"
	"common/oapipublic"
	"common/utils"
	"context"
	"database/sql"
	"net/http"
	"os"
	playerservice "players/service"
	playerstore "players/store"
	playerstoregenerated "players/store/generated"
	sessionstore "sessions/store"
	sessionstoregenerated "sessions/store/generated"
	"strings"
	"time"
	clubs "tzetypes-badminton/clubs/store"
	clubstoregenerated "tzetypes-badminton/clubs/store/generated"
	userservice "users/service"
	userstore "users/store"
	userstoregenerated "users/store/generated"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
	"github.com/sirupsen/logrus"
)

func BearerTokenAuth(sessionStore sessionstore.SessionRepository, userStore userstore.UserRepository) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
				return
			}

			tokenParts := strings.Split(authHeader, " ")
			if len(tokenParts) != 2 || strings.ToLower(tokenParts[0]) != "bearer" {
				http.Error(w, "Invalid Authorization header format", http.StatusUnauthorized)
				return
			}

			token := tokenParts[1]

			session, err := checkToken(r.Context(), sessionStore, token)
			if err != nil || session.ID == "" || session.SessionTokenExpiresAt.Before(time.Now()) {
				http.Error(w, "Invalid token", http.StatusUnauthorized)
				return
			}

			user, err := userStore.FindUserWithID(r.Context(), session.UserID)
			if err != nil || user.ID == "" {
				http.Error(w, "Invalid user", http.StatusUnauthorized)
				return
			}

			// Proceed with the request
			ctx := context.WithValue(r.Context(), common.ContextUser, user)
			ctx = context.WithValue(ctx, common.ContextSession, session)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func checkToken(ctx context.Context, sessionStore sessionstore.SessionRepository, token string) (models.Session, error) {
	session, err := sessionStore.FindSessionWithSessionID(ctx, token)
	if err != nil && !strings.Contains(sql.ErrNoRows.Error(), err.Error()) {
		return models.Session{}, err
	}

	return session, err
}

func getPrivateRouter(conn *pgx.Conn) *chi.Mux {

	apiRoute := commonmiddlewares.NewRouter()
	apiRoute.Use(
		BearerTokenAuth(
			&sessionstore.SessionPostgres{
				Queries: sessionstoregenerated.New(conn),
			},
			&userstore.UserPostgres{
				Queries: userstoregenerated.New(conn),
			}))

	return apiRoute
}

func main() {
	ctx := context.Background()

	dbURI := "postgresql://" + os.Getenv("POSTGRES_USER") + ":" + os.Getenv("POSTGRES_PASSWORD") + "@" + os.Getenv("POSTGRES_HOST") + "/" + os.Getenv("POSTGRES_DB") + "?sslmode=disable"

	err := utils.RunMigrations(dbURI)
	if err != nil {
		panic(err)
	}

	conn, err := pgx.Connect(ctx, dbURI)
	if err != nil {
		panic(err)
	}
	defer conn.Close(ctx)

	service := common.CommonService{
		UserService: userservice.UserService{
			UserStore: &userstore.UserPostgres{
				Queries: userstoregenerated.New(conn),
			},
			SessionStore: &sessionstore.SessionPostgres{
				Queries: sessionstoregenerated.New(conn),
			},
			ClubStore: &clubs.ClubPostgres{
				Queries: clubstoregenerated.New(conn),
			},
		},
		PlayerService: playerservice.PlayerService{
			PlayerStore: &playerstore.PlayerPostgres{
				Queries: playerstoregenerated.New(conn),
			},
			UserStore: &userstore.UserPostgres{
				Queries: userstoregenerated.New(conn),
			},
		},
	}

	handler := common.NewController(service)

	rootRouter := chi.NewRouter()

	apiRouter := getPrivateRouter(conn)
	rootRouter.Mount("/api", oapiprivate.HandlerFromMux(oapiprivate.NewStrictHandler(handler, nil), apiRouter))

	publicRouter := commonmiddlewares.NewRouter()
	rootRouter.Mount("/", oapipublic.HandlerFromMux(oapipublic.NewStrictHandler(handler, nil), publicRouter))

	logrus.Info("Starting HTTP server")

	err = http.ListenAndServe(":"+os.Getenv("APP_PORT"), rootRouter)
	if err != nil {
		logrus.WithError(err).Panic("Unable to start HTTP server")
	}
}
