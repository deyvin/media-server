package service

import (
	"media-server/pkg/model"
	"media-server/pkg/repository"

	"github.com/google/uuid"
)

type CreateMediaService interface {
	Execute(media model.Media) (model.Media, error)
}

type createMediaService struct {
	mediaRepo repository.MediaRepository
}

func NewCreateMediaService(mediaRepo repository.MediaRepository) CreateMediaService {
	return &createMediaService{
		mediaRepo: mediaRepo,
	}
}

func (uc *createMediaService) Execute(media model.Media) (model.Media, error) {
	media.ID = uuid.New().String()
	err := uc.mediaRepo.CreateMedia(media)
	return media, err
}
