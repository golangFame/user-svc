package dummy

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

	Dummy Services // `name:"dummy"`
}

func newServices(i In) (o Out) {
	o = Out{
		Dummy: newDummy(i.Service),
	}
	return
}

func newDummy(genesis *genesis.Service) Services {
	return &service{
		genesis,
	}
}
