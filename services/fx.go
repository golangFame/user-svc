package services

import (
	"github.com/BzingaApp/user-svc/services/dummy"
	"go.uber.org/fx"
)

var Module = fx.Options(
	dummy.Module,
)
