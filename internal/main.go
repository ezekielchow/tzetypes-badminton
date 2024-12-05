package main

import (
	"common"
	commonmiddlewares "common/middlewares"
	"common/models"
	"common/oapiprivate"
	"common/oapipublic"
	"context"
	"database/sql"
	"fmt"
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

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
	"google.golang.org/api/option"
)

var firebaseAuth *auth.Client

func BearerTokenAuth(userStore usersStore.UserRepository) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				log.Println("OI?")
				http.Error(w, "Authorization header missing", http.StatusUnauthorized)
				return
			}

			token := strings.TrimPrefix(authHeader, "Bearer ")
			verifiedToken, err := verifyToken(token)
			if err != nil {
				log.Println("unable to verify token?", err.Error())
				http.Error(w, "Invalid token", http.StatusUnauthorized)
				return
			}

			email, ok := verifiedToken.Claims["email"].(string)
			if !ok {
				http.Error(w, "Email not found in token", http.StatusBadRequest)
				return
			}

			user, err := userStore.FindUserWithEmail(r.Context(), nil, email)
			if err != nil && !strings.Contains(sql.ErrNoRows.Error(), err.Error()) {
				log.Println("invalid user", err.Error())
				http.Error(w, "Invalid user", http.StatusUnauthorized)
				return
			}

			if user.ID == "" {
				user, err = userStore.CreateUser(r.Context(), nil, verifiedToken.UID, email, string(models.AccountTierPlayer))
				if err != nil {
					log.Println("failed to c8 user", err.Error())
					http.Error(w, "failed to c8 user", http.StatusInternalServerError)
					return
				}
			}

			ctx := context.WithValue(r.Context(), common.ContextUser, user)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func initFirebase() {
	serviceAccountJSON := os.Getenv("FIREBASE_SERVICE_ACCOUNT")
	filePath := "/tmp/firebase-service-account.json"

	// Write the secret to a temporary file
	err := os.WriteFile(filePath, []byte(serviceAccountJSON), 0600)
	if err != nil {
		log.Fatalf("Failed to write service account file: %v", err)
	}

	opt := option.WithCredentialsFile(filePath)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v", err)
	}

	firebaseAuth, err = app.Auth(context.Background())
	if err != nil {
		log.Fatalf("error initializing firebase auth: %v", err)
	}
}

func verifyToken(token string) (*auth.Token, error) {
	verifiedToken, err := firebaseAuth.VerifyIDToken(context.Background(), token)
	if err != nil {
		return nil, fmt.Errorf("error verifying token: %v", err)
	}
	return verifiedToken, nil
}

func getPrivateRouter(queries *databasegenerated.Queries) *chi.Mux {

	apiRoute := commonmiddlewares.NewRouter()
	apiRoute.Use(
		BearerTokenAuth(
			&usersStore.UserPostgres{
				Queries: queries,
			}))

	return apiRoute
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rr := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
		start := time.Now()

		next.ServeHTTP(rr, r)

		duration := time.Since(start)
		logEntry := logrus.WithFields(logrus.Fields{
			"method":   r.Method,
			"path":     r.URL.Path,
			"status":   rr.Status(),
			"duration": duration,
			"client":   r.RemoteAddr,
		})

		switch {
		case rr.Status() >= 500:
			logEntry.Error("Server error occurred")
		case rr.Status() >= 400:
			logEntry.Warn("Client error occurred")
		default:
			logEntry.Info("Request handled successfully")
		}
	})
}

func main() {
	ctx := context.Background()

	initFirebase()

	dbURI := os.Getenv("DB_URI")
	if len(dbURI) < 1 {
		dbURI = "postgresql://" + os.Getenv("POSTGRES_USER") + ":" + os.Getenv("POSTGRES_PASSWORD") + "@" + os.Getenv("POSTGRES_HOST") + "/" + os.Getenv("POSTGRES_DB") + "?sslmode=disable"
	}

	db := database.Database{}

	err := db.RunMigrations(dbURI, "file://database/migrations")
	if err != nil {
		panic(err)
	}

	config, err := pgxpool.ParseConfig(dbURI)
	if err != nil {
		panic(err)
	}

	config.ConnConfig.DefaultQueryExecMode = pgx.QueryExecModeCacheDescribe

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		panic(err)
	}
	defer pool.Close() // Ensure that the connection is properly closed on exit

	queries := databasegenerated.New(pool)

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
	rootRouter.Use(loggingMiddleware)

	apiRouter := getPrivateRouter(queries)
	rootRouter.Mount("/api", oapiprivate.HandlerFromMux(oapiprivate.NewStrictHandler(handler, nil), apiRouter))

	publicRouter := commonmiddlewares.NewRouter()
	rootRouter.Mount("/", oapipublic.HandlerFromMux(oapipublic.NewStrictHandler(handler, nil), publicRouter))

	logrus.Info("Starting HTTP server")

	server := &http.Server{
		Addr:              ":" + os.Getenv("APP_PORT"),
		Handler:           rootRouter,
		ReadTimeout:       10 * time.Second, // Allows time for reading typical requests
		WriteTimeout:      15 * time.Second, // Enough time to send responses
		IdleTimeout:       60 * time.Second, // Keeps connections alive for persistent clients
		ReadHeaderTimeout: 2 * time.Second,
	}

	err = server.ListenAndServe()
	if err != nil {
		logrus.WithError(err).Panic("Unable to start HTTP server")
	}
}
