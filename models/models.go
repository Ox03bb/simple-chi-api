package models

import (
	"time"
)

type Model struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

type User struct {
	Model
	Username  string `gorm:"uniqueIndex;not null" validate:"required,min=3,max=20,alphanum"`
	Password  string `gorm:"not null"`
	Email     string `gorm:"uniqueIndex;not null;unique" validate:"required,email"`
	IsAdmin   *bool  `gorm:"default:false"`
	FirstName *string
	LastName  *string
	Phone     *string
}
