package db

import (
	"github.com/BzingaApp/user-svc/internal/genesis"
	"github.com/uptrace/bun"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		newDBS,
	),
)

type in struct {
	fx.In
	Service *genesis.Service
}

type out struct {
	fx.Out

	DB1 *bun.DB `name:"db"`
}

func newDBS(i in) (o out) {
	database := &DB{
		i.Service,
	}
	o = out{
		DB1: newPostgressDB(database),
	}
	return
}
