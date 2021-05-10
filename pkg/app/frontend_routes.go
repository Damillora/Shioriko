package app

import (
	"net/http"

	"github.com/Damillora/Shioriko/pkg/config"
	"github.com/gin-gonic/gin"
)

func InitializeFrontendRoutes(g *gin.Engine) {
	g.NoRoute(frontendHome)
}

func frontendHome(c *gin.Context) {
	baseURL := config.CurrentConfig.BaseURL
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title":    "Shioriko",
		"base_url": baseURL,
	})
}
