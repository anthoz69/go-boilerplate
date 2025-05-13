package main

import (
	"github.com/anthoz69/salepage-api/internal/app"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.NopLogger,
		app.Modules,
	)
	app.Run()
}
