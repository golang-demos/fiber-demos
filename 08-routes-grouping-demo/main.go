package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func apiGroupHandler(c *fiber.Ctx) error {
	accessToken := c.Cookies("Access-Token")
	if accessToken != "" { // Do not allow if Access Token is not set in cookie
		return c.Next()
	} else {
		log.Println("Access token is missing. You need to set a cookie with name \"Access-Token\"")
	}
	return nil
}

func postsGroupHandler(c *fiber.Ctx) error {
	fmt.Println("In Posts groups Handler")
	return c.Next() // This statement is mandatory to execute other routes registered in this group.
}

func postDetailsHandler(c *fiber.Ctx) error {
	fmt.Println("In Posts Details Handler")
	return nil
}
func postDeleteHandler(c *fiber.Ctx) error {
	fmt.Println("In Posts delete Handler")
	return nil
}

func main() {
	app := fiber.New()

	api := app.Group("/api", apiGroupHandler)

	postsApi := api.Group("/posts", postsGroupHandler)

	postsApi.Get("/:id", postDetailsHandler)
	postsApi.Delete("/:id", postDeleteHandler)

	err := app.Listen(":3000")
	if err != nil {
		fmt.Println("Failed to start server")
	}
}
