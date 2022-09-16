package routes

import (
	"day2/controllers"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()
	e.GET("/Books", controllers.ShowBooks)
	e.POST("/Books", controllers.PostBooks)

	return e
}
