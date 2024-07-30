package service

import (
	"media-server/app/entity"
	"media-server/app/usecase"
)

type ListMediasService interface {
	ListMedias() ([]entity.Media, error)
}

type listMediasService struct {
	listMediasUseCase usecase.ListMediasUseCase
}

func NewListMediasService(listMediasUseCase usecase.ListMediasUseCase) ListMediasService {
	return &listMediasService{
		listMediasUseCase: listMediasUseCase,
	}
}

func (s *listMediasService) ListMedias() ([]entity.Media, error) {
	return s.listMediasUseCase.Execute()
}
