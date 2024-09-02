package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

type Client struct {
	Client_id    int16
	Company_name *string
	Website_url  *string
	Logo_url     *string
}

func main() {

	envErr := godotenv.Load()
	if envErr != nil {
		log.Fatal("Error loading .env file")
	}

	ctx := context.Background()
	dbPool, err := pgxpool.New(ctx, os.Getenv("DATABASE_URL"))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	defer dbPool.Close()

	var clients []*Client
	queryError := pgxscan.Select(ctx, dbPool, &clients, `SELECT client_id, company_name, website_url, logo_url FROM client`)

	if queryError != nil {
		fmt.Fprintf(os.Stderr, "Unable to query: %v\n", queryError)
		os.Exit(1)
	}

	fmt.Print(len(clients))
}
