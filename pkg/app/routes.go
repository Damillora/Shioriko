package app

import (
	"github.com/gin-gonic/gin"
)

func InitializeRoutes(g *gin.Engine) {
	InitializeUserRoutes(g)
	InitializeAuthRoutes(g)
	InitializeTagTypeRoutes(g)
	InitializeTagRoutes(g)
	InitializeBlobRoutes(g)
	InitializePostRoutes(g)

	InitializeFrontendRoutes(g)
}
