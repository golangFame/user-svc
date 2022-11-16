package genesis

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
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
}

type out struct {
	fx.Out
	*Service
}

func New(i in) (o out) {

	logger := logrus.Logger{
		Out:       os.Stderr,
		Formatter: new(logrus.TextFormatter),
		Hooks:     make(logrus.LevelHooks),
		Level:     logrus.DebugLevel,
	}

	logger.SetReportCaller(true)

	o = out{
		Service: &Service{
			logger,
			i.Conf,
		},
	}

	return
}
