package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"iecare-api/src/app/modules/accounts/http/controllers"
	"iecare-api/src/app/modules/accounts/http/routes"
	"iecare-api/src/database"
	"os"
)

func main() {
	godotenv.Load()

	dsn := os.Getenv("DATABASE_URL")
	services := database.Connect(dsn)

	services.Drop()
	services.Migrate()
	services.Seed()

	app := fiber.New(fiber.Config{
		AppName:                 "Base Fiber API",
		EnableTrustedProxyCheck: true,
		TrustedProxies:          []string{"0.0.0.0"},
		ProxyHeader:             fiber.HeaderXForwardedFor,
		BodyLimit:               10 * 1024 * 1024, // 10 MB
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization, X-Request-With",
		AllowCredentials: true,
	}))
	app.Use(logger.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message":  "Welcome to theIE Care API",
			"status":   c.Response().StatusCode(),
			"database": services.Stats(),
		})
	})

	// Controllers
	userController := controllers.NewUsersController(services.User)
	roleController := controllers.NewRolesController(services.Role)

	// Routes
	routes.UserRoutes(app, userController)
	routes.RoleRoutes(app, roleController)
	routes.FileRoutes(app)

	_ = app.Listen(os.Getenv("HOST") + ":" + os.Getenv("PORT"))
}
