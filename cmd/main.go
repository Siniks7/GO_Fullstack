package main

import (
	"go_fullstack/config"
	"go_fullstack/internal/home"
	"go_fullstack/internal/sitemap"
	"go_fullstack/internal/vacancy"
	"go_fullstack/pkg/database"
	"go_fullstack/pkg/logger"
	"go_fullstack/pkg/middleware"
	"time"

	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/postgres/v3"
)

func main() {
	config.Init()
	dbConfig := config.NewDatabaseConfig()
	logConfig := config.NewLogConfig()
	customLogger := logger.NewLogger(logConfig)
	app := fiber.New()
	app.Use(recover.New())
	app.Static("/public", "./public")
	app.Static("/robots.txt", "./public/robots.txt")
	dbpool := database.CreateDbPool(dbConfig, customLogger)
	defer dbpool.Close()
	storage := postgres.New(postgres.Config{
		DB:         dbpool,
		Table:      "sessions",
		Reset:      false,
		GCInterval: 10 * time.Second,
	})
	store := session.New(session.Config{
		Storage: storage,
	})
	app.Use(middleware.AuthMiddleware(store))

	// Repositories
	vacancyRepo := vacancy.NewVacancyRepository(dbpool, customLogger)

	app.Use(fiberzerolog.New(fiberzerolog.Config{
		Logger: customLogger,
	}))

	// Handlers
	home.NewHandler(app, customLogger, vacancyRepo, store)
	vacancy.NewHandler(app, customLogger, vacancyRepo)
	sitemap.NewHandler(app)

	app.Listen(":3000")
}
