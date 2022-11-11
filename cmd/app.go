package main

import (
	"github.com/BzingaApp/user-svc/config"
	"github.com/BzingaApp/user-svc/handlers"
	"github.com/BzingaApp/user-svc/internal"
	"github.com/BzingaApp/user-svc/middlewares"
	"github.com/BzingaApp/user-svc/router"
	"github.com/BzingaApp/user-svc/server"
	"github.com/BzingaApp/user-svc/services"
	"go.uber.org/fx"
)

func serverRun() {
	app := fx.New(
		config.Module,
		internal.Module,
		services.Module,
		middlewares.Module,
		handlers.Module,
		router.Module,
		server.Module,
	)
	app.Run()
}
