package db_test

import (
	"os"
	"path/filepath"
	"testing"
	"world-builder/internal/db"
)

func TestInitDB(t *testing.T) {
	tmpDir := t.TempDir()
	dbPath := filepath.Join(tmpDir, "test_world.db")

	database, err := db.InitDB(dbPath)

	if err != nil {
		t.Fatalf("Expected InitDB to succeed, got error: %v", err)
	}

	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
		t.Errorf("Expected database file to be created at %s, but it does not exist", dbPath)
	}

	sqlDB, err := database.DB()
	if err != nil {
		t.Fatalf("Failed to get underlying sql.DB: %v", err)
	}

	if err := sqlDB.Ping(); err != nil {
		t.Errorf("Expected database connection to be active, got ping error: %v", err)
	}
}
