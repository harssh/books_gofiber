package main

import (
	"github.com/gofiber/fiber"
	"github.com/harssh/books_gofiber/book"
	"github.com/harssh/books_gofiber/database"
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

func initDatabase()  {
	var
}

func main() {
	app := fiber.New()
	setupRoutes(app)
	app.Listen(3000)

}
