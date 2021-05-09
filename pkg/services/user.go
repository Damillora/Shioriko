package services

import (
	"github.com/Damillora/Shioriko/pkg/database"
	"github.com/Damillora/Shioriko/pkg/models"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(model models.UserCreateModel) (*database.User, error) {
	passwd, err := bcrypt.GenerateFromPassword([]byte(model.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := database.User{
		ID:       uuid.NewString(),
		Email:    model.Email,
		Username: model.Username,
		Password: string(passwd),
	}
	result := database.DB.Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func GetUser(id string) *database.User {
	var user database.User
	database.DB.Where("id = ?", id).First(&user)

	return &user
}
func GetUserFromUsername(username string) *database.User {
	var user database.User
	database.DB.Where("username = ?", username).First(&user)

	return &user
}

func UpdateUser(id string, model models.UserUpdateModel) (*database.User, error) {
	var user database.User
	result := database.DB.Where("id = ?", id).First(&user)

	if result.Error != nil {
		return nil, result.Error
	}
	user.Email = model.Email
	user.Username = model.Username

	if user.Password != "" {
		passwd, err := bcrypt.GenerateFromPassword([]byte(model.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}
		user.Password = string(passwd)
	}

	result = database.DB.Save(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
