package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

type Post struct {
	Title string `json:"title" xml:"title" form:"title"`
	Body  string `json:"body" xml:"body" form:"body"`
}

func postHandler(c *fiber.Ctx) error {
	var data Post
	c.BodyParser(&data)

	fmt.Println("\n\"" + c.Get("Content-Type") + "\" Body : \n" + string(c.Body()))
	c.JSON(data)
	return nil
}

func main() {
	app := fiber.New()

	app.Post("/post/new", postHandler)

	log.Fatal(app.Listen(":3000"))
}

// curl -X POST -H "Content-Type: application/json" --data "{\"title\":\"john\",\"body\":\"doe\"}" localhost:3000/post/new
// curl -X POST -H "Content-Type: application/xml" --data "<login><title>john</title><body>doe</body></login>" localhost:3000/post/new
// curl -X POST -H "Content-Type: application/x-www-form-urlencoded" --data "title=john&body=doe" localhost:3000/post/new
// curl -X POST -F title=john -F body=doe http://localhost:3000/post/new
