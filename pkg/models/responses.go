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
	Username string `json:"password"`
}

type BlobResponse struct {
	ID     string `json:"id"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

type PostPaginationResponse struct {
	CurrentPage int            `json:"currentPage"`
	PostCount   int            `json:"postCount"`
	Posts       []PostListItem `json:"posts"`
}
