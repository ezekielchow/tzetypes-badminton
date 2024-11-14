package utils

import (
	"database/sql"
	"fmt"
	"log"
	"net"
	"strconv"
	"time"
	"tzetypes-badminton/database"

	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
)

func InitDockerTest() (*dockertest.Pool, *dockertest.Resource, string) {
	// uses a sensible default on windows (tcp/http) and linux/osx (socket)
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not construct pool: %s", err)
	}

	// uses pool to try to connect to Docker
	err = pool.Client.Ping()
	if err != nil {
		log.Fatalf("Could not connect to Docker: %s", err)
	}

	port, err := GetFreePort()
	if err != nil {
		log.Fatalf("Unable to get free port on host container: %s", err)
	}

	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "16.4",
		Env: []string{
			"POSTGRES_PASSWORD=postgres12345",
			"POSTGRES_USER=badminton_user",
			"POSTGRES_DB=test_badminton",
			"listen_addresses = '*'",
		},
		ExposedPorts: []string{"5432"},
		PortBindings: map[docker.Port][]docker.PortBinding{
			"5432/tcp": {
				{HostIP: "0.0.0.0", HostPort: strconv.Itoa(port)},
			},
		},
	}, func(config *docker.HostConfig) {
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{Name: "no"} // Important option when container crash and you want to debug container
	})

	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	hostAndPort := resource.GetHostPort("5432/tcp")
	databaseUrl := fmt.Sprintf("postgres://badminton_user:postgres12345@%s/test_badminton?sslmode=disable", hostAndPort)

	if err := resource.Expire(30); err != nil { // Tell docker to hard kill the container in 30 seconds
		log.Fatalf("Resource couldnt expire: %s", err)
	}

	pool.MaxWait = 30 * time.Second
	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	if err = pool.Retry(func() error {
		db, err := sql.Open("postgres", databaseUrl)
		if err != nil {
			return err
		}
		return db.Ping()
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	db := database.Database{}

	err = db.RunMigrations(databaseUrl, "file://../../database/migrations")
	if err != nil {
		log.Fatalf("Unable to run migrations: %s", err)
	}

	return pool, resource, databaseUrl
}

func GetFreePort() (int, error) {
	// Open a listener on any available address and port
	listener, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		return 0, err
	}
	defer listener.Close()

	// Extract the port number assigned to the listener
	addr := listener.Addr().(*net.TCPAddr)
	return addr.Port, nil
}
