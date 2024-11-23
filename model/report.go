package model

import (
	"time"

	"gorm.io/gorm"
)

type Report struct {
	ID         string `gorm:"primarykey"`
	InformerID string `gorm:"index"`
	TargetID   string
	TargetType string
	Status     string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
