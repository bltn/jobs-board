package main

import (
	"context"
	"github.com/jackc/pgx/v4"
	"jobs-board/http/handlers"
	"jobs-board/repositories"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	dbUrl := os.Getenv("DATABASE_URL")
	if dbUrl == "" {
		log.Fatal("$DATABASE_URL must be set")
	}

	dbConnection, err := pgx.Connect(context.Background(), dbUrl)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer dbConnection.Close(context.Background())

	adzunaRepo := repositories.NewAdzunaJobRepository(dbConnection)
	jobsHandler := handlers.NewJobsHandler(adzunaRepo, jobListingTemplatePath())

	http.HandleFunc("/", jobsHandler.Index)
	err = http.ListenAndServe(":"+port, nil)

	if err != nil {
		log.Fatalf("Unable to ListenAndServe on port %v: %v\n", port, err)
	}
}

func jobListingTemplatePath() string {
	path, _ := filepath.Abs("./static/templates/job-listings.gohtml")
	return path
}

