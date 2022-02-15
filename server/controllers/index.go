package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gookit/config/v2"
	"github.com/twincats/fiber-app/server/models"
)

func Index(c *fiber.Ctx) error {
	config := config.String("app_name")
	return c.Render("index", fiber.Map{
		"Title": config,
	})
}

func TestDB(c *fiber.Ctx) error {
	val := models.GetMangas()
	return c.JSON(val)
}
