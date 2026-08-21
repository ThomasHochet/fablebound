package db

import (
	"fmt"
	"world-builder/internal/models"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func InitDB(dbFileName string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(dbFileName), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("Failed to connect to database: %w", err)
	}

	err = db.AutoMigrate(
		// 1. Independent Lookup Tables (No foreign keys)
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
		return nil, fmt.Errorf("Failed to auto-migrate schema: %w", err)
	}

	return db, nil
}
