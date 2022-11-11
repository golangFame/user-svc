package config

import (
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		New,
	),
)

func New() (config *viper.Viper) {
	config = viper.New()

	for key, meta := range confList {
		// automatic conversion of environment var key to `UPPER_CASE` will happen.
		config.BindEnv(key)

		// read command-line arguments
		pflag.String(key, meta.defaultVal, meta.desc)
	}
	//viper.AutomaticEnv() //replaces the default values//not required really
	pflag.Parse()
	config.BindPFlags(pflag.CommandLine)
	return
}
