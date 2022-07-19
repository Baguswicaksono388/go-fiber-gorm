package controllers

import "github.com/gofiber/fiber/v2"

func UserControllerRead(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{
		"data": "user",
	})
}
