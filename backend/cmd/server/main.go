package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"os"
	"pala/backend/pkg/api"
	"pala/backend/pkg/app"
	"pala/backend/pkg/repository"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "this is the startup error: %s\n", err)
		os.Exit(1)
	}
}

// func run will be responsible for setting up db connections, routers etc
func run() error {
	connectionString := "postgres://postgres:postgres@localhost/postgres?sslmode=disable"

	db, err := setupDatabase(connectionString)
	if err != nil {
		return err
	}

	storage := repository.NewStorage(db)

	err = storage.RunMigrations(connectionString)

	if err != nil {
		return err
	}

	router := gin.Default()
	router.Use(cors.Default())

	voterService := api.NewVoterService(storage)

	weightService := api.NewWeightService(storage)

	server := app.NewServer(router, voterService, weightService)

	// start the server
	err = server.Run()

	if err != nil {
		return err
	}

	return nil
}

func setupDatabase(connString string) (*sql.DB, error) {
	// change "postgres" for whatever supported database you want to use
	db, err := sql.Open("postgres", connString)

	if err != nil {
		return nil, err
	}

	// ping the DB to ensure that it is connected
	err = db.Ping()

	if err != nil {
		return nil, err
	}

	return db, nil
}
