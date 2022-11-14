package handlers

import (
	"github.com/BzingaApp/user-svc/services/app"
	"github.com/BzingaApp/user-svc/services/dummy"
	"github.com/sirupsen/logrus"
	"os"

	"github.com/spf13/viper"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		New,
	),
)

type Handlers struct {
	*HomeHandler
	*DummyHandler
	*AppHandler
}

type In struct {
	fx.In

	Conf          *viper.Viper
	DummyServices dummy.Services
	AppServices   app.Services
}

type Out struct {
	fx.Out
	*Handlers
}

func New(i In) (o Out) {

	Handler := Handler{
		i.Conf,
		&logrus.Logger{
			Out:       os.Stderr,
			Formatter: new(logrus.TextFormatter),
			Hooks:     make(logrus.LevelHooks),
			Level:     logrus.DebugLevel,
		},
	}

	o = Out{
		Handlers: &Handlers{
			&HomeHandler{
				Handler,
			},
			&DummyHandler{
				Handler,
				i.DummyServices,
			},
			&AppHandler{
				Handler,
				i.AppServices,
			},
		},
	}
	return
}
