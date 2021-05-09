package models

type PostReadModel struct {
	ID        string   `json:"id"`
	ImagePath string   `json:"image_path"`
	SourceURL string   `json:"source_url"`
	Tags      []string `json:"tags"`
}
