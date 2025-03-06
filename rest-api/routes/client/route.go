package client

import (
	"context"
	"fmt"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5/pgxpool"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Client struct {
	ClientId          int16   `json:"clientId"`
	CompanyName       *string `json:"companyName"`
	CompanyWebsiteUrl *string `json:"companyWebsiteUrl"`
	CompanyLogoUrl    *string `json:"companyLogoUrl"`
}

var ctx = context.Background()

// context *gin.Context
func getClients(dbPool *pgxpool.Pool) func(context *gin.Context) {
	return func(context *gin.Context) {
		var clients []Client
		queryError := pgxscan.Select(ctx, dbPool, &clients, `SELECT client_id, company_name, company_website_url, company_logo_url FROM client`)

		if queryError != nil {
			fmt.Fprintf(os.Stderr, "Unable to query: %v\n", queryError)
			os.Exit(1)
		}
		fmt.Print(len(clients))
		context.JSON(http.StatusOK, clients)
	}
}

func CreateClientRouter(router *gin.Engine, dbPool *pgxpool.Pool) {
	router.GET("/clients", getClients(dbPool))
}
