package controllers

import (
	"encoding/csv"
	"io"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"

	"dudelkins/internal/interfaces"
	"dudelkins/internal/objects/bo"
	"dudelkins/internal/objects/dto"
	"dudelkins/pkg/logger"
)

type ApplicationController struct {
	ApplicationUploadService interfaces.IApplicationUploadService
	ApplicationViewService   interfaces.IApplicationViewService
}

func (c *ApplicationController) Create(ctx echo.Context) (err error) {
	var (
		logFields = []interface{}{"path", ctx.Path(), "method", ctx.Request().Method}
		request   struct {
			RowsAmount int `json:"rowsAmount" validate:"gte=0"`
		}
	)

	if err = ctx.Bind(&request); err != nil {
		logger.Warnw(err.Error(), logFields...)
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"code": http.StatusBadRequest,
			"msg":  err.Error(),
		})
	}
	if err = ctx.Validate(request); err != nil {
		logger.Warnw(err.Error(), logFields...)
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"code": http.StatusBadRequest,
			"msg":  err.Error(),
		})
	}

	csvFile, err := os.Open("/data/applications.csv")
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

	if request.RowsAmount == 0 {
		request.RowsAmount = 300
	}
	for i := 0; i < request.RowsAmount; i++ {
		record, err := csvReader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			logger.Errorw(err.Error(), logFields...)
			return ctx.JSON(http.StatusInternalServerError, echo.Map{
				"code": http.StatusInternalServerError,
				"msg":  err.Error(),
			})
		}
		if i == 0 {
			continue
		}

		application, err := bo.NewApplicationFromRecord(record)
		if err != nil {
			logger.Errorw(err.Error(), logFields...)
			return ctx.JSON(http.StatusInternalServerError, echo.Map{
				"code": http.StatusInternalServerError,
				"msg":  err.Error(),
			})
		}
		if err = c.ApplicationUploadService.Create(ctx.Request().Context(), application); err != nil {
			logger.Errorw(err.Error(), logFields...)
			return ctx.JSON(http.StatusInternalServerError, echo.Map{
				"code": http.StatusInternalServerError,
				"msg":  err.Error(),
			})
		}

		if i%100 == 0 {
			logger.Debugf("%d", i)
		}
	}

	logger.Infow("", logFields...)

	return ctx.JSON(http.StatusOK, echo.Map{
		"code": http.StatusOK,
	})
}

func (c *ApplicationController) Get(ctx echo.Context) (err error) {
	var (
		logFields = []interface{}{"path", ctx.Path(), "method", ctx.Request().Method}
		request   dto.ApplicationRetrieveOpts
	)

	if err = ctx.Bind(&request); err != nil {
		logger.Warnw(err.Error(), logFields...)
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"code": http.StatusBadRequest,
			"msg":  err.Error(),
		})
	}
	if err = ctx.Validate(request); err != nil {
		logger.Warnw(err.Error(), logFields...)
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"code": http.StatusBadRequest,
			"msg":  err.Error(),
		})
	}

	applications, err := c.ApplicationViewService.Get(ctx.Request().Context(), &bo.ApplicationRetrieveOpts{
		ClosedFrom:     request.ClosedFrom,
		ClosedTo:       request.ClosedTo,
		IsAbnormal:     request.IsAbnormal,
		AnomalyClasses: request.AnomalyClasses,
		CategoryIds:    request.CategoryIds,
		DefectIds:      request.DefectIds,
		Region:         request.Region,
		District:       request.District,
		UNOM:           request.UNOM,
		Limit:          request.Limit,
		Offset:         request.Offset,
	})
	if err != nil {
		logger.Errorw(err.Error(), logFields...)
		return ctx.JSON(http.StatusInternalServerError, echo.Map{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
	}

	logger.Infow("", logFields...)

	return ctx.JSON(http.StatusOK, echo.Map{
		"code": http.StatusOK,
		"data": applications.ToDto(),
	})
}

func (c *ApplicationController) GetAnomalyClassesStats(ctx echo.Context) (err error) {
	var (
		logFields = []interface{}{"path", ctx.Path(), "method", ctx.Request().Method}
		request   dto.ApplicationRetrieveOpts
	)

	if err = ctx.Bind(&request); err != nil {
		logger.Warnw(err.Error(), logFields...)
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"code": http.StatusBadRequest,
			"msg":  err.Error(),
		})
	}
	if err = ctx.Validate(request); err != nil {
		logger.Warnw(err.Error(), logFields...)
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"code": http.StatusBadRequest,
			"msg":  err.Error(),
		})
	}

	anomalyClassesCounterMap, err := c.ApplicationViewService.CountAnomalyClasses(ctx.Request().Context(), &bo.ApplicationRetrieveOpts{
		ClosedFrom:     request.ClosedFrom,
		ClosedTo:       request.ClosedTo,
		AnomalyClasses: request.AnomalyClasses,
		CategoryIds:    request.CategoryIds,
		DefectIds:      request.DefectIds,
		Region:         request.Region,
		District:       request.District,
	})
	if err != nil {
		logger.Errorw(err.Error(), logFields...)
		return ctx.JSON(http.StatusInternalServerError, echo.Map{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
	}

	logger.Infow("", logFields...)

	return ctx.JSON(http.StatusOK, echo.Map{
		"code": http.StatusOK,
		"data": anomalyClassesCounterMap,
	})
}

func (c *ApplicationController) GetNormalAbnormalStats(ctx echo.Context) (err error) {
	var (
		logFields = []interface{}{"path", ctx.Path(), "method", ctx.Request().Method}
		request   dto.ApplicationRetrieveOpts
	)

	if err = ctx.Bind(&request); err != nil {
		logger.Warnw(err.Error(), logFields...)
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"code": http.StatusBadRequest,
			"msg":  err.Error(),
		})
	}
	if err = ctx.Validate(request); err != nil {
		logger.Warnw(err.Error(), logFields...)
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"code": http.StatusBadRequest,
			"msg":  err.Error(),
		})
	}

	normalAbnormalCountersMap, err := c.ApplicationViewService.CountNormalAbnormal(ctx.Request().Context(), &bo.ApplicationRetrieveOpts{
		ClosedFrom:     request.ClosedFrom,
		ClosedTo:       request.ClosedTo,
		AnomalyClasses: request.AnomalyClasses,
		CategoryIds:    request.CategoryIds,
		DefectIds:      request.DefectIds,
		Region:         request.Region,
		District:       request.District,
	})
	if err != nil {
		logger.Errorw(err.Error(), logFields...)
		return ctx.JSON(http.StatusInternalServerError, echo.Map{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
	}

	logger.Infow("", logFields...)

	return ctx.JSON(http.StatusOK, echo.Map{
		"code": http.StatusOK,
		"data": normalAbnormalCountersMap,
	})
}
