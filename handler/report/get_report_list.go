package report

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/leeseika/feature-show/constant"
	reportService "github.com/leeseika/feature-show/services/report"
)

type GetReportListInput struct {
	Page int `form:"page"`
	Size int `form:"size"`
}

func GetReportList(c *gin.Context) {

	// todo: get informerID from jwt token
	informerID := "todo"

	var input GetReportListInput
	if err := c.ShouldBindQuery(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.Page <= 0 {
		input.Page = constant.DefaultPage
	}
	if input.Size <= 0 {
		input.Size = constant.DefaultSize
	}

	queryParam := reportService.QueryParam{
		Page: input.Page,
		Size: input.Size,
	}

	reports, err := reportService.Get().GetReportList(c, informerID, queryParam)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, reports)

}
