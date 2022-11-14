package app

import (
	"github.com/BzingaApp/user-svc/internal/genesis"
	"github.com/uptrace/bun"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		newServices,
	),
)

type In struct {
	fx.In
	*genesis.Service
	DB *bun.DB `name:"db"`
}

type Out struct {
	fx.Out

	AppServices Services
}

func newServices(i In) (o Out) {
	o = Out{
		AppServices: newApp(i.Service, i.DB),
	}
	return
}

func newApp(genesis *genesis.Service, DB *bun.DB) Services {
	return &service{
		genesis,
		DB,
	}
}
