package main

import (
	"crypto/tls"
	"log"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/acme/autocert"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("This is a secure server")
	})

	m := &autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist("jyotfinal.com"),
		Cache:      autocert.DirCache("./certs"),
	}

	// TLS Config
	cfg := &tls.Config{
		GetCertificate: m.GetCertificate,
		NextProtos: []string{
			"http/1.1", "acme-tls/1",
		},
	}

	ln, err := tls.Listen("tcp", ":443", cfg)
	if err != nil {
		panic(err)
	}

	log.Fatal(app.Listener(ln))
}
