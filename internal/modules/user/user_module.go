package user

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(
		NewUserRepository,
		NewUserUseCase,
		NewUserHandler,
		fx.Annotate(
			NewRouteRegistrar,
			fx.ResultTags(`group:"route_registrars"`),
		),
	),
)
