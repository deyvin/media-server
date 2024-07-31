package service

import (
	"media-server/pkg/model"
	"media-server/pkg/repository"
)

type ListMediasService interface {
	Execute() ([]model.Media, error)
}

type listMediasService struct {
	mediaRepo repository.MediaRepository
}

func NewListMediasService(mediaRepo repository.MediaRepository) ListMediasService {
	return &listMediasService{
		mediaRepo: mediaRepo,
	}
}

func (uc *listMediasService) Execute() ([]model.Media, error) {
	return uc.mediaRepo.ListMedias()
}
