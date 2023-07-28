package models

import "time"

// User is for users/owners model
type User struct {
	ID          int
	FirstName   string
	LastName    string
	Email       string
	Password    string
	AccessLevel int
	CreatedAt   time.Time
	UpdateAt    time.Time
}

// UserPassword is for the user password model
type UserPassword struct {
	ID           int
	PasswordHash string
}
