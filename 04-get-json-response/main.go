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

func main() {
	app := fiber.New()

	// Open : http://localhost:3000/check/health
	app.Get("/check/health", jsonHandler)

	err := app.Listen(":3000")
	if err != nil {
		fmt.Println("[ERROR] Failed to start server")
	}
}
