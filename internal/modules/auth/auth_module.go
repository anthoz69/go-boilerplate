package auth

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(
		NewAuthUseCase,
		NewAuthHandler,
		fx.Annotate(
			NewRouteRegistrar,
			fx.ResultTags(`group:"route_registrars"`),
		),
	),
)
