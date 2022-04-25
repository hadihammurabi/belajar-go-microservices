package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hadihammurabi/belajar-go-microservices/product-service/config"
	"github.com/hadihammurabi/belajar-go-microservices/product-service/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductController struct {
}

func (c *ProductController) getCollection() *mongo.Collection {
	return config.DB.Collection("products")
}

func (c *ProductController) GetAllProducts(ctx *fiber.Ctx) error {
	curr, err := c.getCollection().Find(ctx.Context(), bson.D{})
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"error": err,
		})
	}
	defer curr.Close(ctx.Context())

	products := make([]model.Product, 0)
	if err := curr.All(ctx.Context(), &products); err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"error": err,
		})
	}

	return ctx.JSON(fiber.Map{
		"data": products,
	})
}

func (c *ProductController) CreateProduct(ctx *fiber.Ctx) error {
	var productIn *model.Product
	if err := ctx.BodyParser(&productIn); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"error": err,
		})
	}

	result, err := c.getCollection().InsertOne(ctx.Context(), productIn)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"error": err,
		})
	}

	return ctx.JSON(fiber.Map{
		"data": result,
	})
}
