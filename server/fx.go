package server

import (
	"fmt"
	"github.com/BzingaApp/user-svc/enums"
	"github.com/BzingaApp/user-svc/middlewares"
	"github.com/BzingaApp/user-svc/router"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/uptrace/bun"
	"go.uber.org/fx"
	"os"
)

var Module = fx.Options(
	fx.Invoke(
		run, //synchronously

	),
	fx.Provide(initLogrus),
)

type in struct {
	fx.In
	Conf           *viper.Viper
	Middlewares    *middlewares.Middleware
	RouterServices router.Services
	DB             *bun.DB `name:"db"`
}

type out struct {
	fx.Out

	Log *logrus.Logger
}

func run(i in) {
	addr := "0.0.0.0"
	server := &Server{
		i.Middlewares,
		i.RouterServices,
		i.Conf,
		&logrus.Logger{
			Out:       os.Stderr,
			Formatter: new(logrus.TextFormatter),
			Hooks:     make(logrus.LevelHooks),
			Level:     logrus.DebugLevel,
		},
		i.DB,
	}

	r1 := server.setupRouter()
	server.log.Info("running the server on port ", i.Conf.GetString(enums.PORT))

	go func() {
		err := r1.Run(fmt.Sprintf("%s:%s", addr, i.Conf.GetString(enums.PORT)))
		server.log.Fatal(err)
	}()
}
