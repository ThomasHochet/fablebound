package main

import (
	"embed"
	"log"
	"world-builder/internal/db"

	"github.com/wailsapp/wails/v3/pkg/application"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	database, err := db.InitDB("worldBuilder.db")
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Create an instance of the app structure
	appStruct := NewApp(database)

	// Create application with options
	app := application.New(application.Options{
		Name:        "world-builder",
		Description: "world builder application",
		Services: []application.Service{
			application.NewService(appStruct),
		},
		Assets: application.AssetOptions{
			Handler: application.AssetFileServerFS(assets),
		},
	})

	app.Window.NewWithOptions(application.WebviewWindowOptions{
		Title:  "World Builder",
		Width:  1920,
		Height: 1080,
		URL:    "/",
	})

	err = app.Run()
	if err != nil {
		log.Fatal(err)
	}
}
