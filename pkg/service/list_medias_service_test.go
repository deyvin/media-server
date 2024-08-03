package service

import (
	"media-server/pkg/model"
	"media-server/pkg/test"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExecuteListMediasService(t *testing.T) {
	mockRepo := new(test.MockMediaRepository)

	service := listMediasService{
		mediaRepo: mockRepo,
	}

	media := model.Media{
		ID:          "1",
		Title:       "test",
		Description: "test",
		Path:        "/var/null",
	}

	mockRepo.On("ListMedias").Return([]model.Media{media}, nil)

	medias, err := service.Execute()

	assert.NoError(t, err)
	assert.NotNil(t, medias)
	assert.Greater(t, len(medias), 0)
	mockRepo.AssertExpectations(t)
}
