package report

import (
	"github.com/gin-gonic/gin"
	reportService "github.com/leeseika/feature-show/services/report"
)

type AddReportInput struct {
	TargetID   string `json:"target_id" binding:"required"`
	TargetType string `json:"target_type" binding:"required"`
}

func AddReport(c *gin.Context) {
	var input AddReportInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// todo: get informerID from jwt token
	informerID := "todo"

	if err := reportService.Get().AddReport(c, informerID, input.TargetID, input.TargetType); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "success"})
}
