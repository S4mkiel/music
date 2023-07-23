package service

import (
	"context"

	"github.com/S4mkiel/music/domain/entity"
	"github.com/S4mkiel/music/domain/repo"
)

type UserService struct {
	repo repo.UserRepository
}

func NewUserService(repo repo.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) Create(ctx context.Context, user entity.User) (*entity.User, error) {
	u, e := s.repo.Create(ctx, user)
	if e != nil {
		return nil, e
	} else {
		return u, nil
	}
}

