package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func detailsHandler(c *fiber.Ctx) error {
	data := struct {
		Hostname          string
		Protocol          string
		IpAddress         string
		ForwaredIpAddress []string
		Subdomains        []string
		URL               string
		Path              string
	}{}

	// hostname for URL
	data.Hostname = c.Hostname()

	data.Protocol = c.Protocol()

	// IP Address of the user visiting this route
	data.IpAddress = c.IP()

	// IP Addresses taken from X-Forwarded-For Header
	data.ForwaredIpAddress = c.IPs()

	// original request URL
	data.URL = c.OriginalURL()

	data.Path = c.Path()

	data.Subdomains = c.Subdomains()

	return c.JSON(data)
}

func main() {

	app := fiber.New()

	app.Get("/my-details", detailsHandler)

	log.Fatal(app.Listen(":3000"))
}
