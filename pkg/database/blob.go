package database

import (
	"time"
)

type Blob struct {
	ID                string `gorm:"size:36"`
	FilePath          string
	PreviewFilePath   string
	ThumbnailFilePath string
	Width             int
	Height            int
	Hash1             []byte
	Hash2             []byte
	Hash3             []byte
	Hash4             []byte
	CreatedAt         time.Time
	UpdatedAt         time.Time
}
