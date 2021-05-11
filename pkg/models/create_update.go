package models

type UserCreateModel struct {
	Email    string `json:"email" validate:"required,email"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserUpdateModel struct {
	Email    string `json:"email" validate:"required,email"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password"`
}

type TagTypeCreateModel struct {
	Name string `json:"name" validate:"required,ascii"`
}

type TagCreateModel struct {
	Name      string `json:"name" validate:"required,ascii"`
	TagTypeID uint   `json:"tagTypeId" validate:"required"`
}

type TagUpdateModel struct {
	Name      string `json:"name" validate:"required,ascii"`
	TagTypeID uint   `json:"tagTypeId" validate:"required"`
}

type PostCreateModel struct {
	BlobID    string   `json:"blob_id" validate:"required"`
	SourceURL string   `json:"source_url"`
	Tags      []string `json:"tags"`
}

type PostUpdateModel struct {
	SourceURL string   `json:"source_url"`
	Tags      []string `json:"tags"`
}
