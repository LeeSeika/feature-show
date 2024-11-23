package services

import (
	"github.com/leeseika/feature-show/services/report"
	"gorm.io/gorm"
)

func Init(db *gorm.DB) {
	report.Init(db)
}
