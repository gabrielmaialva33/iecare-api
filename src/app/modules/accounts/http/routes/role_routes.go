package routes

import (
	"github.com/gofiber/fiber/v2"
	"iecare-api/src/app/modules/accounts/http/controllers"
	"iecare-api/src/app/shared/http/middlewares"
)

func RoleRoutes(app *fiber.App, controller *controllers.RolesController) {
	api := app.Group("/roles")

	api.Use(middlewares.Acl([]string{"admin", "root"}))

	api.Get("/", controller.List).Name("roles.list")
	api.Get("/:roleId", controller.Get).Name("roles.get")
	api.Post("/", controller.Store).Name("roles.store")
	api.Put("/:roleId", controller.Edit).Name("roles.edit")
	api.Delete("/:roleId", controller.Delete).Name("roles.delete")
}
