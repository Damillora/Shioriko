package app

import (
	"net/http"

	"github.com/Damillora/Shioriko/pkg/services"
	"github.com/gin-gonic/gin"
)

func InitializeTagRoutes(g *gin.Engine) {
	unprotected := g.Group("/api/tag")
	{
		unprotected.GET("/", tagGet)
	}
}

func tagGet(c *gin.Context) {
	tags := services.GetTagAll()
	c.JSON(http.StatusOK, tags)
}
