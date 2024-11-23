package report

import (
	"context"

	"github.com/leeseika/feature-show/constant"
	"github.com/leeseika/feature-show/model"
	"github.com/leeseika/feature-show/pkg/nanoid"
	"go.uber.org/zap"
)

func (rs *reportService) AddReport(ctx context.Context, informerID string, targetID string, targetType string) error {
	// todo: check if the target exists
	// todo: check if the informer has reported the target

	report := model.Report{
		ID:         nanoid.Gen(),
		InformerID: informerID,
		TargetID:   targetID,
		TargetType: targetType,
		Status:     constant.ReportStatusPending,
	}

	if err := rs.db.WithContext(ctx).Create(&report).Error; err != nil {
		zap.L().Error("create report failed", zap.Error(err))
		return err
	}

	return nil
}
