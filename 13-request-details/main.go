package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func detailsHandler(c *fiber.Ctx) error {
	data := struct {
		Hostname          string
		IpAddress         string
		ForwaredIpAddress []string
	}{}

	// hostname for URL
	data.Hostname = c.Hostname()

	// IP Address of the user visiting this route
	data.IpAddress = c.IP()

	// IP Addresses taken from X-Forwarded-For Header
	data.ForwaredIpAddress = c.IPs()

	return c.JSON(data)
}

func main() {

	app := fiber.New()

	app.Get("/my-details", detailsHandler)

	log.Fatal(app.Listen(":3000"))
}
