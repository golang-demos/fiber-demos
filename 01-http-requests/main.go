package main

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func indexHandler(c *fiber.Ctx) error {
	c.SendString("Hello World")
	return nil
}

func postRequestHandler(c *fiber.Ctx) error {
	data := struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	}{}
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	c.JSON(data)
	return nil
}

func deletetRequestHandler(c *fiber.Ctx) error {
	postId := c.Params("postId")
	c.SendString("Deleting Post with Id : " + postId)
	return nil
}

func putRequestHandler(c *fiber.Ctx) error {
	postId := c.Params("postId")
	data := struct {
		PostId      int
		Title       string `json:"title"`
		Description string `json:"description"`
	}{}
	c.BodyParser(&data)
	data.PostId, _ = strconv.Atoi(postId)
	c.JSON(data)
	return nil
}

func main() {
	app := fiber.New()

	// Open : http://localhost:3000/
	app.Get("/", indexHandler)

	// Post Request to http://localhost:3000/post with data {value:1}
	app.Post("/post", postRequestHandler)

	// Delete Request to http://localhost:3000/post/105 with data {postId:1}
	app.Delete("/post/:postId", deletetRequestHandler)

	// Put Request to http://localhost:3000/post/105 with data {postId:1}
	app.Put("/post/:postId", putRequestHandler)

	err := app.Listen(":3000")
	if err != nil {
		fmt.Println("[ERROR] Failed to start server")
	}
}
