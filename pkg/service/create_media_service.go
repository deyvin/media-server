package service

import (
	"media-server/pkg/model"
	"media-server/pkg/repository"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type CreateMediaService interface {
	Execute(media model.Media) (model.Media, error)
}

type createMediaService struct {
	mediaRepo repository.MediaRepository
	validate  *validator.Validate
}

func NewCreateMediaService(mediaRepo repository.MediaRepository) CreateMediaService {
	return &createMediaService{
		mediaRepo: mediaRepo,
		validate:  validator.New(),
	}
}

func (uc *createMediaService) Execute(media model.Media) (model.Media, error) {

	err := uc.validate.Struct(media)
	if err != nil {
		return media, err
	}

	media.ID = uuid.New().String()
	err = uc.mediaRepo.CreateMedia(media)
	return media, err
}
