package db

import (
	src "github.com/S4mkiel/music/infra/db/src/postgres"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

var SourceModule = fx.Module("implementations", fx.Provide(NewSQLImplementations))

type Sources struct {
	UserSQL *src.UserSQLRepository
}

func NewSQLImplementations(db *gorm.DB) *Sources {
	var src = Sources{
		UserSQL: src.NewUserRepository(db),
	}
	return &src
}
