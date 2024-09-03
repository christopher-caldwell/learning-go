package routes

import (
	"example/rest-api-gin/routes/client"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateRouter(router *gin.Engine, dbPool *pgxpool.Pool) {
	client.CreateClientRouter(router, dbPool)
}
