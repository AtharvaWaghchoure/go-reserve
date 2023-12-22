package api

import (
	"github.com/AtharvaWaghchoure/goreserve/types"
	"github.com/gofiber/fiber/v2"
)

func HandlerGetUsers(c *fiber.Ctx) error {
	user := types.User{
		FirstName: "Stain",
		LastName:  "Cleaner",
	}
	return c.JSON(user)
}

func HandlerGetUser(c *fiber.Ctx) error {
	return c.JSON("GetUser handler")
}
