package services_test

import (
	"testing"
	"world-builder/internal/db"
	"world-builder/internal/models"
	"world-builder/internal/services"
)

func lookupModelForCategory(t *testing.T, category string) interface{} {
	t.Helper()
	switch category {
	case "gender":
		return &models.Gender{}
	case "race":
		return &models.Race{}
	case "occupation":
		return &models.Occupation{}
	case "alignment":
		return &models.Alignment{}
	case "status":
		return &models.Status{}
	default:
		t.Fatalf("unexpected lookup category %q", category)
		return nil
	}
}

func traitModelForCategory(t *testing.T, category string) interface{} {
	t.Helper()
	switch category {
	case "personality":
		return &models.PersonalityTrait{}
	case "strength":
		return &models.Strength{}
	case "flaw":
		return &models.Flaw{}
	case "weakness":
		return &models.Weakness{}
	case "fear":
		return &models.Fear{}
	default:
		t.Fatalf("unexpected trait category %q", category)
		return nil
	}
}

// --- LOOKUP SERVICE TESTS ---

func TestLookupService_CreateAndFetch_AllCategories(t *testing.T) {
	database := db.SetupTestDB(t)
	lookupService := services.NewLookupService(database)

	categories := []string{"gender", "race", "occupation", "alignment", "status"}

	for _, category := range categories {
		t.Run(category, func(t *testing.T) {
			label := "Sample_" + category

			created, err := lookupService.Create(category, label)
			if err != nil {
				t.Fatalf("Create(%q) failed: %v", category, err)
			}
			if created == nil {
				t.Fatalf("Create(%q) returned nil item", category)
			}
			if created.ID == 0 {
				t.Errorf("Create(%q) expected non-zero ID", category)
			}
			if created.Label != label {
				t.Errorf("Create(%q) expected label %q, got %q", category, label, created.Label)
			}

			var count int64
			model := lookupModelForCategory(t, category)
			if err := database.Model(model).Where("id = ? AND label = ?", created.ID, label).Count(&count).Error; err != nil {
				t.Fatalf("DB count failed for %q: %v", category, err)
			}
			if count != 1 {
				t.Errorf("expected 1 DB record for %q, got %d", category, count)
			}

			fetched, err := lookupService.Fetch(created.ID, category)
			if err != nil {
				t.Fatalf("Fetch(%q) failed: %v", category, err)
			}
			if fetched.Label != label {
				t.Errorf("Fetch(%q) expected label %q, got %q", category, label, fetched.Label)
			}
		})
	}
}

func TestLookupService_FetchAll_OrderedByLabel(t *testing.T) {
	database := db.SetupTestDB(t)
	lookupService := services.NewLookupService(database)

	empty, err := lookupService.FetchAll("race")
	if err != nil {
		t.Fatalf("FetchAll on empty DB failed: %v", err)
	}
	if len(empty) != 0 {
		t.Errorf("expected 0 items on empty DB, got %d", len(empty))
	}

	_, err = lookupService.Create("race", "Human")
	if err != nil {
		t.Fatalf("seed Human failed: %v", err)
	}
	_, err = lookupService.Create("race", "Elf")
	if err != nil {
		t.Fatalf("seed Elf failed: %v", err)
	}

	items, err := lookupService.FetchAll("race")
	if err != nil {
		t.Fatalf("FetchAll failed: %v", err)
	}
	if len(items) != 2 {
		t.Fatalf("expected 2 items, got %d", len(items))
	}
	if items[0].Label != "Elf" || items[1].Label != "Human" {
		t.Errorf("expected items ordered [Elf, Human], got [%s, %s]", items[0].Label, items[1].Label)
	}
}

func TestLookupService_Update_Success(t *testing.T) {
	database := db.SetupTestDB(t)
	lookupService := services.NewLookupService(database)

	created, err := lookupService.Create("alignment", "Chaotic Neutral")
	if err != nil {
		t.Fatalf("Create failed: %v", err)
	}

	updated, err := lookupService.Update(created.ID, "True Neutral", "alignment")
	if err != nil {
		t.Fatalf("Update failed: %v", err)
	}
	if updated.Label != "True Neutral" {
		t.Errorf("expected updated label 'True Neutral', got %q", updated.Label)
	}

	fetched, err := lookupService.Fetch(created.ID, "alignment")
	if err != nil {
		t.Fatalf("Fetch after Update failed: %v", err)
	}
	if fetched.Label != "True Neutral" {
		t.Errorf("expected persisted label 'True Neutral', got %q", fetched.Label)
	}
}

func TestLookupService_Delete_Success(t *testing.T) {
	database := db.SetupTestDB(t)
	lookupService := services.NewLookupService(database)

	created, err := lookupService.Create("status", "Active")
	if err != nil {
		t.Fatalf("Create failed: %v", err)
	}

	if err := lookupService.Delete(created.ID, "status"); err != nil {
		t.Fatalf("Delete failed: %v", err)
	}

	if _, err := lookupService.Fetch(created.ID, "status"); err == nil {
		t.Fatal("expected error fetching deleted lookup item, got nil")
	}
}

func TestLookupService_InvalidCategory(t *testing.T) {
	database := db.SetupTestDB(t)
	lookupService := services.NewLookupService(database)

	if _, err := lookupService.Create("unknown", "Test"); err == nil {
		t.Error("expected error for invalid category on Create, got nil")
	}
	if _, err := lookupService.Fetch(1, "unknown"); err == nil {
		t.Error("expected error for invalid category on Fetch, got nil")
	}
	if _, err := lookupService.FetchAll("unknown"); err == nil {
		t.Error("expected error for invalid category on FetchAll, got nil")
	}
	if _, err := lookupService.Update(1, "Test", "unknown"); err == nil {
		t.Error("expected error for invalid category on Update, got nil")
	}
	if err := lookupService.Delete(1, "unknown"); err == nil {
		t.Error("expected error for invalid category on Delete, got nil")
	}
}

// --- TRAIT SERVICE TESTS ---

func TestTraitService_CreateAndFetch_AllCategories(t *testing.T) {
	database := db.SetupTestDB(t)
	traitService := services.NewTraitService(database)

	categories := []string{"personality", "strength", "flaw", "weakness", "fear"}

	for _, category := range categories {
		t.Run(category, func(t *testing.T) {
			label := "Sample_" + category

			created, err := traitService.Create(category, label)
			if err != nil {
				t.Fatalf("Create(%q) failed: %v", category, err)
			}
			if created == nil {
				t.Fatalf("Create(%q) returned nil item", category)
			}
			if created.ID == 0 {
				t.Errorf("Create(%q) expected non-zero ID", category)
			}
			if created.Label != label {
				t.Errorf("Create(%q) expected label %q, got %q", category, label, created.Label)
			}

			var count int64
			model := traitModelForCategory(t, category)
			if err := database.Model(model).Where("id = ? AND label = ?", created.ID, label).Count(&count).Error; err != nil {
				t.Fatalf("DB count failed for %q: %v", category, err)
			}
			if count != 1 {
				t.Errorf("expected 1 DB record for %q, got %d", category, count)
			}

			fetched, err := traitService.Fetch(category, created.ID)
			if err != nil {
				t.Fatalf("Fetch(%q) failed: %v", category, err)
			}
			if fetched.Label != label {
				t.Errorf("Fetch(%q) expected label %q, got %q", category, label, fetched.Label)
			}
		})
	}
}

func TestTraitService_FetchAll_OrderedByLabel(t *testing.T) {
	database := db.SetupTestDB(t)
	traitService := services.NewTraitService(database)

	empty, err := traitService.FetchAll("fear")
	if err != nil {
		t.Fatalf("FetchAll on empty DB failed: %v", err)
	}
	if len(empty) != 0 {
		t.Errorf("expected 0 items on empty DB, got %d", len(empty))
	}

	_, err = traitService.Create("fear", "Spiders")
	if err != nil {
		t.Fatalf("seed Spiders failed: %v", err)
	}
	_, err = traitService.Create("fear", "Darkness")
	if err != nil {
		t.Fatalf("seed Darkness failed: %v", err)
	}

	items, err := traitService.FetchAll("fear")
	if err != nil {
		t.Fatalf("FetchAll failed: %v", err)
	}
	if len(items) != 2 {
		t.Fatalf("expected 2 items, got %d", len(items))
	}
	if items[0].Label != "Darkness" || items[1].Label != "Spiders" {
		t.Errorf("expected items ordered [Darkness, Spiders], got [%s, %s]", items[0].Label, items[1].Label)
	}
}

func TestTraitService_Update_SuccessAndNotFound(t *testing.T) {
	database := db.SetupTestDB(t)
	traitService := services.NewTraitService(database)

	created, err := traitService.Create("strength", "Iron Will")
	if err != nil {
		t.Fatalf("Create failed: %v", err)
	}

	updated, err := traitService.Update(created.ID, "Unbreakable Will", "strength")
	if err != nil {
		t.Fatalf("Update failed: %v", err)
	}
	if updated.Label != "Unbreakable Will" {
		t.Errorf("expected updated label 'Unbreakable Will', got %q", updated.Label)
	}

	if _, err := traitService.Update(9999, "Ghost Trait", "strength"); err == nil {
		t.Fatal("expected error updating non-existent trait ID, got nil")
	}
}

func TestTraitService_Delete_Success(t *testing.T) {
	database := db.SetupTestDB(t)
	traitService := services.NewTraitService(database)

	created, err := traitService.Create("flaw", "Arrogance")
	if err != nil {
		t.Fatalf("Create failed: %v", err)
	}

	if err := traitService.Delete("flaw", created.ID); err != nil {
		t.Fatalf("Delete failed: %v", err)
	}

	if _, err := traitService.Fetch("flaw", created.ID); err == nil {
		t.Fatal("expected error fetching deleted trait item, got nil")
	}
}

func TestTraitService_InvalidCategory(t *testing.T) {
	database := db.SetupTestDB(t)
	traitService := services.NewTraitService(database)

	if _, err := traitService.Create("unknown", "Test"); err == nil {
		t.Error("expected error for invalid category on Create, got nil")
	}
	if _, err := traitService.Fetch("unknown", 1); err == nil {
		t.Error("expected error for invalid category on Fetch, got nil")
	}
	if _, err := traitService.FetchAll("unknown"); err == nil {
		t.Error("expected error for invalid category on FetchAll, got nil")
	}
	if _, err := traitService.Update(1, "Test", "unknown"); err == nil {
		t.Error("expected error for invalid category on Update, got nil")
	}
	if err := traitService.Delete("unknown", 1); err == nil {
		t.Error("expected error for invalid category on Delete, got nil")
	}
}
