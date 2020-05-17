package book

import (
	"github.com/gofiber/fiber"
	"github.com/harssh/books_gofiber/database"
	"github.com/jinzhu/gorm"
)

type Book struct {
	gorm.Model
	Title  string `json:"title"`
	Author string `json:"author"`
	Rating int    `json:"rating"`
}

func GetBooks(c *fiber.Ctx) {
	db := database.DBConn
	var books []Book
	db.Find(&books)
	c.JSON(books)
}

func GetBook(c *fiber.Ctx) {
	id := c.Params("id")
  db := database.DBConn
	var book Book
	db.Find(&book, id)
	c.JSON(book)
}

func UpdateBook(c *fiber.Ctx) {
	c.Send("A update Single Books")
}

func NewBook(c *fiber.Ctx) {
	db := database.DBConn
	var book Book
	book.Title = "1984"
	book.Author = "George Orwell"
	book.Rating = 5

	db.Create(&book)
	if db.NewRecord(book) != false {
		c.Send("Record created.")
		} else {
		c.Send("Creation Failed")
	}
 c.JSON(book)
}

func DeleteBook(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn

	var book Book
	db.Find(&book, id)
	if book.Title == "" {
		c.Status(500).Send("No book found with given ID")
		return
	}

	db.Delete(&book)
	c.Send("Book Successfully created.")
}
