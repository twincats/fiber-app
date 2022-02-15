package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func NotFound() fiber.Handler {
	return func(c *fiber.Ctx) error {
		status := map[string]string{
			"error": "Error Not Found",
		}
		return c.Status(404).JSON(status)
	}
}
