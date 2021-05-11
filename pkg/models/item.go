package models

type TagTypeListItem struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type TagListItem struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	TagType string `json:"tagType"`
}

type PostListItem struct {
	ID                 string   `json:"id"`
	ImagePath          string   `json:"image_path"`
	ImageThumbnailPath string   `json:"thumbnail_path"`
	Tags               []string `json:"tags"`
}
