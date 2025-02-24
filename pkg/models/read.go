package models

import "time"

type PostReadModel struct {
	ID               string        `json:"id"`
	ImagePreviewPath string        `json:"preview_path"`
	ImagePath        string        `json:"image_path"`
	SourceURL        string        `json:"source_url"`
	Tags             []TagListItem `json:"tags"`
	Width            int           `json:"width"`
	Height           int           `json:"height"`
	Uploader         string        `json:"uploader"`
	UploadDate       time.Time     `json:"upload_date"`
}

type TagReadModel struct {
	TagID     string `json:"tagId"`
	TagName   string `json:"tagName"`
	TagType   string `json:"tagType"`
	TagNote   string `json:"tagNote"`
	PostCount int    `json:"postCount"`
}
