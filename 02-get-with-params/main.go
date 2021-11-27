package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func greetHandler(c *fiber.Ctx) error {
	name := c.Params("name")
	c.SendString("Hello " + name + "!")
	return nil
}

func multipleParamsHandler(c *fiber.Ctx) error {
	postId := c.Params("postId")
	commentId := c.Params("commentId")
	c.SendString("Showing Comment with id : " + commentId + " for Post with Id : " + postId)
	return nil
}

func wildCardHandler(c *fiber.Ctx) error {
	anyParam := c.Params("*")
	c.SendString("Passed string : " + anyParam)
	return nil
}

func optionalParameterHandler(c *fiber.Ctx) error {
	todoId := c.Params("todoId")
	if todoId != "" {
		c.SendString("Show TODO Id : " + todoId)
	} else {
		c.SendString("Show all TODOs")
	}
	return nil
}

func main() {
	app := fiber.New()

	// Open : http://localhost:3000/greet/Vinay
	app.Get("/greet/:name", greetHandler)

	// Open : http://localhost:3000/posts/12/comments/415
	app.Get("/posts/:postId/comments/:commentId", multipleParamsHandler)

	// Open : http://localhost:3000/wildcard/any/parameter/type
	app.Get("/wildcard/*", wildCardHandler)

	// Open : http://localhost:3000/todos/
	// Open : http://localhost:3000/todos/1
	app.Get("/todos/:todoId?", optionalParameterHandler)

	err := app.Listen(":3000")
	if err != nil {
		fmt.Println("[ERROR] Failed to start server")
	}
}
