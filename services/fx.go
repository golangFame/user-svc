package services

import (
	"github.com/BzingaApp/user-svc/internal/cache"
	"github.com/BzingaApp/user-svc/services/dummy"
	"go.uber.org/fx"
)

var Module = fx.Options(
	cache.Module, //redis
	dummy.Module,
)
