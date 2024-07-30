package controller

import (
	"media-server/app/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MediaController struct {
	listMediasService service.ListMediasService
}

func NewMediaController(s service.ListMediasService) *MediaController {
	return &MediaController{
		listMediasService: s,
	}
}

func (mc *MediaController) ListMedias(c *gin.Context) {
	medias, err := mc.listMediasService.ListMedias()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, medias)
}
