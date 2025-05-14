package app

import (
	"github.com/anthoz69/salepage-api/internal/configs"
	"github.com/anthoz69/salepage-api/internal/modules/auth"
	"github.com/anthoz69/salepage-api/internal/modules/user"
	"github.com/anthoz69/salepage-api/internal/router"
	"go.uber.org/fx"
)

var Modules = fx.Options(
	fx.Provide(
		configs.NewViperConfig,
		configs.NewPostgresDB,
		configs.NewRedisClient,
	),
	user.Module,
	auth.Module,
	fx.Invoke(router.StartFiber),
)
