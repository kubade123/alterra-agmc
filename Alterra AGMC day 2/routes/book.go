package routes

import (
	"project/controllers"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()
	e.GET("/Books", controllers.ShowBooks)
	e.POST("/Books", CreateBooks)
}
