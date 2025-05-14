package auth

import (
	"github.com/anthoz69/salepage-api/internal/middleware"
	"github.com/anthoz69/salepage-api/internal/router"
	"github.com/gofiber/fiber/v2"
)

type routeRegistrar struct {
	handler *AuthHandler
}

func NewRouteRegistrar(h *AuthHandler) router.RouteRegistrar {
	return &routeRegistrar{handler: h}
}

func (r *routeRegistrar) RegisterRoutes(app *fiber.App) {
	api := app.Group("/api")

	auth := api.Group("/auth")
	auth.Use(middleware.JWTAuth())
	auth.Get("/me", r.handler.GetMe)
}
