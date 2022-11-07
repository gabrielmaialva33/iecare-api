package routes

import (
	"github.com/gofiber/fiber/v2"
	"iecare-api/src/app/http/controllers"
	"iecare-api/src/app/http/middlewares"
)

func ServiceRoutes(app *fiber.App, controller *controllers.ServicesController) {
	api := app.Group("/services")

	api.Use(middlewares.Acl([]string{"admin", "root", "user", "provider"}))

	api.Get("/", controller.List).Name("services.list")
	api.Get("/:serviceId", controller.Get).Name("services.get")

	api.Use(middlewares.Acl([]string{"admin", "root", "provider"}))

	api.Post("/", controller.Store).Name("services.store")
	api.Put("/:serviceId", controller.Edit).Name("services.edit")
	api.Delete("/:serviceId", controller.Delete).Name("services.delete")
}
