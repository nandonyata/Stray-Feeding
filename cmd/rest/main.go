package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nandonyata/Stray-Fedding/config"
)

const PORT = ":3000"

func main() {
	app := fiber.New()
	config.NewMongoDatabase()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	app.Post("/register", func(c *fiber.Ctx) error {
		return c.SendString("This is register")
	})

	app.Listen(PORT)
}
