package main

import (
	"go_fullstack/config"
	"go_fullstack/internal/home"

	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/rs/zerolog"
)

func main() {
	app := fiber.New()
	app.Use(recover.New())
	config.Init()
	config.NewDatabaseConfig()
	logConfig := config.NewLogConfig()
	zerolog.SetGlobalLevel(zerolog.Level(logConfig.Level))
	app.Use(fiberzerolog.New())
	home.NewHandler(app)
	app.Listen(":3000")
}
