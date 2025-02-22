package models

type TokenResponse struct {
	Token string `json:"token"`
}

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type UserProfileResponse struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}

type BlobResponse struct {
	ID         string `json:"id"`
	Width      int    `json:"width"`
	Height     int    `json:"height"`
	PreviewUrl string `json:"previewUrl"`
}

type BlobSimilarResponse struct {
	ID         string                   `json:"id"`
	Width      int                      `json:"width"`
	Height     int                      `json:"height"`
	PreviewUrl string                   `json:"previewUrl"`
	Similar    []PostSimilarityListItem `json:"similar"`
}
type PostPaginationResponse struct {
	CurrentPage int            `json:"currentPage"`
	TotalPage   int            `json:"totalPage"`
	PostCount   int            `json:"postCount"`
	Posts       []PostListItem `json:"posts"`
	Tags        []TagListItem  `json:"tags"`
}

type PostCountResponse struct {
	PostCount int `json:"postCount"`
}
