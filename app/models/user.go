package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

// User contains information about a user
type User struct {
	ID           int
	Username     string
	Password     string
	PasswordHash string `db:"-"`
	ProgramID    int
	DateCreated  time.Time
}

// TableName gets the name of the database table
func (u User) TableName() string {
	return "users"
}

// HashPassword generates a hashed a password
func (u *User) HashPassword() (err error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	u.PasswordHash = string(bytes)
	return
}

// CheckPasswordHash validates a password against a hashed password
func (u *User) CheckPasswordHash(password string) (isValid bool) {
	err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password))
	if err != nil {
		return
	}

	isValid = true
	return
}
