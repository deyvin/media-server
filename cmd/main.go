package main

import (
	"media-server/app/controller"
	"media-server/app/database"
	"media-server/app/repository"
	"media-server/app/service"
	"media-server/app/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()
	db := database.DB

	mediaRepo := repository.PgMediaRepository(db)
	listMediasUseCase := usecase.NewListMediasUseCase(mediaRepo)
	mediaService := service.NewListMediasService(listMediasUseCase)
	mediaController := controller.NewMediaController(mediaService)

	r := gin.Default()
	r.GET("/medias", mediaController.ListMedias)

	r.Run()
}
