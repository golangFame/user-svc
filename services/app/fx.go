package app

import (
	"github.com/BzingaApp/user-svc/internal/genesis"
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
}

type Out struct {
	fx.Out

	AppServices Services
}

func newServices(i In) (o Out) {
	o = Out{
		AppServices: newApp(i.Service),
	}
	return
}

func newApp(genesis *genesis.Service) Services {
	return &service{
		genesis,
	}
}
