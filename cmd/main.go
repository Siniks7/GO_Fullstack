package main

import (
	"go_fullstack/config"
	"go_fullstack/internal/home"
	"go_fullstack/internal/vacancy"
	"go_fullstack/pkg/database"
	"go_fullstack/pkg/logger"

	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	config.Init()
	dbConfig := config.NewDatabaseConfig()
	logConfig := config.NewLogConfig()
	customLogger := logger.NewLogger(logConfig)
	app := fiber.New()
	app.Use(recover.New())
	app.Static("/public", "./public")
	dbpool := database.CreateDbPool(dbConfig, customLogger)
	defer dbpool.Close()
	app.Use(fiberzerolog.New(fiberzerolog.Config{
		Logger: customLogger,
	}))
	home.NewHandler(app, customLogger)
	vacancy.NewHandler(app, customLogger)
	app.Listen(":3000")
}
