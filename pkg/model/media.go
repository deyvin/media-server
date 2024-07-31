package model

import "time"

type Media struct {
	ID          string     `gorm:"primaryKey;size:50" json:"id"`
	Title       string     `gorm:"not null" json:"title"`
	Description string     `gorm:"not null" json:"description"`
	Path        string     `gorm:"not null" json:"path"`
	CreatedAt   *time.Time `gorm:"defaut:current_timestamp" json:"created_at"`
	UpdatedAt   *time.Time
}

func (Media) TableName() string {
	return "medias"
}
