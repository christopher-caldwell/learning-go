package routes

import (
	"example/rest-api-gin/routes/albums"
	"github.com/gin-gonic/gin"
)

func CreateRouter(router *gin.Engine) {
	albums.CreateAlbumRouter(router)
}
