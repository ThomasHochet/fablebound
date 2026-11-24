package services_test

import (
	"testing"
	"world-builder/internal/db"
	"world-builder/internal/models"
	"world-builder/internal/services"
)

func TestRelationshipService_CreateAndGetByID_WithPreloads(t *testing.T) {
	database := db.SetupTestDB(t)
	relService := services.NewRelationshipService(database)

	sourceChar := models.Character{Firstname: "Aragorn", Surname: "Elessar"}
	targetChar := models.Character{Firstname: "Arwen", Surname: "Undómiel"}
	if err := database.Create(&sourceChar).Error; err != nil {
		t.Fatalf("failed to seed source character: %v", err)
	}
	if err := database.Create(&targetChar).Error; err != nil {
		t.Fatalf("failed to seed target character: %v", err)
	}

	created, err := relService.Create(models.Relationship{
		SourceCharID: sourceChar.ID,
		TargetCharID: targetChar.ID,
	})
	if err != nil {
		t.Fatalf("Create failed: %v", err)
	}
	if created == nil || created.ID == 0 {
		t.Fatalf("expected created relationship with non-zero ID, got %+v", created)
	}

	fetched, err := relService.GetByID(created.ID)
	if err != nil {
		t.Fatalf("GetByID failed: %v", err)
	}
	if fetched.SourceCharacter.ID != sourceChar.ID {
		t.Errorf("expected preloaded SourceCharacter ID %d, got %d", sourceChar.ID, fetched.SourceCharacter.ID)
	}
	if fetched.TargetCharacter.ID != targetChar.ID {
		t.Errorf("expected preloaded TargetCharacter ID %d, got %d", targetChar.ID, fetched.TargetCharacter.ID)
	}
}

func TestRelationshipService_Create_RejectsSelfRelationship(t *testing.T) {
	database := db.SetupTestDB(t)
	relService := services.NewRelationshipService(database)

	created, err := relService.Create(models.Relationship{
		SourceCharID: 7,
		TargetCharID: 7,
	})
	if err == nil {
		t.Fatal("expected error when SourceCharID equals TargetCharID, got nil")
	}
	if created != nil {
		t.Errorf("expected nil relationship on error, got %+v", created)
	}
}

func TestRelationshipService_GetByID_NotFound(t *testing.T) {
	database := db.SetupTestDB(t)
	relService := services.NewRelationshipService(database)

	fetched, err := relService.GetByID(9999)
	if err == nil {
		t.Fatal("expected error fetching non-existent relationship, got nil")
	}
	if fetched != nil {
		t.Errorf("expected nil relationship when not found, got %+v", fetched)
	}
}

func TestRelationshipService_GetByCharacterID_SourceAndTarget(t *testing.T) {
	database := db.SetupTestDB(t)
	relService := services.NewRelationshipService(database)

	charA := models.Character{Firstname: "Frodo", Surname: "Baggins"}
	charB := models.Character{Firstname: "Samwise", Surname: "Gamgee"}
	charC := models.Character{Firstname: "Gollum", Surname: ""}
	for _, c := range []*models.Character{&charA, &charB, &charC} {
		if err := database.Create(c).Error; err != nil {
			t.Fatalf("failed to seed character %q: %v", c.Firstname, err)
		}
	}

	// charA as source
	if _, err := relService.Create(models.Relationship{
		SourceCharID: charA.ID,
		TargetCharID: charB.ID,
	}); err != nil {
		t.Fatalf("seed rel 1 failed: %v", err)
	}

	// charA as target
	if _, err := relService.Create(models.Relationship{
		SourceCharID: charC.ID,
		TargetCharID: charA.ID,
	}); err != nil {
		t.Fatalf("seed rel 2 failed: %v", err)
	}

	// Unrelated pair (charB -> charC)
	if _, err := relService.Create(models.Relationship{
		SourceCharID: charB.ID,
		TargetCharID: charC.ID,
	}); err != nil {
		t.Fatalf("seed rel 3 failed: %v", err)
	}

	rels, err := relService.GetByCharacterID(charA.ID)
	if err != nil {
		t.Fatalf("GetByCharacterID failed: %v", err)
	}
	if len(rels) != 2 {
		t.Fatalf("expected 2 relationships involving charA, got %d", len(rels))
	}

	// Empty result for character with no relationships
	none, err := relService.GetByCharacterID(9999)
	if err != nil {
		t.Fatalf("GetByCharacterID for unused ID failed: %v", err)
	}
	if len(none) != 0 {
		t.Errorf("expected 0 relationships, got %d", len(none))
	}
}

func TestRelationshipService_Update_SuccessAndGuards(t *testing.T) {
	database := db.SetupTestDB(t)
	relService := services.NewRelationshipService(database)

	if _, err := relService.Update(models.Relationship{ID: 0}); err == nil {
		t.Error("expected error updating relationship with ID 0, got nil")
	}

	char1 := models.Character{Firstname: "Legolas", Surname: "Greenleaf"}
	char2 := models.Character{Firstname: "Gimli", Surname: "Gloin's son"}
	_ = database.Create(&char1).Error
	_ = database.Create(&char2).Error

	created, err := relService.Create(models.Relationship{
		SourceCharID: char1.ID,
		TargetCharID: char2.ID,
	})
	if err != nil {
		t.Fatalf("seed failed: %v", err)
	}

	// Updating fields on the existing record
	updatedInput := *created
	updated, err := relService.Update(updatedInput)
	if err != nil {
		t.Fatalf("Update failed: %v", err)
	}
	if updated == nil || updated.ID != created.ID {
		t.Errorf("expected updated relationship with ID %d, got %+v", created.ID, updated)
	}

	// Note: RelationshipService.Update currently uses s.db.Save(&rel), which upserts
	// if ID does not exist unless switched to .Model(&models.Relationship{}).Where(...).Updates(...)
}

func TestRelationshipService_Delete_SuccessAndNotFound(t *testing.T) {
	database := db.SetupTestDB(t)
	relService := services.NewRelationshipService(database)

	char1 := models.Character{Firstname: "Gandalf", Surname: "Grey"}
	char2 := models.Character{Firstname: "Saruman", Surname: "White"}
	_ = database.Create(&char1).Error
	_ = database.Create(&char2).Error

	created, err := relService.Create(models.Relationship{
		SourceCharID: char1.ID,
		TargetCharID: char2.ID,
	})
	if err != nil {
		t.Fatalf("seed failed: %v", err)
	}

	if err := relService.Delete(created.ID); err != nil {
		t.Fatalf("Delete failed: %v", err)
	}
	if _, err := relService.GetByID(created.ID); err == nil {
		t.Error("expected error fetching deleted relationship, got nil")
	}

	if err := relService.Delete(9999); err == nil {
		t.Error("expected error deleting non-existent relationship, got nil")
	}
}
