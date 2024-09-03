package main

import (
	"context"
	"example/rest-api-gin/routes"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var ctx = context.Background()

func main() {
	envErr := godotenv.Load()
	if envErr != nil {
		log.Fatal("Error loading .env file")
	}

	dbPool, connectionErr := pgxpool.New(ctx, os.Getenv("DATABASE_URL"))
	if connectionErr != nil {
		_, formatErr := fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", connectionErr)
		if formatErr != nil {
			log.Fatal(formatErr)
		}
		os.Exit(1)
	}
	defer dbPool.Close()

	router := gin.Default()
	routes.CreateRouter(router, dbPool)
	routerError := router.Run("localhost:8080")

	if routerError != nil {
		log.Fatal(routerError)
	}
}
