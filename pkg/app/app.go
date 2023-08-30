package app

import (
	"embed"
	"net/http"
	"os"

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

	previewDir := config.CurrentConfig.DataDirectory + "/preview"
	thumbnailDir := config.CurrentConfig.DataDirectory + "/thumbnail"

	if _, err := os.Stat(previewDir); os.IsNotExist(err) {
		os.Mkdir(previewDir, 0755)
	}
	if _, err := os.Stat(thumbnailDir); os.IsNotExist(err) {
		os.Mkdir(thumbnailDir, 0755)
	}

	database.Initialize()
}

func Start() {
	g := gin.Default()

	g.StaticFile("/", "./web/static/index.html")
	g.Static("/_app", "./web/static/_app")
	g.Static("/data", config.CurrentConfig.DataDirectory)
	
	g.Use(cors.Default())

	InitializeRoutes(g)

	g.Run()
}
