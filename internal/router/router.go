package router

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/anthoz69/salepage-api/shared/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
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

	app.Use(logger.New())
	app.Use(cors.New())
	app.Use(recover.New())

	app.Use(func(c *fiber.Ctx) error {
		lang := utils.GetLang(c)
		c.Locals("lang", lang)
		return c.Next()
	})

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
			log.Println("Graceful shutdown server via fx onstop...")
			return app.ShutdownWithContext(shutdownCtx)
		},
	})

	// Optional: support SIGINT/SIGTERM shutdown
	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit
		log.Println("Graceful shutdown via os signal")
		_ = app.Shutdown()
	}()

	return app
}
