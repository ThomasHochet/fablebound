package services_test

import (
	"testing"
	"world-builder/internal/db"
	"world-builder/internal/models"
	"world-builder/internal/services"
)

func TestLocationService_CreateAndGetByID_WithHierarchy(t *testing.T) {
	database := db.SetupTestDB(t)
	locService := services.NewLocationService(database)

	parent, err := locService.Create(models.Location{
		Name:    "Beleriand",
		Type:    "Region",
		Subtype: "Continent",
	})
	if err != nil {
		t.Fatalf("failed to create parent location: %v", err)
	}
	if parent == nil || parent.ID == 0 {
		t.Fatalf("expected valid created parent location, got %+v", parent)
	}

	child, err := locService.Create(models.Location{
		Name:     "Gondolin",
		Type:     "Settlement",
		Subtype:  "Hidden City",
		ParentID: &parent.ID,
	})
	if err != nil {
		t.Fatalf("failed to create child location: %v", err)
	}

	fetchedChild, err := locService.GetByID(child.ID)
	if err != nil {
		t.Fatalf("GetByID(child) failed: %v", err)
	}
	if fetchedChild.Name != "Gondolin" {
		t.Errorf("expected name 'Gondolin', got %q", fetchedChild.Name)
	}
	if fetchedChild.Parent == nil || fetchedChild.Parent.ID != parent.ID {
		t.Errorf("expected preloaded Parent ID %d, got %+v", parent.ID, fetchedChild.Parent)
	}

	fetchedParent, err := locService.GetByID(parent.ID)
	if err != nil {
		t.Fatalf("GetByID(parent) failed: %v", err)
	}
	if len(fetchedParent.Children) != 1 || fetchedParent.Children[0].ID != child.ID {
		t.Errorf("expected parent to preload child ID %d, got %+v", child.ID, fetchedParent.Children)
	}
}

func TestLocationService_Create_RejectsSelfParent(t *testing.T) {
	database := db.SetupTestDB(t)
	locService := services.NewLocationService(database)

	selfID := int64(42)
	created, err := locService.Create(models.Location{
		ID:       selfID,
		Name:     "Self Loop",
		ParentID: &selfID,
	})
	if err == nil {
		t.Fatal("expected error when location is its own parent, got nil")
	}
	if created != nil {
		t.Errorf("expected nil location on error, got %+v", created)
	}
}

func TestLocationService_GetByID_NotFound(t *testing.T) {
	database := db.SetupTestDB(t)
	locService := services.NewLocationService(database)

	fetched, err := locService.GetByID(9999)
	if err == nil {
		t.Fatal("expected error for non-existent location ID, got nil")
	}
	if fetched != nil {
		t.Errorf("expected nil location when not found, got %+v", fetched)
	}
}

func TestLocationService_GetRootsAndGetAll(t *testing.T) {
	database := db.SetupTestDB(t)
	locService := services.NewLocationService(database)

	rootB, err := locService.Create(models.Location{Name: "Valinor", Type: "Continent"})
	if err != nil {
		t.Fatalf("seed rootB failed: %v", err)
	}
	rootA, err := locService.Create(models.Location{Name: "Aman", Type: "Continent"})
	if err != nil {
		t.Fatalf("seed rootA failed: %v", err)
	}
	_, err = locService.Create(models.Location{
		Name:     "Valmar",
		Type:     "City",
		ParentID: &rootB.ID,
	})
	if err != nil {
		t.Fatalf("seed child failed: %v", err)
	}

	roots, err := locService.GetRoots()
	if err != nil {
		t.Fatalf("GetRoots failed: %v", err)
	}
	if len(roots) != 2 {
		t.Fatalf("expected 2 root locations, got %d", len(roots))
	}
	// Ordered by name ASC: Aman, then Valinor
	if roots[0].Name != rootA.Name || roots[1].Name != rootB.Name {
		t.Errorf("expected roots ordered [Aman, Valinor], got [%s, %s]", roots[0].Name, roots[1].Name)
	}

	all, err := locService.GetAll()
	if err != nil {
		t.Fatalf("GetAll failed: %v", err)
	}
	if len(all) != 3 {
		t.Fatalf("expected 3 total locations, got %d", len(all))
	}
}

func TestLocationService_GetTypesAndGetSubtypes(t *testing.T) {
	database := db.SetupTestDB(t)
	locService := services.NewLocationService(database)

	seeds := []models.Location{
		{Name: "Loc 1", Type: "City", Subtype: "Capital"},
		{Name: "Loc 2", Type: "City", Subtype: "Port"},
		{Name: "Loc 3", Type: "Forest", Subtype: "Capital"},
		{Name: "Loc 4", Type: "", Subtype: ""}, // Should be ignored by DISTINCT non-empty filter
	}
	for _, s := range seeds {
		if _, err := locService.Create(s); err != nil {
			t.Fatalf("seed %q failed: %v", s.Name, err)
		}
	}

	types, err := locService.GetTypes()
	if err != nil {
		t.Fatalf("GetTypes failed: %v", err)
	}
	if len(types) != 2 || types[0] != "City" || types[1] != "Forest" {
		t.Errorf("expected types [City, Forest], got %+v", types)
	}

	subtypes, err := locService.GetSubtypes()
	if err != nil {
		t.Fatalf("GetSubtypes failed: %v", err)
	}
	if len(subtypes) != 2 || subtypes[0] != "Capital" || subtypes[1] != "Port" {
		t.Errorf("expected subtypes [Capital, Port], got %+v", subtypes)
	}
}

func TestLocationService_Update_SuccessAndGuards(t *testing.T) {
	database := db.SetupTestDB(t)
	locService := services.NewLocationService(database)

	// Invalid ID
	if _, err := locService.Update(models.Location{ID: 0, Name: "No ID"}); err == nil {
		t.Error("expected error updating location with ID 0, got nil")
	}

	created, err := locService.Create(models.Location{
		Name:    "Moria",
		Type:    "Underground",
		Subtype: "Mine",
	})
	if err != nil {
		t.Fatalf("seed failed: %v", err)
	}

	// Self-parenting guard on Update
	selfID := created.ID
	badUpdate := *created
	badUpdate.ParentID = &selfID
	if _, err := locService.Update(badUpdate); err == nil {
		t.Error("expected error when updating location to be its own parent, got nil")
	}

	// Valid update
	goodUpdate := *created
	goodUpdate.Name = "Khazad-dûm"
	goodUpdate.Subtype = "Dwarven Realm"
	updated, err := locService.Update(goodUpdate)
	if err != nil {
		t.Fatalf("Update failed: %v", err)
	}
	if updated.Name != "Khazad-dûm" || updated.Subtype != "Dwarven Realm" {
		t.Errorf("unexpected updated values: %+v", updated)
	}

	// Non-existent ID
	if _, err := locService.Update(models.Location{ID: 9999, Name: "Ghost"}); err == nil {
		t.Error("expected error updating non-existent location, got nil")
	}
}

func TestLocationService_Delete_SuccessAndNotFound(t *testing.T) {
	database := db.SetupTestDB(t)
	locService := services.NewLocationService(database)

	created, err := locService.Create(models.Location{Name: "Angband", Type: "Fortress"})
	if err != nil {
		t.Fatalf("seed failed: %v", err)
	}

	if err := locService.Delete(created.ID); err != nil {
		t.Fatalf("Delete failed: %v", err)
	}
	if _, err := locService.GetByID(created.ID); err == nil {
		t.Error("expected error fetching deleted location, got nil")
	}

	if err := locService.Delete(9999); err == nil {
		t.Error("expected error deleting non-existent location, got nil")
	}
}
