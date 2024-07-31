package controller

import (
	"media-server/pkg/service"
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
	medias, err := mc.listMediasService.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, medias)
}
