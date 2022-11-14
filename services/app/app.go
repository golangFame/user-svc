package app

import (
	"github.com/BzingaApp/user-svc/internal/genesis"
	"github.com/uptrace/bun"
)

type service struct {
	*genesis.Service
	DB *bun.DB `name:"db"`
}
