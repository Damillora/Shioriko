package database

import (
	"time"
)

type Post struct {
	ID        string `gorm:"size:36"`
	UserID    string
	User      User
	BlobID    string
	Blob      Blob `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	SourceURL string
	Tags      []Tag `gorm:"many2many:post_tags;constraint:OnDelete:CASCADE"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
