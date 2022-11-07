package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/joho/godotenv"
	"iecare-api/src/app/http/controllers"
	"iecare-api/src/app/http/routes"
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
		AppName:                 "IECare API",
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

	// Controllers
	userController := controllers.NewUsersController(services.User)
	roleController := controllers.NewRolesController(services.Role)
	providerController := controllers.NewProvidersController(services.Provider)
	servicesController := controllers.NewServicesController(services.Service)
	categoryController := controllers.NewCategoriesController(services.Category)

	// Routes
	app.Get("/", monitor.New(monitor.Config{
		Title: "IECare API",
	}))

	app.Get("/database", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message":  "Welcome to theIE Care API",
			"status":   c.Response().StatusCode(),
			"database": services.Stats(),
		})
	})

	routes.UserRoutes(app, userController)
	routes.RoleRoutes(app, roleController)
	routes.ProviderRoutes(app, providerController)
	routes.CategoryRoutes(app, categoryController)
	routes.ServiceRoutes(app, servicesController)
	routes.FileRoutes(app)

	_ = app.Listen(os.Getenv("HOST") + ":" + os.Getenv("PORT"))
}
