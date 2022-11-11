package router

import (
	"github.com/BzingaApp/user-svc/handlers"

	"github.com/spf13/viper"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		newServices,
	),
)

type In struct {
	fx.In

	Conf *viper.Viper

	Handlers *handlers.Handlers
}

func newServices(i In) (Services Services) {
	h := i.Handlers
	return newService(i.Conf, h)
}

func newService(
	conf *viper.Viper,
	h *handlers.Handlers,
) *service {
	return &service{
		dummyHandler: h.DummyHandler,
		conf:         conf,
	}
}