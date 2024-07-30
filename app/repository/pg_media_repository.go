package repository

import (
	"media-server/app/entity"

	"gorm.io/gorm"
)

type MediaRepository interface {
	ListMedias() ([]entity.Media, error)
}

type mediasRepository struct {
	db *gorm.DB
}

func PgMediaRepository(db *gorm.DB) MediaRepository {
	return &mediasRepository{db}
}

func (r *mediasRepository) ListMedias() ([]entity.Media, error) {
	var medias []entity.Media
	result := r.db.Find(&medias)
	return medias, result.Error
}
