package app

import (
	"github.com/gin-gonic/gin"
)

func InitializeFrontendRoutes(g *gin.Engine) {
	g.NoRoute(frontendHome)
}

func frontendHome(c *gin.Context) {    
	c.File("./web/static/index.html")
}
