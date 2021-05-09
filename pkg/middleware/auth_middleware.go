package middleware

import (
	"strings"

	"github.com/Damillora/Shioriko/pkg/models"
	"github.com/Damillora/Shioriko/pkg/services"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientToken := c.Request.Header.Get("Authorization")
		if clientToken == "" {
			c.JSON(403, models.ErrorResponse{
				Code:    403,
				Message: "Authorization required",
			})
			c.Abort()
			return
		}

		extractedToken := strings.Split(clientToken, "Bearer ")

		if len(extractedToken) == 2 {
			clientToken = strings.TrimSpace(extractedToken[1])
		} else {
			c.JSON(400, models.ErrorResponse{
				Code:    400,
				Message: "Incorrect Format of Authorization Token",
			})
			c.Abort()
			return
		}

		claims, err := services.ValidateToken(clientToken)
		if err != nil {
			c.JSON(401, models.ErrorResponse{
				Code:    401,
				Message: err.Error(),
			})
			c.Abort()
			return
		}
		user := services.GetUser(claims["sub"].(string))
		c.Set("user", user)

		c.Next()
	}
}
