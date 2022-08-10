package route

import (
	"go-fiber-gorm/controllers"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(r *fiber.App) {
	r.Get("/user", controllers.UserControllerGetAll)
	r.Post("/user", controllers.UserControllerCreate)
	r.Get("/user/:id", controllers.UserControllerGetById)
	r.Put("/user/:id", controllers.UserControllerUpdate)
	r.Delete("/user/:id", controllers.UserControllerDelete)
}
