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

type In struct {
	fx.In
	Conf *viper.Viper
}

type Out struct {
	fx.Out
	*Service
}

func New(i In) (o Out) {

	logger := logrus.Logger{
		Out:       os.Stderr,
		Formatter: new(logrus.TextFormatter),
		Hooks:     make(logrus.LevelHooks),
		Level:     logrus.DebugLevel,
	}

	logger.SetReportCaller(true)

	o = Out{
		Service: &Service{
			logger,
			i.Conf,
		},
	}

	return
}
