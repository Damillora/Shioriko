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
	ID string `json:"id"`
}

type PostPaginationResponse struct {
	CurrentPage int            `json:"currentPage"`
	TotalPage   int            `json:"totalPage"`
	Posts       []PostListItem `json:"posts"`
}
