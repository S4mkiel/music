package repo

import (
	"context"

	"github.com/S4mkiel/music/domain/entity"
)

type UserRepository interface {
	Create(context.Context, entity.User) (*entity.User, error)
}