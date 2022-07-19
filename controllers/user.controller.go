package controllers

import (
	"go-fiber-gorm/database"
	"go-fiber-gorm/model/entity"
	"log"

	"github.com/gofiber/fiber/v2"
)

func UserControllerGetAll(ctx *fiber.Ctx) error {
	var users []entity.User
	result := database.DB.Find(&users)

	if result.Error != nil {
		log.Println(result.Error)
	}

	return ctx.JSON(users)
}
