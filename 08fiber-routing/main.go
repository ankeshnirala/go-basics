package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func handleRoute(c *fiber.Ctx) error {
	fmt.Println(c)
	return c.SendString("Welcome to the GOPHERS world")
}

func main() {
	app := fiber.New()

	app.Get("/", handleRoute)
	app.Get("/:value", func(c *fiber.Ctx) error {
		return c.SendString(c.Params("value"))
	})

	app.Listen(":3000")
}
