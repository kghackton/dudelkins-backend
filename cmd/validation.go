package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type CustomValidator struct {
	Validator validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}

func registerValidation(router *echo.Echo) {
	vld := *validator.New()
	router.Validator = &CustomValidator{vld}
	return
}
