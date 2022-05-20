package model

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var Books = []Book{
	{ID: "1", Title: "Harry Potter", Author: "J. K. ROWLING"},
	{ID: "2", Title: "The Lord of The Rings", Author: "J. R. R. Tolkien"},
	{ID: "3", Title: "The Wizard of Oz", Author: "L. Frank Baum"},
}
