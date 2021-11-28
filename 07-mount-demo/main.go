package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func postsHandler(c *fiber.Ctx) error {
	c.SendString("Show All Posts")
	return nil
}

func postsDetailsHandler(c *fiber.Ctx) error {
	c.SendString("Details for post with ID " + c.Params("id"))
	return nil
}

func main() {

	postsApp := fiber.New()

	postsApp.Get("/posts", postsHandler)
	postsApp.Get("/posts/:id", postsDetailsHandler)

	app := fiber.New()

	app.Mount("/api/v1", postsApp) // "/api/v1" and "/posts" get attached

	err := app.Listen(":3000")
	if err != nil {
		fmt.Println("Failed to start server")
	}
}
