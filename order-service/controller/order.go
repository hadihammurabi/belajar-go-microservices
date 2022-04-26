package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hadihammurabi/belajar-go-microservices/order-service/config"
	"github.com/hadihammurabi/belajar-go-microservices/order-service/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type OrderController struct {
}

func (c *OrderController) getCollection() *mongo.Collection {
	return config.DB.Collection("orders")
}

func (c *OrderController) GetAllOrders(ctx *fiber.Ctx) error {
	curr, err := c.getCollection().Find(ctx.Context(), bson.D{})
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"error": err,
		})
	}
	defer curr.Close(ctx.Context())

	orders := make([]model.Order, 0)
	if err := curr.All(ctx.Context(), &orders); err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"error": err,
		})
	}

	return ctx.JSON(fiber.Map{
		"data": orders,
	})
}

func (c *OrderController) CreateOrder(ctx *fiber.Ctx) error {
	var orderIn *model.Order
	if err := ctx.BodyParser(&orderIn); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"error": err,
		})
	}

	result, err := c.getCollection().InsertOne(ctx.Context(), orderIn)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"error": err,
		})
	}

	return ctx.JSON(fiber.Map{
		"data": result,
	})
}

func (c *OrderController) GetById(ctx *fiber.Ctx) error {
	id := ctx.Params("id", "")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"error": err,
		})
	}

	result := c.getCollection().FindOne(ctx.Context(), bson.M{
		"_id": objectID,
	})
	err = result.Err()
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"error": err,
		})
	}

	var order *model.Order
	err = result.Decode(&order)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"error": err,
		})
	}

	return ctx.JSON(fiber.Map{
		"data": order,
	})
}
