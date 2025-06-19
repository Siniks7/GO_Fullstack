package main

import (
	"go_fullstack/config"
	"go_fullstack/internal/home"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	config.Init()
	dbConf := config.NewDatabaseConfig()
	log.Println(dbConf)
	home.NewHandler(app)
	app.Listen(":3000")
}
