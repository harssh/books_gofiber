package main

import (
	"fmt"

	"github.com/gofiber/fiber"
	"github.com/harssh/books_gofiber/book"
	"github.com/harssh/books_gofiber/database"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func helloWorld(c *fiber.Ctx) {
	c.Send("Hello, World!")
}

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/book/:id", book.GetBook)
	app.Get("/api/v1/books", book.GetBooks)
	app.Put("/api/v1/book/:id", book.UpdateBook)
	app.Post("/api/v1/book", book.NewBook)
	app.Delete("api/v1/book/:id", book.DeleteBook)

}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "books.db")
	if err != nil {
		panic(err)
	}
	fmt.Println("DB connected.")
	database.DBConn.AutoMigrate(&book.Book{})
	fmt.Println("DB migrated")
}

func main() {
	app := fiber.New()
	initDatabase()

	defer database.DBConn.Close()
	setupRoutes(app)
	app.Listen(3000)

}
