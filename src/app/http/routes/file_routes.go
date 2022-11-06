package routes

import (
	"github.com/gofiber/fiber/v2"
	"iecare-api/src/app/http/controllers"
	"iecare-api/src/app/http/middlewares"
)

func FileRoutes(app *fiber.App) {
	app.Static("/files/uploads", "public/uploads/").Name("uploads")

	api := app.Group("/files")

	api.Use(middlewares.Acl([]string{"admin", "root", "user", "provider"}))
	api.Post("/", controllers.Store).Name("files.store")
}
