package main

import (
	"context"
	"fmt"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

type Client struct {
	ClientId          int16              `json:"clientId"`
	DtCreated         pgtype.Timestamptz `json:"dtCreated"`
	DtModified        pgtype.Timestamptz `json:"dtModified"`
	StripeCustomerSid string             `json:"stripeCustomerSid"`
	CompanyName       string             `json:"companyName"`
	CompanyWebsiteUrl string             `json:"companyWebsiteUrl"`
	CompanyLogoUrl    string             `json:"companyLogoUrl"`
	CompanyPhone      string             `json:"companyPhone"`
}

type ClientRepo interface {
	GetMany(ctx context.Context) ([]Client, error)
	//Add(client *Client) (Client, error)
	//Update(client *Client) (Client, error)
}

type PostgresClientRepo struct {
	dbReadHandle  *pgxpool.Pool
	dbWriteHandle *pgxpool.Pool
}

func NewPostgresClientRepo(dbReadHandle *pgxpool.Pool, dbWriteHandle *pgxpool.Pool) *PostgresClientRepo {
	return &PostgresClientRepo{dbReadHandle: dbReadHandle, dbWriteHandle: dbWriteHandle}
}

type ClientService struct {
	repo ClientRepo
}

func NewClientService(repo ClientRepo) *ClientService {
	return &ClientService{repo}
}

func (repo *PostgresClientRepo) GetMany(ctx context.Context) ([]Client, error) {
	var clients []Client
	queryError := pgxscan.Select(ctx, repo.dbReadHandle, &clients, `SELECT client_id, company_name, stripe_customer_sid, company_logo_url, company_phone FROM client`)

	if queryError != nil {
		fmt.Fprintf(os.Stderr, "Unable to query: %v\n", queryError)
		return nil, queryError
	}

	fmt.Print(len(clients))

	return clients, nil
}

func main() {
	envErr := godotenv.Load()
	if envErr != nil {
		log.Fatal("Error loading .env file")
	}

	ctx := context.Background()
	dbReadHandle, readHandleErr := pgxpool.New(ctx, os.Getenv("DATABASE_URL"))
	if readHandleErr != nil {
		log.Fatal("Read handle connection failed")
	}
	dbWriteHandle, writeHandleErr := pgxpool.New(ctx, os.Getenv("DATABASE_URL"))
	if writeHandleErr != nil {
		log.Fatal("Write handle connection failed")
	}

	clientRepo := NewPostgresClientRepo(dbReadHandle, dbWriteHandle)
	clientService := NewClientService(clientRepo)

	router := gin.Default()
	router.GET("/clients", func(context *gin.Context) {
		manyClients, manyClientsErr := clientService.repo.GetMany(ctx)
		if manyClientsErr != nil {
			log.Fatal("Many clients query failed")
		}
		fmt.Println("Len", len(manyClients))
		context.JSON(http.StatusOK, manyClients)
	})

	routerError := router.Run("localhost:8080")
	if routerError != nil {
		log.Fatal("Router failed to start")
	}
}
