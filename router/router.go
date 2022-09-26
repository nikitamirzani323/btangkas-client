package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Init() *fiber.App {
	app := fiber.New()
	app.Use(cors.New())
	app.Use(logger.New())
	// Custom config
	app.Static("/", "frontend/public", fiber.Static{
		Compress:  true,
		ByteRange: true,
		Browse:    true,
	})
	return app
}
