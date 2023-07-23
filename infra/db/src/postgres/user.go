package db

import (
	"context"
	"errors"

	"github.com/S4mkiel/music/domain/entity"
	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
)

type UserSQLRepository struct {
	orm *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserSQLRepository {
	return &UserSQLRepository{orm: db}
}

func (db UserSQLRepository) Create(ctx context.Context, u entity.User) (*entity.User, error) {
	err := db.orm.Create(&u).Error
	if err != nil {
		if pgError := err.(*pgconn.PgError); errors.Is(err, pgError) {
			switch pgError.Code {
			case "23505":
				if db.orm.Where(entity.User{Email: u.Email, Name: u.Name, Date: u.Date, City: u.City, Contact: u.Contact}).Take(&entity.User{}).Error == nil {
					return nil, errors.New("internal server error")
				}
			}
		}
		return nil, err
	}
	return &u, nil
}
