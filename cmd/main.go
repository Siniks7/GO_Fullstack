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
	"github.com/gofiber/fiber/v2/middleware/session"
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
	store := session.New()

	// Repositories
	vacancyRepo := vacancy.NewVacancyRepository(dbpool, customLogger)

	app.Use(fiberzerolog.New(fiberzerolog.Config{
		Logger: customLogger,
	}))
	home.NewHandler(app, customLogger, vacancyRepo, store)
	vacancy.NewHandler(app, customLogger, vacancyRepo)
	app.Listen(":3000")
}
