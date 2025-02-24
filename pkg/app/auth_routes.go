package app

import (
	"net/http"

	"github.com/Damillora/phoebe/pkg/database"
	"github.com/Damillora/phoebe/pkg/middleware"
	"github.com/Damillora/phoebe/pkg/models"
	"github.com/Damillora/phoebe/pkg/services"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func InitializeAuthRoutes(g *gin.Engine) {
	g.POST("/api/auth/login", createToken)

	protected := g.Group("/api/auth").Use(middleware.AuthMiddleware())
	{
		protected.POST("/token", createTokenLoggedIn)
	}
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

func createTokenLoggedIn(c *gin.Context) {
	result, ok := c.Get("user")
	if ok {
		user := result.(*database.User)
		if user != nil {
			token := services.CreateToken(user)
			c.JSON(http.StatusOK, models.TokenResponse{
				Token: token,
			})
		}
	} else {
		c.JSON(http.StatusUnauthorized, models.ErrorResponse{
			Code:    http.StatusUnauthorized,
			Message: "No authorized user",
		})
	}
}
