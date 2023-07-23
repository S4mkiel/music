package module

import (
	"github.com/S4mkiel/music/domain/service"
	"go.uber.org/fx"
)

var Service = fx.Module("services",
	fx.Provide(service.NewUserService),
)
