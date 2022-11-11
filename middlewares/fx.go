package middlewares

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/uptrace/bun"
	"go.uber.org/fx"
	"os"
)

var Module = fx.Options(
	fx.Provide(
		New,
	),
)

type In struct {
	fx.In
	Conf *viper.Viper
	DB   *bun.DB `name:"db"`
}

type Out struct {
	fx.Out
	*Middleware
}

func New(i In) (o Out) {
	m := &Middleware{
		&logrus.Logger{
			Out:       os.Stderr,
			Formatter: new(logrus.TextFormatter),
			Hooks:     make(logrus.LevelHooks),
			Level:     logrus.DebugLevel,
		},
		i.DB,
		i.Conf,
	}
	o = Out{
		Middleware: m,
	}
	return
}
