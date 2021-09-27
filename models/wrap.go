package models

import (
	"github.com/labstack/echo/v4"
)

type Error struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

type JSONResponseData struct {
	Code    int64       `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

type JSONResponse struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

func (j JSONResponseData) Response(c echo.Context) error {
	return c.JSON(int(j.Code), j)
}

func (j JSONResponse) Response(c echo.Context) error {
	return c.JSON(int(j.Code), j)
}
