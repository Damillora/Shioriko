package database

import (
	"time"
)

type User struct {
	ID        string `gorm:"size:36"`
	Email     string
	Username  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
