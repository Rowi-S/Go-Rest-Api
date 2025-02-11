package main

import (
	"music/web-api/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/albums", handlers.GetAlbums)
	router.GET("/albums/:id", handlers.GetAlbumById)
	router.POST("/albums", handlers.PostAlbum)
	router.PATCH("/albums/:id", handlers.PatchAlbum)
	router.DELETE("/albums/:id", handlers.DeleteAlbum)

	router.Run("localhost:8080")
}
