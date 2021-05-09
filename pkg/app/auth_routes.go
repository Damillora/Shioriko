package app

import (
	"net/http"

	"github.com/Damillora/Shioriko/pkg/models"
	"github.com/Damillora/Shioriko/pkg/services"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func InitializeAuthRoutes(g *gin.Engine) {
	g.POST("/api/auth/login", createToken)
}
func createToken(c *gin.Context) {
	var model models.LoginFormModel
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

	user := services.Login(model.Username, model.Password)

	if user != nil {
		token := services.CreateToken(user)
		c.JSON(http.StatusOK, models.TokenResponse{
			Token: token,
		})

	} else {
		c.JSON(http.StatusUnauthorized, models.ErrorResponse{
			Code:    http.StatusUnauthorized,
			Message: "Wrong username or password",
		})
	}
}
