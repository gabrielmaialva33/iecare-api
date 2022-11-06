package routes

import (
	"github.com/gofiber/fiber/v2"
	"iecare-api/src/app/http/controllers"
	"iecare-api/src/app/http/middlewares"
)

func ProviderRoutes(app *fiber.App, controller *controllers.ProvidersController) {
	api := app.Group("/providers")

	api.Use(middlewares.Acl([]string{"admin", "root", "user", "provider"}))

	api.Get("/", controller.List).Name("providers.list")
	api.Get("/:providerId", controller.Get).Name("providers.get")

	api.Use(middlewares.Acl([]string{"admin", "root", "provider"}))

	api.Post("/", controller.Store).Name("providers.store")
	api.Put("/:providerId", controller.Edit).Name("providers.edit")
	api.Delete("/:providerId", controller.Delete).Name("providers.delete")
}
