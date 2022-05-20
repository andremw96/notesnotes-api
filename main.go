package main

import (
	"andre/notesnotes-api/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	r.GET("/books", listBooksHandler)
	r.POST("/books", createBookHandler)
	r.DELETE("/books/:id", removeBookHandler)

	r.Run()
}

func listBooksHandler(c *gin.Context) {
	c.JSON(http.StatusOK, model.Books)
}

func createBookHandler(c *gin.Context) {
	var book model.Book

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	model.Books = append(model.Books, book)
	c.JSON(http.StatusCreated, book)
}

func removeBookHandler(c *gin.Context) {
	id := c.Param("id")
	for i, a := range model.Books {
		if a.ID == id {
			model.Books = append(model.Books[:i], model.Books[i+1:]...)
			break
		}
	}
	c.Status(http.StatusNoContent)
}

// r.GET("/", func(c *gin.Context) {
// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "Hello World",
// 	})
// })
