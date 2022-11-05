package routes

import (
	"github.com/gofiber/fiber/v2"
	"iecare-api/src/app/modules/accounts/http/controllers"
	"iecare-api/src/app/shared/http/middlewares"
)

func FileRoutes(app *fiber.App) {
	api := app.Group("/files")

	api.Use(middlewares.Acl([]string{"admin", "root", "user"}))

	api.Post("/", controllers.Store).Name("files.store")
	api.Static("/uploads", "public/uploads/")
}
