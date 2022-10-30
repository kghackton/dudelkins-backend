package controllers

import (
	"encoding/csv"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"

	"dudelkins/internal/interfaces"
	"dudelkins/internal/objects/bo"
	"dudelkins/pkg/logger"
)

type ApplicationController struct {
	ApplicationService interfaces.IApplicationService
}

func (c *ApplicationController) Create(ctx echo.Context) (err error) {
	var (
		logFields = []interface{}{"path", ctx.Path(), "method", ctx.Request().Method}
	)

	csvFile, err := os.Open("/data/3000.csv")
	if err != nil {
		logger.Errorw(err.Error(), logFields...)
		return ctx.JSON(http.StatusInternalServerError, echo.Map{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
	}
	defer csvFile.Close()

	csvReader := csv.NewReader(csvFile)
	csvReader.Comma = '$'
	csvReader.ReuseRecord = true

	records, err := csvReader.ReadAll()
	if err != nil {
		logger.Errorw(err.Error(), logFields...)
		return ctx.JSON(http.StatusInternalServerError, echo.Map{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
	}

	logger.Debugf("record: %+v", records[1])

	application, err := bo.NewApplicationFromRecord(records[1])
	if err != nil {
		logger.Errorw(err.Error(), logFields...)
		return ctx.JSON(http.StatusInternalServerError, echo.Map{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
	}
	if err = c.ApplicationService.Create(ctx.Request().Context(), application); err != nil {
		logger.Errorw(err.Error(), logFields...)
		return ctx.JSON(http.StatusInternalServerError, echo.Map{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
	}

	logger.Infow("", logFields...)

	return ctx.JSON(http.StatusOK, echo.Map{
		"code": http.StatusOK,
	})
}
