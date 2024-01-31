package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/s809616134/hotel-backend/types"
)

func HandleGetUsers(c *fiber.Ctx) error {
	u := types.User{
		FirstName: "jasdfas",
	}
	return c.JSON(u)
}

func HandleGetUser(c *fiber.Ctx) error {
	return c.JSON("James2")
}
