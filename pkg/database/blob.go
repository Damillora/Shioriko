package database

import (
	"time"
)

type Blob struct {
	ID        string `gorm:"size:36"`
	FilePath  string
	Width     int
	Height    int
	CreatedAt time.Time
	UpdatedAt time.Time
}
