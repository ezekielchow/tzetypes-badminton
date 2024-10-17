package users

import (
	"common/utils"
	"context"
	"log"
	"os"
	sessionstore "sessions/store"
	"testing"
	clubs "tzetypes-badminton/clubs/store"
	"tzetypes-badminton/database"
	databasegenerated "tzetypes-badminton/database/generated"
	users "users/store"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ory/dockertest/v3"
)

var DBPool *dockertest.Pool
var DBResource *dockertest.Resource
var DBUri string
var pgxPool *pgxpool.Pool
var pgxConn *pgx.Conn

func TestMain(m *testing.M) {

	pool, resource, url := utils.InitDockerTest()
	DBPool = pool
	DBResource = resource
	DBUri = url

	defer func() {
		if err := pool.Purge(resource); err != nil {
			log.Fatalf("Could not purge resource: %s", err)
		}
	}()

	// Run tests
	code := m.Run()

	ctx := context.Background()
	defer pgxPool.Close()
	defer pgxConn.Close(ctx)

	// Exit with the appropriate code
	os.Exit(code)
}

func InitService(ctx context.Context) UserService {
	db := database.Database{}

	conn, err := pgx.Connect(ctx, DBUri)
	if err != nil {
		log.Fatalf("unable to connect to test db: %s", err)
	}
	pgxConn = conn

	pool, err := db.Open(ctx, DBUri)
	if err != nil {
		panic(err)
	}
	pgxPool = pool

	queries := databasegenerated.New(conn)

	userService := UserService{
		UserStore: &users.UserPostgres{
			Queries: queries,
		},
		SessionStore: &sessionstore.SessionPostgres{
			Queries: queries,
		},
		ClubStore: &clubs.ClubPostgres{
			Queries: queries,
		},
		PgxPool: pgxPool,
	}

	return userService
}
