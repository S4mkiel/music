package db

import (
	"github.com/S4mkiel/music/domain/repo"
	db "github.com/S4mkiel/music/infra/db/src"
	"go.uber.org/fx"
)

var Module = fx.Module("database",
	PostgresModule,
	db.SourceModule,
	fx.Provide(func(s *db.Sources) repo.UserRepository { return s.UserSQL }),
)
