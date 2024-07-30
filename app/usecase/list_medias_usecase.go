package usecase

import (
	"media-server/app/entity"
	"media-server/app/repository"
)

type ListMediasUseCase interface {
	Execute() ([]entity.Media, error)
}

type listMediasUseCase struct {
	mediaRepo repository.MediaRepository
}

func NewListMediasUseCase(mediaRepo repository.MediaRepository) ListMediasUseCase {
	return &listMediasUseCase{
		mediaRepo: mediaRepo,
	}
}

func (uc *listMediasUseCase) Execute() ([]entity.Media, error) {
	return uc.mediaRepo.ListMedias()
}
