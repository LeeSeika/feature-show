package report

import (
	"context"

	"github.com/leeseika/feature-show/model"
	"go.uber.org/zap"
)

type QueryParam struct {
	Page int
	Size int
}

func (rs *reportService) GetReportList(ctx context.Context, informerID string, queryParam QueryParam) ([]*model.Report, error) {
	reports := make([]*model.Report, 0)
	if err := rs.db.WithContext(ctx).Where("informer_id = ?", informerID).Offset((queryParam.Page - 1) * queryParam.Size).Limit(queryParam.Size).Find(&reports).Error; err != nil {
		zap.L().Error("get report list failed", zap.Error(err))
		return nil, err
	}

	return reports, nil
}
