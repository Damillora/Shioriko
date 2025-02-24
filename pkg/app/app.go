package app

import (
	"io/fs"
	"net/http"
	"os"

	"github.com/Damillora/phoebe/pkg/config"
	"github.com/Damillora/phoebe/pkg/database"
	"github.com/Damillora/phoebe/pkg/web"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

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

	webFS := web.WebAssets()
	webAssets, _ := fs.Sub(webFS, "_app")

	g.NoRoute(func(c *gin.Context) {
		//c.String(http.StatusOK, "AAA")
		c.FileFromFS("./app.html", http.FS(webFS))
	})
	g.StaticFS("/_app", http.FS(webAssets))
	g.Static("/data", config.CurrentConfig.DataDirectory)
	g.Use(cors.Default())

	InitializeRoutes(g)

	g.Run()
}
