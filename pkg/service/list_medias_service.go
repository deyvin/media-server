package service

import (
	"media-server/pkg/entity"
	"media-server/pkg/repository"
)

type ListMediasService interface {
	Execute() ([]entity.Media, error)
}

type listMediasService struct {
	mediaRepo repository.MediaRepository
}

func NewListMediasService(mediaRepo repository.MediaRepository) ListMediasService {
	return &listMediasService{
		mediaRepo: mediaRepo,
	}
}

func (uc *listMediasService) Execute() ([]entity.Media, error) {
	return uc.mediaRepo.ListMedias()
}
