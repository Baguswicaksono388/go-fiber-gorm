package main

import (
	"go-fiber-gorm/database"
	"go-fiber-gorm/database/migration"
	"go-fiber-gorm/route"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Initialize Database
	database.DatabaseInit()

	// Run Migration
	migration.RunMigration()

	app := fiber.New()

	// Initialize the route
	route.RouteInit(app)

	app.Listen(":3000")
}
