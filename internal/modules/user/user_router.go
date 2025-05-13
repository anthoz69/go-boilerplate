package user

import (
	"github.com/anthoz69/salepage-api/internal/router"
	"github.com/gofiber/fiber/v2"
)

type routeRegistrar struct {
	handler *UserHandler
}

func NewRouteRegistrar(h *UserHandler) router.RouteRegistrar {
	return &routeRegistrar{handler: h}
}

func (r *routeRegistrar) RegisterRoutes(app *fiber.App) {
	api := app.Group("/api")

	users := api.Group("/users")
	users.Post("", r.handler.CreateUser)
}
