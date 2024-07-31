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
	listmediaService := service.NewListMediasService(mediaRepo)
	createMediaService := service.NewCreateMediaService(mediaRepo)

	mediaController := controller.NewMediaController(listmediaService, createMediaService)

	r := gin.Default()
	r.GET("/medias", mediaController.ListMedias)
	r.POST("/medias", mediaController.CreateMedia)

	r.Run()
}
