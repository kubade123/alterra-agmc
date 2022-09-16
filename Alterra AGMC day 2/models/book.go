package models

type Book struct {
	Id    string `json:"id" form:"id"`
	Title string `json:"title" form:"title"`
}
