package models

import (
	"database/sql"
	"time"
)

type Model struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

type User struct {
	Model
	Username  string       `gorm:"uniqueIndex;not null" validate:"required,min=3,max=20,alphanum"`
	Password  string       `gorm:"not null"`
	Email     string       `gorm:"uniqueIndex;not null;unique" validate:"required,email"`
	IsAdmin   sql.NullBool `gorm:"default:false"`
	FirstName sql.NullString
	LastName  sql.NullString
	Phone     sql.NullString
}
