package entity

import "time"

type Base struct {
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;autoCreateTime" valid:"required"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at;autoUpdateTime" valid:"-"`
}
