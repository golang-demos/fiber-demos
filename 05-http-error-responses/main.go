package main

import "github.com/gofiber/fiber/v2"

func notFoundHandler(c *fiber.Ctx) error {
	return fiber.NewError(fiber.StatusNotFound, "Not Found")
}

func internalServerErrorHandler(c *fiber.Ctx) error {
	return fiber.NewError(fiber.StatusInternalServerError, "Internal Server Error")
}

func movedPermanentlyHandler(c *fiber.Ctx) error {
	return fiber.NewError(fiber.StatusMovedPermanently, "Moved Permanently")
}

func main() {
	app := fiber.New()

	// Visit http://localhost:3000/404
	app.Get("/404", notFoundHandler)

	// Visit http://localhost:3000/500
	app.Get("/500", internalServerErrorHandler)

	// Visit http://localhost:3000/301
	app.Get("/301", movedPermanentlyHandler)

	app.Listen(":3000")
}
