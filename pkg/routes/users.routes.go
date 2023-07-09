package routes

import (
	"fmt"
	"notes-app/pkg/db"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Usucorreo      string `validate:"required,email,min=4,max=32"`
	Usucontrasenia string `validate:"required,min=6,max=32"`
	Usutoken       string
}

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

var validate = validator.New()

func ValidateStruct(user User) []*ErrorResponse {
	var errors []*ErrorResponse
	err := validate.Struct(user)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}

func CreateUser(c *fiber.Ctx) error {

	user := new(User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Datos invalidos")
	}

	errors := ValidateStruct(*user)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(user.Usucontrasenia), 14)

	if err != nil {
		fmt.Println("Ha ocurrido un error al generar la contrasenia")
	}

	user.Usucontrasenia = string(bytes)

	result := db.DB.Create(&user)

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error al crear el registro")
	} else {
		return c.Status(fiber.StatusOK).SendString("Usuario creado")
	}

}
func GetUsers(c *fiber.Ctx) error {

	return c.SendString("Get all users")
}
