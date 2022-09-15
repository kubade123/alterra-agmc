package main

import (
	"net/http"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/Status",func (e echo.Context)  error {
		return e.JSON(http.StatusOK,map[string]string("Status": "OK"))
	})
}