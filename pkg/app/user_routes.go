package app

import (
	"net/http"

	"github.com/Damillora/phoebe/pkg/config"
	"github.com/Damillora/phoebe/pkg/database"
	"github.com/Damillora/phoebe/pkg/middleware"
	"github.com/Damillora/phoebe/pkg/models"
	"github.com/Damillora/phoebe/pkg/services"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func InitializeUserRoutes(g *gin.Engine) {
	g.POST("/api/user/register", registerUser)

	protected := g.Group("/api/user").Use(middleware.AuthMiddleware())
	{
		protected.GET("/profile", userProfile)
		protected.PUT("/update", userUpdate)
		protected.PUT("/update-password", userUpdatePassword)
	}
}

func registerUser(c *gin.Context) {
	var disableRegistration = config.CurrentConfig.DisableRegistration
	if disableRegistration == "true" {
		c.JSON(http.StatusForbidden, models.ErrorResponse{
			Code:    http.StatusForbidden,
			Message: "Registration is disabled",
		})
		c.Abort()
		return
	}

	var model models.UserCreateModel
	err := c.ShouldBindJSON(&model)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		c.Abort()
		return
	}

	validate := validator.New()
	err = validate.Struct(model)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		c.Abort()
		return
	}

	user, err := services.CreateUser(model)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Cannot create user",
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, models.UserProfileResponse{
		Email:    user.Email,
		Username: user.Username,
	})
}

func userProfile(c *gin.Context) {
	result, ok := c.Get("user")
	if ok {
		user := result.(*database.User)
		c.JSON(http.StatusOK, models.UserProfileResponse{
			Email:    user.Email,
			Username: user.Username,
		})
	} else {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "User does not exist",
		})
	}
}

func userUpdate(c *gin.Context) {
	var model models.UserUpdateModel

	err := c.ShouldBindJSON(&model)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		c.Abort()
		return
	}

	validate := validator.New()
	err = validate.Struct(model)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		c.Abort()
		return
	}

	result, ok := c.Get("user")
	if ok {
		user := result.(*database.User)
		services.UpdateUserProfile(user.ID, model)
		c.JSON(http.StatusOK, nil)
	} else {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "User does not exist",
		})
	}
}

func userUpdatePassword(c *gin.Context) {
	var model models.UserUpdatePasswordModel

	err := c.ShouldBindJSON(&model)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		c.Abort()
		return
	}

	validate := validator.New()
	err = validate.Struct(model)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		c.Abort()
		return
	}

	result, ok := c.Get("user")
	if ok {
		user := result.(*database.User)
		services.UpdateUserPassword(user.ID, model)
		c.JSON(http.StatusOK, nil)
	} else {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "User does not exist",
		})
	}
}
