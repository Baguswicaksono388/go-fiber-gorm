package controllers

import (
	"go-fiber-gorm/database"
	"go-fiber-gorm/model/entity"
	"go-fiber-gorm/model/request"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func UserControllerGetAll(ctx *fiber.Ctx) error {
	var users []entity.User
	result := database.DB.Debug().Find(&users)

	if result.Error != nil {
		log.Println(result.Error)
	}

	return ctx.JSON(fiber.Map{
		"message": "Success get all user",
		"data":    users,
	})
}

func UserControllerCreate(ctx *fiber.Ctx) error {
	user := new(request.UserCreateRequest)

	if err := ctx.BodyParser(&user); err != nil {
		return err
	}

	validate := validator.New()
	errValidate := validate.Struct(user)
	if errValidate != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Validation error",
			"data":    errValidate.Error(),
		})
	}

	newUser := entity.User{
		Name:    user.Name,
		Email:   user.Email,
		Address: user.Address,
		Phone:   user.Phone,
	}

	errorCreateUser := database.DB.Debug().Create(&newUser).Error

	if errorCreateUser != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "Error create user",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "Success create user",
		"data":    newUser,
	})
}
