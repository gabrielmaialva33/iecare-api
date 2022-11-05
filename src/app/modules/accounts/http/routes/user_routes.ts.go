package routes

import (
	"github.com/gofiber/fiber/v2"
	"iecare-api/src/app/modules/accounts/http/controllers"
	"iecare-api/src/app/shared/http/middlewares"
)

func UserRoutes(app *fiber.App, controller *controllers.UsersController) {
	app.Post("/sign_in", controller.SignIn).Name("sign_in")
	app.Post("/sign_up", controller.SignUp).Name("sign_up")

	api := app.Group("/users")

	api.Use(middlewares.Acl([]string{"admin", "root", "user"}))

	api.Get("/", controller.List).Name("users.list")
	api.Get("/:userId", controller.Get).Name("users.get")
	api.Post("/", controller.Store).Name("users.store").Name("users.store")
	api.Put("/:userId", controller.Edit).Name("users.edit")
	api.Delete("/:userId", controller.Delete).Name("users.delete")
}
