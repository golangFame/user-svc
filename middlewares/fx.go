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

type in struct {
	fx.In
	Conf *viper.Viper
	DB   *bun.DB `name:"db"`
}

type out struct {
	fx.Out
	*Middleware
}

func New(i in) (o out) {
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
	o = out{
		Middleware: m,
	}
	return
}
