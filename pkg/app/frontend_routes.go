package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitializeFrontendRoutes(g *gin.Engine) {
	g.NoRoute(frontendHome)
}

func frontendHome(c *gin.Context) {    
	c.Redirect(http.StatusFound, "/")
}
