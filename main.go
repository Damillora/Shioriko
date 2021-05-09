package main

import (
	"embed"

	"github.com/Damillora/Shioriko/pkg/app"
)

//go:embed web/static
var staticFiles embed.FS

//go:embed web/template
var templateFiles embed.FS

func main() {
	app.Initialize()
	app.Start()
}
