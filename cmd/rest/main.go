package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/nandonyata/Stray-Fedding/config"
	"github.com/nandonyata/Stray-Fedding/controller"
	"github.com/nandonyata/Stray-Fedding/exception"
	"github.com/nandonyata/Stray-Fedding/repository"
	"github.com/nandonyata/Stray-Fedding/service"
)

const PORT = ":3000"

func main() {
	// Setup Database
	database := config.NewMongoDatabase()

	// Setup Repository
	userRepository := repository.NewUserRepository(database)

	// Setup Service
	userService := service.NewUserService(&userRepository)

	// Setup Controller
	userController := controller.NewUserController(&userService)

	// Setup Fiber
	app := fiber.New(config.NewFiberConfig())
	app.Use(recover.New())

	// Setup Routing
	userController.Route(app)

	err := app.Listen(PORT)
	exception.PanicIfNeeded(err)
}
