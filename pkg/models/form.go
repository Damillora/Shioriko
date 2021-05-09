package models

type LoginFormModel struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}
