package report

import (
	"context"
	"sync"

	"github.com/leeseika/feature-show/model"
	"gorm.io/gorm"
)

type ReportService interface {
	AddReport(ctx context.Context, informerID string, targetID string, targetType string) error

	GetReportList(ctx context.Context, informerID string, queryParam QueryParam) ([]*model.Report, error)
}

type reportService struct {
	db *gorm.DB
}

var _reportService ReportService
var _initOnce sync.Once

func Init(db *gorm.DB) {
	_initOnce.Do(func() {
		_reportService = &reportService{
			db: db,
		}
	})
}

func Get() ReportService {
	return _reportService
}
