package main

import (
	"media-server/pkg/controller"
	"media-server/pkg/database"
	"media-server/pkg/repository"
	"media-server/pkg/service"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()
	db := database.DB

	mediaRepo := repository.PgMediaRepository(db)
	mediaService := service.NewListMediasService(mediaRepo)
	mediaController := controller.NewMediaController(mediaService)

	r := gin.Default()
	r.GET("/medias", mediaController.ListMedias)

	r.Run()
}
