package controller

import (
	"media-server/pkg/model"
	"media-server/pkg/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type MediaController struct {
	listMediasService  service.ListMediasService
	createMediaService service.CreateMediaService
}

func NewMediaController(listMediasService service.ListMediasService, createMediaService service.CreateMediaService) *MediaController {
	return &MediaController{
		listMediasService:  listMediasService,
		createMediaService: createMediaService,
	}
}

func (mc *MediaController) ListMedias(c *gin.Context) {
	medias, err := mc.listMediasService.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, medias)
}

func (mc *MediaController) CreateMedia(c *gin.Context) {
	var media model.Media
	if err := c.ShouldBindJSON(&media); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdMedia, err := mc.createMediaService.Execute(media)
	if err != nil {
		if _, ok := err.(validator.ValidationErrors); ok {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusCreated, createdMedia)

	c.JSON(http.StatusCreated, media)
}
