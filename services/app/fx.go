package app

import (
	"github.com/BzingaApp/user-svc/internal/cache"
	"github.com/BzingaApp/user-svc/internal/genesis"
	"github.com/uptrace/bun"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		newServices,
	),
)

type in struct {
	fx.In
	*genesis.Service
	DB            *bun.DB `name:"db"`
	CacheServices cache.Services
}

type out struct {
	fx.Out

	AppServices Services
}

func newServices(i in) (o out) {
	o = out{
		AppServices: newApp(i.Service, i.DB, i.CacheServices),
	}
	return
}

func newApp(genesis *genesis.Service, DB *bun.DB, cacheServices cache.Services) Services {
	return &service{
		genesis,
		DB,
		cacheServices,
	}
}
