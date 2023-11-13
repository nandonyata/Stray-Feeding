package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/nandonyata/Stray-Fedding/exception"
	"github.com/nandonyata/Stray-Fedding/model"
	"github.com/nandonyata/Stray-Fedding/service"
)

func NewUserController(servis *service.UserService) UserController {
	return UserController{
		UserService: *servis,
	}
}

type UserController struct {
	UserService service.UserService
}

func (controller *UserController) Route(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("AAA")
	})

	app.Post("/api/register", controller.Create)
}

func (controller *UserController) Create(c *fiber.Ctx) error {
	var req model.CreateUserRequest

	err := c.BodyParser(&req)
	req.Id = uuid.New().String()

	exception.PanicIfNeeded(err)

	resp := controller.UserService.Create(req)

	return c.JSON(model.WebResponse{
		Code:   201,
		Status: "OK",
		Data:   resp,
	})
}
