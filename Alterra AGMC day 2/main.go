package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Book struct {
	Id    string `json:"id" form:"id"`
	Title string `json:"title" form:"title"`
}

func ShowStatus(e echo.Context) error {
	return e.JSON(http.StatusOK, map[string]string{"Status": "OK"})
}

func ShowBooks(e echo.Context) error {
	book := Book{Id: "1", Title: "REST API Tutorial"}
	return e.JSON(http.StatusOK, book)
}

func CreateBooks(e echo.Context) error {
	id := e.FormValue("id")
	title := e.FormValue("title")

	var book Book
	book.Id = id
	book.Title = title

	return e.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success create book",
		"book":     book,
	})
}

func main() {
	e := echo.New()
	e.GET("/Status", ShowStatus)
	e.GET("/Books", ShowBooks)
	e.POST("/Books", CreateBooks)
	e.Start(":8080")
}
