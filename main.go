package main

import (
	"fmt"
	"notes-app/pkg/db"
	"notes-app/pkg/models"
	"notes-app/pkg/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {

	db.GetConnectionDB()
	db.DB.AutoMigrate(models.User{}, models.Category{})
	app := fiber.New()

	userGroup := app.Group("/users")
	userGroup.Post("/create", routes.CreateUser)
	userGroup.Get("", routes.GetUsers)

	categoryGroup := app.Group("/category")
	categoryGroup.Post("/create", routes.CreateCategory)
	categoryGroup.Get("/", routes.GetCategories)
	categoryGroup.Post("/delete", routes.DeleteCategory)

	authGroup := app.Group("/auth")
	authGroup.Post("/login", routes.LoginUser)

	app.Listen(":8002")

	fmt.Println("Database connected")

}
