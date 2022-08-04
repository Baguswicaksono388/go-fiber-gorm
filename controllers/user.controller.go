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

func UserControllerGetById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var user entity.User
	result := database.DB.Debug().Where("id = ?", id).First(&user)

	if result.Error != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	// userResponse := response.UserResponse{
	// 	ID:        user.ID,
	// 	Email:     user.Email,
	// 	Name:      user.Name,
	// 	Address:   user.Address,
	// 	Phone:     user.Phone,
	// 	CreatedAt: user.CreatedAt,
	// 	UpdatedAt: user.UpdatedAt,
	// }

	return ctx.JSON(fiber.Map{
		"message": "Success get user by id",
		"data":    user,
	})
}

func UserControllerUpdate(ctx *fiber.Ctx) error {
	userUpdate := new(request.UserUpdateRequest)
	var user entity.User

	if err := ctx.BodyParser(&userUpdate); err != nil {
		return err
	}

	validate := validator.New()
	errValidate := validate.Struct(userUpdate)
	if errValidate != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Validation error",
			"data":    errValidate.Error(),
		})
	}

	// Check if user exist
	id := ctx.Params("id")
	result := database.DB.Debug().Where("id = ?", id).First(&user)
	if result.Error != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	// Update Data User
	if userUpdate.Name != "" || userUpdate.Email != "" || userUpdate.Address != "" || userUpdate.Phone != "" {
		user.Name = userUpdate.Name
		user.Email = userUpdate.Email
		user.Address = userUpdate.Address
		user.Phone = userUpdate.Phone
	}

	errorUpdateUser := database.DB.Debug().Save(&user).Error
	if errorUpdateUser != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "Error update user",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "Success update user",
		"data":    user,
	})

}
