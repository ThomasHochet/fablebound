package services_test

import (
	"testing"
	"time"
	"world-builder/internal/db"
	"world-builder/internal/models"
	"world-builder/internal/services"
)

func seedLoreCategory(t *testing.T, loreService *services.LoreService, label string) *models.LoreCategory {
	t.Helper()
	cat, err := loreService.CreateCategory(label)
	if err != nil {
		t.Fatalf("failed to seed lore category %q: %v", label, err)
	}
	return cat
}

// --- LORE CATEGORY TESTS ---

func TestLoreService_CategoryCRUD_AndBackstory(t *testing.T) {
	database := db.SetupTestDB(t)
	loreService := services.NewLoreService(database)

	// Empty list initially
	emptyCats, err := loreService.GetAllCategories()
	if err != nil {
		t.Fatalf("GetAllCategories on empty DB failed: %v", err)
	}
	if len(emptyCats) != 0 {
		t.Errorf("expected 0 categories, got %d", len(emptyCats))
	}

	// Create categories including "Backstory" and test alphabetical ordering
	backstoryCat := seedLoreCategory(t, loreService, "Backstory")
	seedLoreCategory(t, loreService, "Historical Event")
	seedLoreCategory(t, loreService, "Mythology")

	cats, err := loreService.GetAllCategories()
	if err != nil {
		t.Fatalf("GetAllCategories failed: %v", err)
	}
	if len(cats) != 3 {
		t.Fatalf("expected 3 categories, got %d", len(cats))
	}
	if cats[0].Label != "Backstory" || cats[1].Label != "Historical Event" || cats[2].Label != "Mythology" {
		t.Errorf("unexpected category ordering: %+v", cats)
	}

	// GetBackstoryLoreCategory
	foundBackstory, err := loreService.GetBackstoryLoreCategory()
	if err != nil {
		t.Fatalf("GetBackstoryLoreCategory failed: %v", err)
	}
	if foundBackstory.ID != backstoryCat.ID || foundBackstory.Label != "Backstory" {
		t.Errorf("expected Backstory category ID %d, got %+v", backstoryCat.ID, foundBackstory)
	}

	// UpdateCategory validation & success
	if _, err := loreService.UpdateCategory(0, "Valid"); err == nil {
		t.Error("expected error updating category with ID 0, got nil")
	}
	if _, err := loreService.UpdateCategory(backstoryCat.ID, ""); err == nil {
		t.Error("expected error updating category with empty label, got nil")
	}

	updatedCat, err := loreService.UpdateCategory(backstoryCat.ID, "Origin Story")
	if err != nil {
		t.Fatalf("UpdateCategory failed: %v", err)
	}
	if updatedCat.Label != "Origin Story" {
		t.Errorf("expected label 'Origin Story', got %q", updatedCat.Label)
	}

	// DeleteCategory success & not found
	if err := loreService.DeleteCategory(updatedCat.ID); err != nil {
		t.Fatalf("DeleteCategory failed: %v", err)
	}
	if err := loreService.DeleteCategory(9999); err == nil {
		t.Error("expected error deleting non-existent category, got nil")
	}
}

// --- LORE ENTRY CRUD & QUERIES ---

func TestLoreService_Create_RequiresCategoryID(t *testing.T) {
	database := db.SetupTestDB(t)
	loreService := services.NewLoreService(database)

	created, err := loreService.Create(models.Lore{
		Title:      "Unfiled Chronicle",
		CategoryID: 0,
	})
	if err == nil {
		t.Fatal("expected error when CategoryID is 0, got nil")
	}
	if created != nil {
		t.Errorf("expected nil lore entry on validation error, got %+v", created)
	}
}

func TestLoreService_CreateAndGetByID_WithPreloads(t *testing.T) {
	database := db.SetupTestDB(t)
	loreService := services.NewLoreService(database)

	cat := seedLoreCategory(t, loreService, "Prophecy")
	char := models.Character{Firstname: "Elendil", Surname: "Tall"}
	if err := database.Create(&char).Error; err != nil {
		t.Fatalf("failed to seed character: %v", err)
	}

	charID := char.ID
	created, err := loreService.Create(models.Lore{
		Title:       "Oath of Elendil",
		CategoryID:  cat.ID,
		CharacterID: &charID,
	})
	if err != nil {
		t.Fatalf("Create failed: %v", err)
	}
	if created == nil || created.ID == 0 {
		t.Fatalf("expected created lore with non-zero ID, got %+v", created)
	}

	fetched, err := loreService.GetByID(created.ID)
	if err != nil {
		t.Fatalf("GetByID failed: %v", err)
	}
	if fetched.LoreCategory.ID != cat.ID {
		t.Errorf("expected preloaded LoreCategory ID %d, got %d", cat.ID, fetched.LoreCategory.ID)
	}
	if fetched.Character == nil || fetched.Character.ID != char.ID {
		t.Errorf("expected preloaded Character ID %d, got %+v", char.ID, fetched.Character)
	}
}

func TestLoreService_GetByID_NotFound(t *testing.T) {
	database := db.SetupTestDB(t)
	loreService := services.NewLoreService(database)

	fetched, err := loreService.GetByID(9999)
	if err == nil {
		t.Fatal("expected error fetching non-existent lore entry, got nil")
	}
	if fetched != nil {
		t.Errorf("expected nil lore entry when not found, got %+v", fetched)
	}
}

func TestLoreService_Queries_CharacterCategoryAndGeneral(t *testing.T) {
	database := db.SetupTestDB(t)
	loreService := services.NewLoreService(database)

	catHistory := seedLoreCategory(t, loreService, "History")
	catLegend := seedLoreCategory(t, loreService, "Legend")

	char1 := models.Character{Firstname: "Isildur", Surname: "Son of Elendil"}
	char2 := models.Character{Firstname: "Anárion", Surname: "Son of Elendil"}
	_ = database.Create(&char1).Error
	_ = database.Create(&char2).Error

	char1ID := char1.ID
	char2ID := char2.ID
	now := time.Now()

	entries := []models.Lore{
		{
			Title:       "General World Lore",
			CategoryID:  catHistory.ID,
			CharacterID: nil,
			CreatedAt:   now.Add(-2 * time.Hour),
		},
		{
			Title:       "Isildur's Journal",
			CategoryID:  catLegend.ID,
			CharacterID: &char1ID,
			CreatedAt:   now.Add(-1 * time.Hour),
		},
		{
			Title:       "Anárion's Defense",
			CategoryID:  catHistory.ID,
			CharacterID: &char2ID,
			CreatedAt:   now,
		},
	}

	for _, e := range entries {
		if _, err := loreService.Create(e); err != nil {
			t.Fatalf("failed to seed lore %q: %v", e.Title, err)
		}
	}

	// GetAll (newest first)
	all, err := loreService.GetAll()
	if err != nil {
		t.Fatalf("GetAll failed: %v", err)
	}
	if len(all) != 3 {
		t.Fatalf("expected 3 lore entries, got %d", len(all))
	}
	if all[0].Title != "Anárion's Defense" {
		t.Errorf("expected newest entry first ('Anárion's Defense'), got %q", all[0].Title)
	}

	// GetByCharacterID
	char1Lore, err := loreService.GetByCharacterID(char1.ID)
	if err != nil {
		t.Fatalf("GetByCharacterID failed: %v", err)
	}
	if len(char1Lore) != 1 || char1Lore[0].Title != "Isildur's Journal" {
		t.Errorf("expected 1 lore entry for char1, got %+v", char1Lore)
	}

	// GetByCategoryID
	historyLore, err := loreService.GetByCategoryID(catHistory.ID)
	if err != nil {
		t.Fatalf("GetByCategoryID failed: %v", err)
	}
	if len(historyLore) != 2 {
		t.Fatalf("expected 2 lore entries in History category, got %d", len(historyLore))
	}

	// GetGeneralLore (CharacterID IS NULL)
	generalLore, err := loreService.GetGeneralLore()
	if err != nil {
		t.Fatalf("GetGeneralLore failed: %v", err)
	}
	if len(generalLore) != 1 || generalLore[0].Title != "General World Lore" {
		t.Errorf("expected 1 general lore entry, got %+v", generalLore)
	}
}

func TestLoreService_UpdateAndDelete(t *testing.T) {
	database := db.SetupTestDB(t)
	loreService := services.NewLoreService(database)

	if _, err := loreService.Update(models.Lore{ID: 0, CategoryID: 1}); err == nil {
		t.Error("expected error updating lore with ID 0, got nil")
	}

	cat := seedLoreCategory(t, loreService, "Chronicle")
	created, err := loreService.Create(models.Lore{
		Title:      "Draft Entry",
		CategoryID: cat.ID,
	})
	if err != nil {
		t.Fatalf("seed failed: %v", err)
	}

	updatedInput := *created
	updatedInput.Title = "Final Entry"
	updated, err := loreService.Update(updatedInput)
	if err != nil {
		t.Fatalf("Update failed: %v", err)
	}
	if updated.Title != "Final Entry" {
		t.Errorf("expected updated title 'Final Entry', got %q", updated.Title)
	}

	if err := loreService.Delete(created.ID); err != nil {
		t.Fatalf("Delete failed: %v", err)
	}
	if _, err := loreService.GetByID(created.ID); err == nil {
		t.Error("expected error fetching deleted lore entry, got nil")
	}
	if err := loreService.Delete(9999); err == nil {
		t.Error("expected error deleting non-existent lore entry, got nil")
	}
}
