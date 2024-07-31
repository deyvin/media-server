package model

import "time"

type Media struct {
	ID          string     `gorm:"primaryKey;size:50" json:"id"`
	Title       string     `gorm:"not null" json:"title" validate:"required"`
	Description string     `gorm:"not null" json:"description" validate:"required"`
	Path        string     `gorm:"not null" json:"path" validate:"required"`
	CreatedAt   *time.Time `gorm:"defaut:current_timestamp" json:"created_at"`
	UpdatedAt   *time.Time
}

func (Media) TableName() string {
	return "medias"
}
