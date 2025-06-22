package main

import (
	"go_fullstack/config"
	"go_fullstack/internal/home"
	"go_fullstack/pkg/logger"

	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	app := fiber.New()
	config.Init()
	config.NewDatabaseConfig()
	logConfig := config.NewLogConfig()
	customLogger := logger.NewLogger(logConfig)
	app.Use(recover.New())
	app.Use(fiberzerolog.New(fiberzerolog.Config{
		Logger: customLogger,
	}))
	home.NewHandler(app, customLogger)
	app.Listen(":3000")
}
