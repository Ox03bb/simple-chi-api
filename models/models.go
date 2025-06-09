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
	Username  string       `gorm:"uniqueIndex;not null"`
	Password  string       `gorm:"not null"`
	Email     string       `gorm:"uniqueIndex;not null"`
	IsAdmin   sql.NullBool `gorm:"default:false"`
	FirstName sql.NullString
	LastName  sql.NullString
	Phone     sql.NullString
}
