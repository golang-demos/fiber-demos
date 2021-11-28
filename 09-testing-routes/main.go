package main

import (
	"fmt"
	"io/ioutil"
	"net/http/httptest"

	"github.com/gofiber/fiber/v2"
)

func greetHandler(c *fiber.Ctx) error {
	name := c.Params("name")
	c.SendString("Hello " + name + "!")
	return nil
}

func postCreateHandler(c *fiber.Ctx) error {
	data := struct {
		Title string `json:"title"`
		Body  string `json:"body"`
	}{}
	c.BodyParser(&data)
	return c.JSON(data)
}

func main() {
	app := fiber.New()

	app.Get("/greet/:name", greetHandler)

	req := httptest.NewRequest("GET", "http://localhost:3000/greet/Vinay", nil)
	resp, _ := app.Test(req)

	fmt.Println("Status : ", resp.StatusCode)
	if resp.StatusCode == fiber.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(body))
	}

}
