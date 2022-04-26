package controller

import "github.com/gofiber/fiber/v2"

func route(app *fiber.App) {
	orderController := OrderController{}
	orders := app.Group("/orders")
	orders.Get("", orderController.GetAllOrders)
	orders.Post("", orderController.CreateOrder)
	orders.Get("/:id", orderController.GetById)
}

func Init(app *fiber.App) {
	route(app)
}
