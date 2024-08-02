package service

import (
	"media-server/pkg/model"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockMediaRepository struct {
	mock.Mock
}

func (m *MockMediaRepository) CreateMedia(media model.Media) (model.Media, error) {
	args := m.Called(media)
	return args.Get(0).(model.Media), args.Error(1)
}

func (m *MockMediaRepository) ListMedias() ([]model.Media, error) {
	args := m.Called()
	return args.Get(0).([]model.Media), args.Error(1)
}

func TestExecuteCreateMediaService(t *testing.T) {
	mockRepo := new(MockMediaRepository)
	validator := validator.New()

	service := createMediaService{
		mediaRepo: mockRepo,
		validate:  validator,
	}

	media := model.Media{
		Title:       "test",
		Description: "test",
		Path:        "/var/null",
	}

	mockRepo.On("CreateMedia", mock.MatchedBy(func(m model.Media) bool {
		return m.Title == media.Title && m.Description == media.Description && m.Path == media.Path
	})).Return(media, nil)

	createdMedia, err := service.Execute(media)

	assert.NoError(t, err)
	assert.Equal(t, media.ID, createdMedia.ID)
	assert.Equal(t, media.Title, createdMedia.Title)
	assert.Equal(t, media.Description, createdMedia.Description)
	mockRepo.AssertExpectations(t)
}

func TestExecuteCreateMediaServiceWithRequiredError(t *testing.T) {
	mockRepo := new(MockMediaRepository)
	validator := validator.New()

	service := createMediaService{
		mediaRepo: mockRepo,
		validate:  validator,
	}

	media := model.Media{
		Title:       "",
		Description: "test",
		Path:        "/var/null",
	}

	createdMedia, err := service.Execute(media)

	assert.Error(t, err)
	assert.Equal(t, "Key: 'Media.Title' Error:Field validation for 'Title' failed on the 'required' tag", err.Error())
	assert.Equal(t, media.ID, createdMedia.ID)
}