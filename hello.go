package main

import (
	controllers "goScraper/controllers"

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
	api.Get("/search", controllers.ControllerSearch)
	api.Get("/detail", controllers.ControllerDetail)
	api.Get("/video", controllers.ControllerVideo)

	app.Listen(5000)
}
