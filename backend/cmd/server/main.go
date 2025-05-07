package main

import (
	"database/sql"
	"fmt"
	"os"
	"pala/backend/pkg/api"
	"pala/backend/pkg/app"
	"pala/backend/pkg/repository"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "this is the startup error: %s\n", err)
		os.Exit(1)
	}
}

func run() error {
	connectionString := "postgres://postgres:example@db:5432/postgres?sslmode=disable"

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
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
	}))

	voterService := api.NewVoterService(storage)
	voteService := api.NewVoteService(storage)
	candidateService := api.NewCandidateService(storage)
	partyService := api.NewPartyService(storage)
	electionService := api.NewElectionService(storage)

	server := app.NewServer(router, voterService, voteService, partyService, candidateService, electionService)

	err = server.Run()

	if err != nil {
		return err
	}

	return nil
}

func setupDatabase(connString string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connString)

	if err != nil {
		return nil, err
	}

	err = db.Ping()

	if err != nil {
		return nil, err
	}

	return db, nil
}
