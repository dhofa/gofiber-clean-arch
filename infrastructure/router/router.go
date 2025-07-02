package router

import (
	"github.com/dhofa/gofiber-clean-arch/internal/handler"
	"github.com/gofiber/fiber/v2"
)

type RouteRegistry struct {
	UserHandler *handler.UserHandler
}

func Setup(app *fiber.App, reg *RouteRegistry) {
	api := app.Group("/api/v1/")

	// Modular route registration
	reg.UserHandler.Route(api.Group("/users"))
}
