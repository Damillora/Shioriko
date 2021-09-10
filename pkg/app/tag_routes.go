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
		unprotected.GET("/autocomplete", tagAutocomplete)
	}
}

func tagGet(c *gin.Context) {
	tags := services.GetTagAll()
	c.JSON(http.StatusOK, tags)
}

func tagAutocomplete(c *gin.Context) {
	tags := services.GetTagAutocomplete()
	c.JSON(http.StatusOK, tags)
}
