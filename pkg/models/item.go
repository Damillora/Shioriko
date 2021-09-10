package models

type TagTypeListItem struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type TagListItem struct {
	TagID     string `json:"tagId"`
	TagName   string `json:"tagName"`
	TagType   string `json:"tagType"`
	PostCount int    `json:"postCount"`
}

type TagAutocompleteListItem struct {
	Name string `json:"name"`
}

type PostListItem struct {
	ID                 string   `json:"id"`
	ImagePath          string   `json:"image_path"`
	ImageThumbnailPath string   `json:"thumbnail_path"`
	Tags               []string `json:"tags"`
}

type PostSimilarityListItem struct {
	ID       string `json:"id"`
	Distance int    `json:"distance"`
}
