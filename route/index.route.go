package route

import (
	"go-fiber-gorm/config"
	"go-fiber-gorm/controllers"
	"go-fiber-gorm/middleware"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(r *fiber.App) {
	r.Static("/public", config.ProjectRootPath+"/public/asset")
	r.Post("/login", controllers.LoginController)
	r.Get("/user", middleware.UserMiddleware, controllers.UserControllerGetAll)
	r.Post("/user", controllers.UserControllerCreate)
	r.Get("/user/:id", controllers.UserControllerGetById)
	r.Put("/user/:id", controllers.UserControllerUpdate)
	r.Delete("/user/:id", controllers.UserControllerDelete)
}
