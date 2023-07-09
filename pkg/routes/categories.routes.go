package routes

import "github.com/gofiber/fiber/v2"

func CreateCategory(c *fiber.Ctx) error {

	return c.SendString("Create category")
}

func GetCategories(c *fiber.Ctx) error {
	return c.SendString("Get categories")
}

func DeleteCategory(c *fiber.Ctx) error {
	return c.SendString("Delete Category")
}
