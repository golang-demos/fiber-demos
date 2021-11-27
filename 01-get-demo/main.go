package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func indexHandler(c *fiber.Ctx) error {
	c.SendString("Hello World")
	return nil
}

func main() {
	app := fiber.New()

	// Open : http://localhost:3000/
	app.Get("/", indexHandler)

	err := app.Listen(":3000")
	if err != nil {
		fmt.Println("[ERROR] Failed to start server")
	}
}
