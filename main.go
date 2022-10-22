package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity string `json:"quantity"`
}

var books = []book{
	{ID: "1", Title: "Buku pertama", Author: "Author pertama", Quantity: "1"},
	{ID: "2", Title: "Buku kedua", Author: "Author kedua", Quantity: "2"},
	{ID: "3", Title: "Buku ketiga", Author: "Author ketiga", Quantity: "3"},
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}
func main() {
	router := gin.Default()

	router.GET("/books", getBooks)

	router.Run("localhost:8080")
}
