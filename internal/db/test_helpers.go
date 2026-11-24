package db

import (
	"testing"
	"world-builder/internal/models"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func SetupTestDB(t *testing.T) *gorm.DB {
	t.Helper()

	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect to test database: %v", err)
	}

	err = db.AutoMigrate(
		&models.Gender{},
		&models.Race{},
		&models.Occupation{},
		&models.Alignment{},
		&models.Status{},
		&models.PersonalityTrait{},
		&models.Strength{},
		&models.Flaw{},
		&models.Weakness{},
		&models.Fear{},
		&models.LoreCategory{},

		// 2. Base Entity Tables
		&models.Faction{},
		&models.Location{},

		// 3. Dependent Tables (Reference base entities/lookups)
		&models.Affiliation{}, // References Faction
		&models.Character{},   // References Gender, Race, Faction, Affiliation, etc.
		&models.Lore{},        // References LoreCategory, Character

		// 4. Intermediary / Self-Referencing Models
		&models.Relationship{}, // References Character (Source & Target)

		// 5. The Biggest page to see
		&models.Article{},
	)
	if err != nil {
		t.Fatalf("Failed to migrate test database: %v", err)
	}
	return db
}
