package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hadihammurabi/belajar-go-microservices/user-service/config"
	"github.com/hadihammurabi/belajar-go-microservices/user-service/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserController struct {
}

func (c *UserController) getCollection() *mongo.Collection {
	return config.DB.Collection("users")
}

func (c *UserController) GetAllUsers(ctx *fiber.Ctx) error {
	curr, err := c.getCollection().Find(ctx.Context(), bson.D{})
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"error": err,
		})
	}
	defer curr.Close(ctx.Context())

	users := make([]model.User, 0)
	if err := curr.All(ctx.Context(), &users); err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"error": err,
		})
	}

	return ctx.JSON(fiber.Map{
		"data": users,
	})
}

func (c *UserController) Createuser(ctx *fiber.Ctx) error {
	var userIn *model.User
	if err := ctx.BodyParser(&userIn); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"error": err,
		})
	}

	result, err := c.getCollection().InsertOne(ctx.Context(), userIn)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"error": err,
		})
	}

	return ctx.JSON(fiber.Map{
		"data": result,
	})
}

func (c *UserController) GetById(ctx *fiber.Ctx) error {
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

	var user *model.User
	err = result.Decode(&user)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"error": err,
		})
	}

	return ctx.JSON(fiber.Map{
		"data": user,
	})
}
