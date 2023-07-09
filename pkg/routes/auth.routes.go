package routes

import (
	"notes-app/pkg/db"

	"github.com/gofiber/fiber/v2"
)

type UserLogin struct {
	Usucorreo      string
	Usucontrasenia string
}

func LoginUser(c *fiber.Ctx) error {

	user := new(UserLogin)

	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Datos invalidos")
	}

	usuo := db.DB.Where("usucorreo = ? ", user.Usucorreo).First(&user)

	return c.JSON(usuo)

}
