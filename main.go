package main

import (
	"andre/notesnotes-api/handler"
	"andre/notesnotes-api/model"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	//db.InitializeDb()

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&model.Book{})

	handler := handler.NewHandler(db)

	r := gin.New()

	r.GET("/books", handler.ListBooksHandler)
	r.POST("/books", handler.CreateBookHandler)
	r.DELETE("/books/:id", handler.RemoveBookHandler)

	r.Run()
}
