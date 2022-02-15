package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/twincats/fiber-app/server/controllers"
)

func SetupRouter(app *fiber.App) {

	app.Get("/", controllers.Index)
	app.Get("/test", controllers.TestDB)

}
