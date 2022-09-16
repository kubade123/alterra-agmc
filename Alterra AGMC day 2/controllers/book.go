package controllers

import (
	"day2/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

var (
	listData []models.Book
)

func init() {
	listData = []models.Book{}
}

func ShowBooks(e echo.Context) error {
	//book := models.Book{ID: 1, Title: "REST API Tutorial"}
	return e.JSON(http.StatusOK, listData)
}

// func CreateBooks(e echo.Context) error {
// 	id := e.FormValue("id")
// 	title := e.FormValue("title")

// 	var book models.Book
// 	book.Id = id
// 	book.Title = title

// 	return e.JSON(http.StatusOK, map[string]interface{}{
// 		"messages": "success create book",
// 		"book":     book,
// 	})
// }

func CreateBooks(bookBaru models.Book) models.Book {
	if len(listData) == 0 {
		bookBaru.ID = 1
	} else {
		bookBaru.ID = listData[len(listData)-1].ID + 1
	}
	listData = append(listData, bookBaru)

	return bookBaru
}

func PostBooks(e echo.Context) error {
	var dataBaru models.Book
	err := e.Bind(&dataBaru)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Error dari server",
			"status":  false,
		})
	}
	data := CreateBooks(dataBaru)

	return e.JSON(http.StatusCreated, map[string]interface{}{
		"message": "Success",
		"status":  true,
		"data":    data,
	})
}
