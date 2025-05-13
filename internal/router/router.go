package router

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"go.uber.org/fx"
)

type RouteRegistrar interface {
	RegisterRoutes(app *fiber.App)
}

type RouteGroup struct {
	fx.In
	Registrars []RouteRegistrar `group:"route_registrars"`
}

func StartFiber(
	lc fx.Lifecycle,
	v *viper.Viper,
	r RouteGroup,
) *fiber.App {
	app := fiber.New()

	// Register routes from all modules
	for _, r := range r.Registrars {
		r.RegisterRoutes(app)
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				if err := app.Listen(v.GetString("app.port")); err != nil {
					panic(err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			shutdownCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
			defer cancel()
			return app.ShutdownWithContext(shutdownCtx)
		},
	})

	// Optional: support SIGINT/SIGTERM shutdown
	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit
		_ = app.Shutdown()
	}()

	return app
}
