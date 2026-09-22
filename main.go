package main

import (
	"embed"
	"log"
	"os"
	"path/filepath"
	"world-builder/internal/db"

	"github.com/wailsapp/wails/v3/pkg/application"
)

var assets embed.FS

func main() {
	database, err := db.InitDB("worldBuilder.db")
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	os.Setenv("WEBKIT_DISABLE_COMPOSITING_MODE", "0")
	os.Setenv("WEBKIT_DISABLE_DMABUF_RENDERER", "1")
	os.Setenv("GDK_BACKEND", "x11")

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

func initStorage() (string, string) {
	configDir, _ := os.UserConfigDir()
	appDir := filepath.Join(configDir, "world-builder")

	imagesDir := filepath.Join(appDir, "assets", "portraits")
	os.MkdirAll(imagesDir, 0755)

	dbPath := filepath.Join(appDir, "worldBuilder.db")
	return dbPath, imagesDir
}
