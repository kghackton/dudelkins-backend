package controllers

import (
	"context"
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
			RowsSkip   int `json:"rowsSkip" validate:"gte=0"`
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

	if request.RowsAmount == 0 {
		request.RowsAmount = 300
	}
	// to skip csv header
	if request.RowsSkip == 0 {
		request.RowsSkip = 1
	}

	go func() {
		csvFile, err := os.Open("/data/applications.csv")
		if err != nil {
			logger.Errorw(err.Error(), logFields...)
			return
		}
		defer csvFile.Close()

		csvReader := csv.NewReader(csvFile)
		csvReader.Comma = '$'
		csvReader.ReuseRecord = true

		for i := 0; i < request.RowsSkip; i++ {
			_, err := csvReader.Read()
			if err != nil {
				if err == io.EOF {
					logger.Debugf("rowsSkip greater than rows amount in csv")
					return
				}
				logger.Errorw(err.Error(), logFields...)
				continue
			}
		}

		workerPool := make(chan struct{}, 6)
		for i := 0; i < request.RowsAmount; i++ {
			record, err := csvReader.Read()
			if err != nil {
				if err == io.EOF {
					break
				}
				logger.Errorw(err.Error(), logFields...)
				continue
			}

			application, err := bo.NewApplicationFromRecord(record)
			if err != nil {
				logger.Errorw(err.Error(), logFields...)
				continue
			}

			workerPool <- struct{}{}
			go func() {
				defer func() {
					<-workerPool
				}()

				if err = c.ApplicationUploadService.Create(context.TODO(), application); err != nil {
					logger.Errorw(err.Error(), logFields...)
					return
				}
			}()

			if i%500 == 0 || i == request.RowsAmount-1 {
				logger.Debugf("%d", i)
			}
		}
	}()

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

	if request.Limit == nil {
		defaultLimit := 100
		request.Limit = &defaultLimit
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

func (c *ApplicationController) GetSingle(ctx echo.Context) (err error) {
	var (
		logFields = []interface{}{"path", ctx.Path(), "method", ctx.Request().Method}
		request   struct {
			RootId int `param:"rootId"`
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

	application, err := c.ApplicationViewService.GetSingle(ctx.Request().Context(), request.RootId)
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
		"data": application.ToDto(),
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

func (c *ApplicationController) GetAnomalyClassesWithCreationHourStats(ctx echo.Context) (err error) {
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

	anomalyClassCountersWithCreationHourMap, err := c.ApplicationViewService.CountAnomalyClassesWithCreationHour(ctx.Request().Context(), &bo.ApplicationRetrieveOpts{
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
		"data": anomalyClassCountersWithCreationHourMap,
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
