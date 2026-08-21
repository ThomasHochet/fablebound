package main

import (
	"embed"
	"log"
	"world-builder/internal/db"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	database, err := db.InitDB("worldBuilder.db")
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Create an instance of the app structure
	app := NewApp(database)

	// Create application with options
	err = wails.Run(&options.App{
		Title:  "world-builder",
		Width:  1920,
		Height: 1080,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
