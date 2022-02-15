package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/twincats/fiber-app/server/controllers"
)

func ApiRoute(api fiber.Router) {

	api.Get("/manga", controllers.TestDB)
}
