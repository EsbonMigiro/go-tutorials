package main

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var books = []book{
	{ID: "1", Title: "GoBook", Author: "Michael", Quantity: 2},
	{ID: "2", Title: "java", Author: "Dan", Quantity: 4},
	{ID: "3", Title: "js", Author: "Ben", Quantity: 6},
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func createBook(c *gin.Context) {
	var newBook book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}
	books = append(books, newBook)

	c.IndentedJSON(http.StatusCreated, newBook)
}
func bookById(c *gin.Context) {
	id := c.Param("id")

	book, err := getBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found."})
		return
	}
	c.IndentedJSON(http.StatusOK, book)
}

func getBookById(id string) (*book, error) {
	for i, b := range books {
		if b.ID == id {
			return &books[i], nil
		}

	}
	return nil, errors.New("Book not found")
}

func checkoutBook(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "bad request sent"})
		return
	}
	book, err := getBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
		return
	}
	if book.Quantity <= 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "no book found"})
		return
	}
	book.Quantity -= 1
	c.IndentedJSON(http.StatusOK, book)

}

func main() {
	err := OpenDatabase()

	if err != nil {

		log.Printf("error while connecting to the database: %v", err)
	}

	router := gin.Default()

	router.GET("/get-books", getBooks)
	router.POST("/create-books", createBook)
	router.GET("/get-book/:id", bookById)
	router.PATCH("/checkout", checkoutBook)

	router.Run("localhost:8080")
}
