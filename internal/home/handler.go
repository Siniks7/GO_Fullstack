package home

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

type HomeHandler struct {
	router       fiber.Router
	customLogger *zerolog.Logger
}

func NewHandler(router fiber.Router, customLogger *zerolog.Logger) {
	h := &HomeHandler{
		router:       router,
		customLogger: customLogger,
	}
	api := h.router.Group("/api")
	api.Get("/", h.home)
	api.Get("/error", h.error)
}

func (h *HomeHandler) home(c *fiber.Ctx) error {
	data := struct {
		Count   int
		IsAdmin bool
		CanUse  bool
	}{Count: 0, IsAdmin: true, CanUse: true}
	return c.Render("page", data)
}

func (h *HomeHandler) error(c *fiber.Ctx) error {
	// h.customLogger.Info().
	// 	Bool("isAdmin", true).
	// 	Str("email", "siniks7@yandex.ru").
	// 	Int("id", 10).
	// 	Msg("Инфо")
	return c.SendString("Error")
}
