package main

import (
	"common"
	commonmiddlewares "common/middlewares"
	"common/models"
	"common/oapiprivate"
	"common/oapipublic"
	"context"
	"database/sql"
	games "games/service"
	gamestore "games/store"
	"log"
	"net/http"
	"os"
	playerservice "players/service"
	playerstore "players/store"
	sessionstore "sessions/store"
	"strings"
	"time"
	clubs "tzetypes-badminton/clubs/store"
	database "tzetypes-badminton/database"
	databasegenerated "tzetypes-badminton/database/generated"
	usersService "users/service"
	usersStore "users/store"

	"github.com/go-chi/chi/v5"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5"
	"github.com/sirupsen/logrus"
)

func BearerTokenAuth(sessionStore sessionstore.SessionRepository, userStore usersStore.UserRepository) func(next http.Handler) http.Handler {
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

			user, err := userStore.FindUserWithID(r.Context(), nil, session.UserID)
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
	session, err := sessionStore.FindSessionWithSessionID(ctx, nil, token)
	if err != nil && !strings.Contains(sql.ErrNoRows.Error(), err.Error()) {
		return models.Session{}, err
	}

	return session, err
}

func getPrivateRouter(queries *databasegenerated.Queries) *chi.Mux {

	apiRoute := commonmiddlewares.NewRouter()
	apiRoute.Use(
		BearerTokenAuth(
			&sessionstore.SessionPostgres{
				Queries: queries,
			},
			&usersStore.UserPostgres{
				Queries: queries,
			}))

	return apiRoute
}

func main() {
	ctx := context.Background()

	for _, env := range os.Environ() {
		log.Println(env)
	}

	dbURI := os.Getenv("DB_URI")
	if len(dbURI) < 1 {
		dbURI = "postgresql://" + os.Getenv("POSTGRES_USER") + ":" + os.Getenv("POSTGRES_PASSWORD") + "@" + os.Getenv("POSTGRES_HOST") + "/" + os.Getenv("POSTGRES_DB") + "?sslmode=disable"
	}

	db := database.Database{}

	err := db.RunMigrations(dbURI, "file://database/migrations")
	if err != nil {
		panic(err)
	}

	conn, err := pgx.Connect(ctx, dbURI)
	if err != nil {
		panic(err)
	}
	defer conn.Close(ctx)

	pool, err := db.Open(ctx, dbURI)
	if err != nil {
		panic(err)
	}
	defer pool.Close() // Ensure that the connection is properly closed on exit

	queries := databasegenerated.New(conn)

	service := common.CommonService{
		UserService: usersService.UserService{
			UserStore: &usersStore.UserPostgres{
				Queries: queries,
			},
			SessionStore: &sessionstore.SessionPostgres{
				Queries: queries,
			},
			ClubStore: &clubs.ClubPostgres{
				Queries: queries,
			},
			PgxPool: pool,
		},
		PlayerService: playerservice.PlayerService{
			PlayerStore: &playerstore.PlayerPostgres{
				Queries: queries,
			},
			UserStore: &usersStore.UserPostgres{
				Queries: queries,
			},
			ClubStore: &clubs.ClubPostgres{
				Queries: queries,
			},
			PgxPool: pool,
		},
		GameService: games.GameService{
			ClubStore: &clubs.ClubPostgres{
				Queries: queries,
			},
			GameStore: &gamestore.GamePostgres{
				Queries: queries,
			},
			PgxPool: pool,
		},
	}

	handler := common.NewController(service)

	rootRouter := chi.NewRouter()

	apiRouter := getPrivateRouter(queries)
	rootRouter.Mount("/api", oapiprivate.HandlerFromMux(oapiprivate.NewStrictHandler(handler, nil), apiRouter))

	publicRouter := commonmiddlewares.NewRouter()
	rootRouter.Mount("/", oapipublic.HandlerFromMux(oapipublic.NewStrictHandler(handler, nil), publicRouter))

	logrus.Info("Starting HTTP server")

	err = http.ListenAndServe(":"+os.Getenv("APP_PORT"), rootRouter)
	if err != nil {
		logrus.WithError(err).Panic("Unable to start HTTP server")
	}
}
