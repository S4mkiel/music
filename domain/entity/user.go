package entity

import "github.com/google/uuid"

type User struct {
	Id      uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Email   string    `json:"email" gorm:"type:varchar(255);unique;not null" valid:"email,required"`
	Name    string    `json:"name" gorm:"type:varchar(255);not null" valid:"required"`
	Date    string    `json:"birth_date" gorm:"type:varchar(255);not null" valid:"required"`
	City    string    `json:"city" gorm:"type:varchar(255);not null" valid:"required"`
	Contact string    `json:"contact" gorm:"type:varchar(255);not null" valid:"required"`
	Base    `gorm:"embedded"`
}
