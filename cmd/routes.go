package main

import (
	"crypto/subtle"
	"sync"

	"dudelkins/infrastructure"
	"dudelkins/pkg/logger"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func initRouter(k *infrastructure.Kernel, port string) (routerStartFunc func(wg *sync.WaitGroup) (err error)) {
	router := echo.New()
	router.Use(middleware.Recover(), middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		// Be careful to use constant time comparison to prevent timing attacks
		if subtle.ConstantTimeCompare([]byte(username), []byte("hi")) == 1 &&
			subtle.ConstantTimeCompare([]byte(password), []byte("test")) == 1 {
			return true, nil
		}
		return false, nil
	}))
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

			rootId := applications.Group("/:rootId")
			rootId.GET("", applicationController.GetSingle)

			stats := applications.Group("/stats")
			{
				stats.GET("/anomalyClasses", applicationController.GetAnomalyClassesStats)
				stats.GET("/anomalyClassesHour", applicationController.GetAnomalyClassesWithCreationHourStats)
				stats.GET("/normalAbnormal", applicationController.GetNormalAbnormalStats)
			}
		}
	}
}
