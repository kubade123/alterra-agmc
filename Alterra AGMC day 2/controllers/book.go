package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func ShowBooks(e echo.Context) error {
	book := models.Book{Id: "1", Title: "REST API Tutorial"}
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
