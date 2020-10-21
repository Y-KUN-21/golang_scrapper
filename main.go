package main

import (
	controllers "goanimey/controllers"

	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"
)

func main() {
	app := fiber.New()
	api := app.Group("api")

	app.Use(middleware.Recover())
	app.Use(middleware.Logger())
	api.Get("/popular", controllers.ControllerPopular)
	api.Get("/recent", controllers.ControllerRecent)
	api.Get("/random", controllers.ControllerRandom)
	api.Get("/search", controllers.ControllerSearch)
	api.Get("/detail", controllers.ControllerDetail)
	api.Get("/video", controllers.ControllerVideo)
	api.Get("/test", controllers.ControllerTest)

	// port := os.Getenv("PORT")
	app.Settings.CaseSensitive = true
	app.Settings.StrictRouting = true
	app.Listen("5000")
}
