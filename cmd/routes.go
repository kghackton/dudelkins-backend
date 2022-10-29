package main

import (
	"dudelkins/infrastructure"
	"dudelkins/pkg/logger"
	"sync"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func initRouter(k *infrastructure.Kernel, port string) (routerStartFunc func(wg *sync.WaitGroup) (err error)) {
	router := echo.New()
	router.Use(middleware.Recover())
	router.HTTPErrorHandler = func(err error, c echo.Context) {
		logger.Error(c.Request().Context(), "err: %s, request: %+v", err.Error(), c.Request())

		router.DefaultHTTPErrorHandler(err, c)
	}

	registerRoutes(router, k)

	return func(wg *sync.WaitGroup) (err error) {
		defer wg.Done()

		if err = router.Start(":" + port); err != nil {
			return err
		}

		return
	}
}

func registerRoutes(router *echo.Echo, injector infrastructure.IInjector) {

}
