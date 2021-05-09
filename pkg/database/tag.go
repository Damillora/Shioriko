package database

import (
	"time"
)

type Tag struct {
	ID        string `gorm:"size:36"`
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	TagTypeID uint
	TagType   TagType
	Posts     []Post `gorm:"many2many:post_tags"`
}
