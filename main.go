package main

import (
	"embed"
	"wails-temp/service"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	App := service.NewApp()
	// Create application with options
	err := wails.Run(&options.App{
		Title:  "wails-temp",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        App.Startup,
		Bind: []interface{}{
			App,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
