package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Open : http://localhost:3000/index.html
	pwd, _ := os.Getwd()
	rootPath := pwd + "/public"
	fmt.Println("Root Path : ", rootPath)
	app.Static("/", rootPath)

	err := app.Listen(":3000")
	if err != nil {
		fmt.Println("[ERROR] Failed to start server")
	}
}
