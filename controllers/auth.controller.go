package controllers

import (
	"go-fiber-gorm/database"
	"go-fiber-gorm/helpers"
	"go-fiber-gorm/model/entity"
	"go-fiber-gorm/model/request"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func LoginController(ctx *fiber.Ctx) error {
	loginRequest := new(request.LoginRequest)

	if err := ctx.BodyParser(&loginRequest); err != nil {
		return err
	}

	validate := validator.New()
	errValidate := validate.Struct(loginRequest)
	if errValidate != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Validation error",
			"data":    errValidate.Error(),
		})
	}

	// Check Available User
	var user entity.User
	result := database.DB.Debug().Where("email = ?", loginRequest.Email).First(&user)

	if result.Error != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "wrong credential",
		})
	}

	// Check Validation Password
	isValid := helpers.CheckPasswordHash(loginRequest.Password, user.Password)

	if !isValid {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "wrong credential",
		})
	}

	return ctx.JSON(fiber.Map{
		"token": "secret",
	})
}
