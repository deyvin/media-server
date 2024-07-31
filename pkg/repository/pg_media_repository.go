package repository

import (
	"media-server/pkg/model"

	"gorm.io/gorm"
)

type MediaRepository interface {
	ListMedias() ([]model.Media, error)
}

type mediasRepository struct {
	db *gorm.DB
}

func PgMediaRepository(db *gorm.DB) MediaRepository {
	return &mediasRepository{db}
}

func (r *mediasRepository) ListMedias() ([]model.Media, error) {
	var medias []model.Media
	result := r.db.Find(&medias)
	return medias, result.Error
}