package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/hadihammurabi/belajar-go-microservices/product-service/config"
	"github.com/hadihammurabi/belajar-go-microservices/product-service/controller"
)

func main() {
	if err := config.ConfigureDatabase(); err != nil {
		panic(err)
	}

	app := fiber.New()
	controller.Init(app)
	log.Fatal(app.Listen(":8001"))
}
