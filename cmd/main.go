package main

import (
	"go_fullstack/config"
	"go_fullstack/internal/home"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())
	app.Use(recover.New())
	config.Init()
	config.NewDatabaseConfig()
	logConfig := config.NewLogConfig()
	log.SetLevel(log.Level(logConfig.Level))
	home.NewHandler(app)
	app.Listen(":3000")
}
