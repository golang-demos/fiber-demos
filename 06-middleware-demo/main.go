package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func jsonHandler(c *fiber.Ctx) error {
	response := make(map[string]string)
	response["health"] = "1"
	return c.JSON(response)
}

// Middleware : Set JSON Content-Type header
func jsonContentTypeMiddleware(c *fiber.Ctx) error {
	c.Set("Content-Type", "application/json")
	return c.Next()
}

func main() {
	app := fiber.New()

	app.Use(jsonContentTypeMiddleware)

	app.Get("/check", jsonHandler)

	err := app.Listen(":3000")
	if err != nil {
		fmt.Println("Failed to start server")
	}
}
