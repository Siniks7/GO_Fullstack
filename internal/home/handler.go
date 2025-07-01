package home

import (
	"go_fullstack/internal/vacancy"
	"go_fullstack/pkg/tadapter.go"
	"go_fullstack/views"
	"math"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/rs/zerolog"
)

type HomeHandler struct {
	router       fiber.Router
	customLogger *zerolog.Logger
	repository   *vacancy.VacancyRepository
	store        *session.Store
}

type User struct {
	Id   int
	Name string
}

func NewHandler(router fiber.Router, customLogger *zerolog.Logger, repository *vacancy.VacancyRepository, store *session.Store) {
	h := &HomeHandler{
		router:       router,
		customLogger: customLogger,
		repository:   repository,
		store:        store,
	}
	h.router.Get("/", h.home)
	h.router.Get("/404", h.error)
	h.router.Get("/login", h.login)
}

func (h *HomeHandler) home(c *fiber.Ctx) error {
	PAGE_ITEMS := 2
	page := c.QueryInt("page", 1)
	sess, err := h.store.Get(c)
	if err != nil {
		panic(err)
	}
	if name, ok := sess.Get("name").(string); ok {
		h.customLogger.Info().Msg(name)
	}
	count := h.repository.CountAll()
	vacancies, err := h.repository.GetAll(PAGE_ITEMS, (page-1)*PAGE_ITEMS)
	if err != nil {
		h.customLogger.Error().Msg(err.Error())
		return c.SendStatus(500)
	}
	component := views.Main(vacancies, int(math.Ceil((float64(count) / float64(PAGE_ITEMS)))), page)
	return tadapter.Render(c, component, http.StatusOK)
}

func (h *HomeHandler) login(c *fiber.Ctx) error {
	component := views.Login()
	sess, err := h.store.Get(c)
	if err != nil {
		panic(err)
	}
	sess.Set("name", "Антон")
	if err := sess.Save(); err != nil {
		panic(err)
	}
	return tadapter.Render(c, component, http.StatusOK)
}

func (h *HomeHandler) error(c *fiber.Ctx) error {
	return c.SendString("Error")
}
