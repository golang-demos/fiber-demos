package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func downloadHandler(c *fiber.Ctx) error {
	return c.Download("./resources/sunrise.jpg", "sunrise.jpg")
}

func main() {
	app := fiber.New()

	app.Get("/download", downloadHandler)

	log.Fatal(app.Listen(":3000"))
}
