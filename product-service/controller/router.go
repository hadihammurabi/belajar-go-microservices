package controller

import "github.com/gofiber/fiber/v2"

func route(app *fiber.App) {
	productController := ProductController{}
	products := app.Group("/products")
	products.Get("", productController.GetAllProducts)
	products.Post("", productController.CreateProduct)
	products.Get("/:id", productController.GetById)
}

func Init(app *fiber.App) {
	route(app)
}
