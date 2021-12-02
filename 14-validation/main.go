package main

import (
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type Exam struct {
	Name  string `json:"name" validate:"required,min=3,max=32"`
	Marks int    `json:"marks" validate:"required,number"`
}

type Student struct {
	Name       string `json:"name" validate:"required,min=3,max=32"`
	IsAssigned *bool  `json:"isassigned" validate:"required"`
	Email      string `json:"email" validate:"required,email,min=6,max=32"`
	Exam       Exam   `json:"exam" validate:"dive"`
}

type ErrorResp struct {
	Field string
	Tag   string
	Value string
}

func Validate(student Student) []*ErrorResp {
	var errors []*ErrorResp
	validate := validator.New()
	err := validate.Struct(student)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResp
			element.Field = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}

func AddStudentHandler(c *fiber.Ctx) error {
	student := new(Student)

	if err := c.BodyParser(student); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	errors := Validate(*student)
	if errors != nil {
		return c.JSON(errors)
	}

	return c.JSON(student)
}

func main() {
	app := fiber.New()

	app.Post("/register/student", AddStudentHandler)

	log.Fatal(app.Listen(":3000"))

	// Running a test with the following curl commands
	/*
		curl --location --request POST 'http://localhost:3000/register/student' --header 'Content-Type: application/json' --data-raw '{ "name":"Vi", "isassigned": true, "exam":{ "name": "S2" } }'
	*/
}
