package routes

import (
	"github.com/gofiber/fiber/v2"
	"iecare-api/src/app/http/controllers"
	"iecare-api/src/app/http/middlewares"
)

func CategoryRoutes(app *fiber.App, controller *controllers.CategoriesController) {
	api := app.Group("/categories")

	api.Use(middlewares.Acl([]string{"admin", "root", "user", "provider"}))

	api.Get("/", controller.List).Name("categories.list")
	api.Get("/:categoryId", controller.Get).Name("categories.get")

	api.Use(middlewares.Acl([]string{"admin", "root"}))

	api.Post("/", controller.Store).Name("categories.store")
	api.Put("/:categoryId", controller.Edit).Name("categories.edit")
	api.Delete("/:categoryId", controller.Delete).Name("categories.delete")
}
