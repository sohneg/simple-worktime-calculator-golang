package main

import (
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	time := NewTime()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "Simple Worktime Calculator",
		Width:  570,
		Height: 400,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        time.startup,
		Bind: []interface{}{
			time,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
