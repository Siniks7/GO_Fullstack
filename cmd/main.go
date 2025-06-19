package main

import (
	"go_fullstack/internal/home"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	home.NewHandler(app)
	app.Listen(":3000")
}
