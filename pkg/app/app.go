package app

import (
	"embed"
	"net/http"

	"github.com/Damillora/Shioriko/pkg/config"
	"github.com/Damillora/Shioriko/pkg/database"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

type embedFileSystem struct {
	http.FileSystem
}

func (e embedFileSystem) Exists(prefix string, path string) bool {
	_, err := e.Open(path)
	if err != nil {
		return false
	}
	return true
}

func EmbedFolder(fsEmbed embed.FS) static.ServeFileSystem {
	return embedFileSystem{
		FileSystem: http.FS(fsEmbed),
	}
}

func Initialize() {
	config.InitializeConfig()
	database.Initialize()
}

func Start() {
	g := gin.Default()

	g.Static("/static", "./web/static")
	g.Static("/data", config.CurrentConfig.DataDirectory)

	g.LoadHTMLGlob("web/template/**/*")

	g.Use(cors.Default())

	InitializeRoutes(g)

	g.Run()
}
