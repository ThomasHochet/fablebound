package db

import (
	"fmt"
	"reflect"
	"world-builder/internal/models"

	"github.com/glebarez/sqlite"
	"github.com/wailsapp/wails/v3/pkg/application"
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

	RegisterWailsNotifier(db)

	return db, nil
}

func RegisterWailsNotifier(db *gorm.DB) {
	db.Callback().Create().After("gorm:create").Register("wails_notify_create", func(d *gorm.DB) {
		emitDBEvent(d, "create")
	})

	db.Callback().Update().After("gorm:update").Register("wails_notify_update", func(d *gorm.DB) {
		emitDBEvent(d, "update")
	})

	db.Callback().Delete().After("gorm:delete").Register("wails_notify_delete", func(d *gorm.DB) {
		emitDBEvent(d, "delete")
	})
}

func emitDBEvent(d *gorm.DB, action string) {
	if d.Error != nil || d.Statement == nil || d.Statement.Table == "" {
		return
	}

	val := d.Statement.ReflectValue
	for val.Kind() == reflect.Pointer || val.Kind() == reflect.Interface {
		val = val.Elem()
	}

	// Guard against slice/array association callbacks where primary key lookup fails
	if !val.IsValid() || val.Kind() != reflect.Struct {
		return
	}

	var id interface{}
	if d.Statement.Schema != nil && d.Statement.Schema.PrioritizedPrimaryField != nil {
		id, _ = d.Statement.Schema.PrioritizedPrimaryField.ValueOf(d.Statement.Context, d.Statement.ReflectValue)
	}

	// 1. Verify GORM is firing the hook
	fmt.Printf("--> GORM Hook Fired: Action='%s', Table='%s', ID='%v'\n", action, d.Statement.Table, id)

	app := application.Get()
	if app == nil {
		// 2. Catch if the Wails app instance isn't available
		fmt.Println("--> ERROR: Wails application.Get() returned nil!")
		return
	}

	fmt.Println("--> Emitting Wails event 'db:change' to frontend...")

	app.Event.Emit("db:change", map[string]interface{}{
		"table":  d.Statement.Table,
		"action": action,
		"id":     id,
	})
}
