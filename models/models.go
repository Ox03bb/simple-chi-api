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
	Username  string `gorm:"uniqueIndex;not null"`
	Password  string `gorm:"not null"`
	Email     string `gorm:"uniqueIndex;not null"`
	IsAdmin   bool   `gorm:"default:false"`
	FirstName *string
	LastName  *string
	Phone     *string
}
