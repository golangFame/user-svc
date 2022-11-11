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

type In struct {
	fx.In
	Service *genesis.Service
}

type Out struct {
	fx.Out

	DB1 *bun.DB `name:"db"`
}

func newDBS(i In) (o Out) {
	database := &DB{
		i.Service,
	}
	o = Out{
		DB1: newPostgressDB(database),
	}
	return
}
