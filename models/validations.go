package models

import (
	"errors"
	"regexp"

	"gorm.io/gorm"
)

//! User

func (user *User) UserNameValidation() bool {

	if user.Username == "" {
		return false
	}
	// Username must be 3-30 characters, only letters, numbers, and underscores
	re := regexp.MustCompile(`^[a-zA-Z0-9_]{3,30}$`)
	return re.MatchString(user.Username)
}

func (user *User) EmailValidation() bool {
	if user.Email == "" {
		return false
	}
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(user.Email)
}

func (user *User) PasswordValidation() bool {
	if user.Password == "" || user.Password == user.Username || user.Password == user.Email {
		return false
	}
	// Password must be 8-30 characters, at least one letter and one number
	re := regexp.MustCompile(`^[A-Za-z\d_]{8,30}$`)

	return re.MatchString(user.Password)
}

func (u *User) BeforeSave(tx *gorm.DB) (err error) {

	if u.Username != "" {
		if !u.UserNameValidation() {
			return errors.New("invalid username")
		}
	}
	if u.Email != "" {
		if !u.EmailValidation() {
			return errors.New("invalid email")
		}
	}
	return
}
