package main

import (
	"sync"

	"dudelkins/infrastructure"
	"dudelkins/pkg/logger"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func initRouter(k *infrastructure.Kernel, port string) (routerStartFunc func(wg *sync.WaitGroup) (err error)) {
	router := echo.New()
	router.Use(middleware.Recover())
	router.HTTPErrorHandler = func(err error, c echo.Context) {
		logger.Errorf("err: %s, request: %+v", err.Error(), c.Request())

		router.DefaultHTTPErrorHandler(err, c)
	}

	registerRoutes(router, k)
	registerValidation(router)

	return func(wg *sync.WaitGroup) (err error) {
		defer wg.Done()

		if err = router.Start(":" + port); err != nil {
			return err
		}

		return
	}
}

func registerRoutes(router *echo.Echo, injector infrastructure.IInjector) {
	applicationController := injector.InjectApplicationController()

	api := router.Group("/api")
	{
		applications := api.Group("/applications")
		{
			applications.POST("", applicationController.Create)
			applications.GET("", applicationController.Get)

			stats := applications.Group("/stats")
			{
				stats.GET("/anomalyClasses", applicationController.GetAnomalyClassesStats)
				stats.GET("/normalAbnormal", applicationController.GetNormalAbnormalStats)
			}
		}
	}
}
