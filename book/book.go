package book

import(
	"github.com/gofiber/fiber"
)

func GetBooks(c *fiber.Ctx)  {
	c.Send("All Books")
}

func GetBook(c *fiber.Ctx)  {
	c.Send("A Single Books")
}

func UpdateBook(c *fiber.Ctx)  {
	c.Send("A update Single Books")
}

func NewBook(c *fiber.Ctx)  {
	c.Send("Add new Book")
}

func DeleteBook(c *fiber.Ctx)  {
	c.Send("Delete a book")
}