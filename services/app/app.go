package app

import (
	"github.com/BzingaApp/user-svc/internal/cache"
	"github.com/BzingaApp/user-svc/internal/genesis"
	"github.com/uptrace/bun"
)

type service struct {
	*genesis.Service
	db            *bun.DB `name:"db"`
	cacheServices cache.Services
}
