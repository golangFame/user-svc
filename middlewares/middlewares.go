package middlewares

import (
	"github.com/sirupsen/logrus"
	"github.com/uptrace/bun"

	"github.com/spf13/viper"
)

type Middleware struct {
	log  *logrus.Logger //not yet used
	db   *bun.DB
	conf *viper.Viper
}
