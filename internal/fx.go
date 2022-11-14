package internal

import (
	"github.com/BzingaApp/user-svc/internal/cache"
	"github.com/BzingaApp/user-svc/internal/db"
	"github.com/BzingaApp/user-svc/internal/genesis"
	"go.uber.org/fx"
)

var Module = fx.Options(
	cache.Module,
	db.Module,
	genesis.Module,
)
