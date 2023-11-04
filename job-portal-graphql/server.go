package main

import (
	"context"
	"errors"
	"fmt"
	"graphql/database"
	"graphql/graph"
	"graphql/models"
	"graphql/repo"
	"graphql/service"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8081"

func main() {

	serready, err := startApplication()
	if err != nil {
		log.Fatal(err)
		return
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		S: serready,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

}
func startApplication() (service.UserService, error) {

	//initialize to database connection
	db, err := database.Connection()
	if err != nil {
		return nil, err
	}
	//to get db instance
	pg, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get database instance: %w ", err)
	}
	//create context to ping database for check table alive or not
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	err = pg.PingContext(ctx)
	if err != nil {

		return nil, fmt.Errorf("database is not connected: %w", err)
	}
	//
	err = db.Migrator().AutoMigrate(&models.User{}, &models.Company{}, &models.Job{})
	if err != nil {

		return nil, fmt.Errorf("database is not connected: %w", err)
	}

	repoinit, err := repo.NewRepo(db)
	if err != nil {
		return nil, errors.New("repo not initialzed")
	}

	servc, err := service.NewService(repoinit)
	if err != nil {
		return nil, errors.New("service not initialzed")
	}

	return servc, nil

}
