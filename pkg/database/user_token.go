package database

import "time"

type UserToken struct {
	ID        uint
	UserID    string
	User      User
	Token     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
