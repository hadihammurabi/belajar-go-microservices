package controller

import "github.com/gofiber/fiber/v2"

func route(app *fiber.App) {
	userController := UserController{}
	users := app.Group("/users")
	users.Get("", userController.GetAllUsers)
	users.Post("", userController.Createuser)
	users.Get("/:id", userController.GetById)
}

func Init(app *fiber.App) {
	route(app)
}
