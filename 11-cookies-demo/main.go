package main

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

var CookieName string = "test-cookie-name"

func showCookieHandler(c *fiber.Ctx) error {
	cookieValue := c.Cookies(CookieName)
	c.SendString("Cookie Value : " + cookieValue)
	return nil
}
func addCookieHandler(c *fiber.Ctx) error {
	data := struct {
		Value string `json:"value"`
	}{}
	c.BodyParser(&data)
	c.Cookie(&fiber.Cookie{
		Name:     CookieName,
		Value:    data.Value,
		Expires:  time.Now().Add(120 * time.Minute),
		HTTPOnly: true,
	})
	return c.SendString("Cookie Created")
}
func deleteCookieHandler(c *fiber.Ctx) error {
	c.ClearCookie(CookieName)
	return c.SendString("Done")
}
func updateCookieHandler(c *fiber.Ctx) error {
	data := struct {
		Value string `json:"value"`
	}{}
	c.BodyParser(&data)
	c.Cookie(&fiber.Cookie{
		Name:     CookieName,
		Value:    data.Value,
		Expires:  time.Now().Add(120 * time.Minute),
		HTTPOnly: true,
	})
	return c.SendString("Cookie Updated")
}

func main() {
	app := fiber.New()

	app.Get("/cookie", showCookieHandler)
	app.Post("/cookie", addCookieHandler)
	app.Delete("/cookie", deleteCookieHandler)
	app.Put("/cookie", updateCookieHandler)

	log.Fatal(app.Listen(":3000"))
}
