package services_test

import (
	"fmt"
	"testing"
	"world-builder/internal/db"
	"world-builder/internal/models"
	"world-builder/internal/services"
)

func TestCharacterService_CreateAndGetByID_WithPreloadsAndAssociations(t *testing.T) {
	database := db.SetupTestDB(t)
	charSvc := services.NewCharacterService(database)
	lookupSvc := services.NewLookupService(database)
	traitSvc := services.NewTraitService(database)
	locSvc := services.NewLocationService(database)
	factionSvc := services.NewFactionService(database)
	loreSvc := services.NewLoreService(database)

	// Seed lookups, location, faction, traits, lore, and target character
	race, err := lookupSvc.Create("race", "Elf")
	if err != nil {
		t.Fatalf("seed race failed: %v", err)
	}
	alignment, err := lookupSvc.Create("alignment", "Lawful Good")
	if err != nil {
		t.Fatalf("seed alignment failed: %v", err)
	}
	birthplace, err := locSvc.Create(models.Location{Name: "Rivendell", Type: "Valley"})
	if err != nil {
		t.Fatalf("seed birthplace failed: %v", err)
	}
	faction, err := factionSvc.CreateFaction(models.Faction{Name: "White Council"})
	if err != nil {
		t.Fatalf("seed faction failed: %v", err)
	}
	trait, err := traitSvc.Create("personality", "Patient")
	if err != nil {
		t.Fatalf("seed personality trait failed: %v", err)
	}
	strength, err := traitSvc.Create("strength", "Keen Sight")
	if err != nil {
		t.Fatalf("seed strength failed: %v", err)
	}
	loreCat, err := loreSvc.CreateCategory("Backstory")
	if err != nil {
		t.Fatalf("seed lore category failed: %v", err)
	}

	targetChar, err := charSvc.Create(models.Character{
		Firstname: "Elrond",
		Surname:   "Half-elven",
	})
	if err != nil {
		t.Fatalf("seed target character failed: %v", err)
	}

	raceID := race.ID
	alignID := alignment.ID
	birthID := birthplace.ID
	factionID := faction.ID

	backstory, err := loreSvc.Create(models.Lore{
		Title:      "First Age Travels",
		CategoryID: loreCat.ID,
	})
	if err != nil {
		t.Fatalf("seed backstory lore failed: %v", err)
	}

	created, err := charSvc.Create(models.Character{
		Firstname:            "Glorfindel",
		Surname:              "of Rivendell",
		Nickname:             "Golden Lord",
		RaceID:               &raceID,
		AlignmentID:          &alignID,
		BirthplaceLocationID: &birthID,
		FactionID:            &factionID,
		PersonalityTraits: []models.PersonalityTrait{
			{ID: trait.ID, Label: trait.Label},
		},
		Strengths: []models.Strength{
			{ID: strength.ID, Label: strength.Label},
		},
		Relationships: []models.Relationship{
			{TargetCharID: targetChar.ID},
		},
		Backstories: []models.Lore{
			*backstory,
		},
	})
	if err != nil {
		t.Fatalf("Create character failed: %v", err)
	}
	if created == nil || created.ID == 0 {
		t.Fatalf("expected created character with non-zero ID, got %+v", created)
	}

	fetched, err := charSvc.GetByID(created.ID)
	if err != nil {
		t.Fatalf("GetByID failed: %v", err)
	}
	if fetched.Firstname != "Glorfindel" || fetched.Surname != "of Rivendell" {
		t.Errorf("unexpected character name: %+v", fetched)
	}
	if fetched.Race == nil || fetched.Race.ID != race.ID {
		t.Errorf("expected preloaded Race ID %d, got %+v", race.ID, fetched.Race)
	}
	if fetched.Alignment == nil || fetched.Alignment.ID != alignment.ID {
		t.Errorf("expected preloaded Alignment ID %d, got %+v", alignment.ID, fetched.Alignment)
	}
	if fetched.Birthplace == nil || fetched.Birthplace.ID != birthplace.ID {
		t.Errorf("expected preloaded Birthplace ID %d, got %+v", birthplace.ID, fetched.Birthplace)
	}
	if fetched.Faction == nil || fetched.Faction.ID != faction.ID {
		t.Errorf("expected preloaded Faction ID %d, got %+v", faction.ID, fetched.Faction)
	}
	if len(fetched.PersonalityTraits) != 1 || fetched.PersonalityTraits[0].ID != trait.ID {
		t.Errorf("expected 1 PersonalityTrait ID %d, got %+v", trait.ID, fetched.PersonalityTraits)
	}
	if len(fetched.Strengths) != 1 || fetched.Strengths[0].ID != strength.ID {
		t.Errorf("expected 1 Strength ID %d, got %+v", strength.ID, fetched.Strengths)
	}
	if len(fetched.Relationships) != 1 || fetched.Relationships[0].TargetCharID != targetChar.ID {
		t.Errorf("expected 1 Relationship to target %d, got %+v", targetChar.ID, fetched.Relationships)
	}
	if len(fetched.Backstories) != 1 || fetched.Backstories[0].ID != backstory.ID {
		t.Errorf("expected 1 Backstory ID %d, got %+v", backstory.ID, fetched.Backstories)
	}
}

func TestCharacterService_GetByID_NotFound(t *testing.T) {
	database := db.SetupTestDB(t)
	charSvc := services.NewCharacterService(database)

	fetched, err := charSvc.GetByID(9999)
	if err == nil {
		t.Fatal("expected error fetching non-existent character, got nil")
	}
	if fetched != nil {
		t.Errorf("expected nil character when not found, got %+v", fetched)
	}
}

func TestCharacterService_GetAllCharacters_Limit(t *testing.T) {
	database := db.SetupTestDB(t)
	charSvc := services.NewCharacterService(database)

	for i := 1; i <= 3; i++ {
		if _, err := charSvc.Create(models.Character{
			Firstname: fmt.Sprintf("Char%d", i),
			Surname:   "Test",
		}); err != nil {
			t.Fatalf("seed character %d failed: %v", i, err)
		}
	}

	limited, err := charSvc.GetAllCharacters(2)
	if err != nil {
		t.Fatalf("GetAllCharacters(2) failed: %v", err)
	}
	if len(limited) != 2 {
		t.Errorf("expected 2 characters with limit=2, got %d", len(limited))
	}

	// limit <= 0 defaults to 200; with 3 seeded records it should return all 3
	defaultLimit, err := charSvc.GetAllCharacters(0)
	if err != nil {
		t.Fatalf("GetAllCharacters(0) failed: %v", err)
	}
	if len(defaultLimit) != 3 {
		t.Errorf("expected 3 characters with default limit, got %d", len(defaultLimit))
	}
}

func TestCharacterService_GetAllCharacterAsLookup(t *testing.T) {
	database := db.SetupTestDB(t)
	charSvc := services.NewCharacterService(database)

	_, err := charSvc.Create(models.Character{
		Firstname: "Boromir",
		Surname:   "of Gondor",
		Nickname:  "Captain of the White Tower",
	})
	if err != nil {
		t.Fatalf("seed Boromir failed: %v", err)
	}

	_, err = charSvc.Create(models.Character{
		Firstname: "Faramir",
		Surname:   "of Gondor",
	})
	if err != nil {
		t.Fatalf("seed Faramir failed: %v", err)
	}

	lookups, err := charSvc.GetAllCharacterAsLookup()
	if err != nil {
		t.Fatalf("GetAllCharacterAsLookup failed: %v", err)
	}
	if len(lookups) != 2 {
		t.Fatalf("expected 2 lookup items, got %d", len(lookups))
	}

	// Ordered by firstname ASC, surname ASC -> Boromir first, then Faramir
	if lookups[0].Label != "Boromir of Gondor (Captain of the White Tower)" {
		t.Errorf("unexpected Boromir lookup label: %q", lookups[0].Label)
	}
	if lookups[1].Label != "Faramir of Gondor" {
		t.Errorf("unexpected Faramir lookup label: %q", lookups[1].Label)
	}
}

func TestCharacterService_Update_ReplacesTraits(t *testing.T) {
	database := db.SetupTestDB(t)
	charSvc := services.NewCharacterService(database)
	traitSvc := services.NewTraitService(database)

	if _, err := charSvc.Update(models.Character{ID: 0, Firstname: "Invalid"}); err == nil {
		t.Error("expected error updating character with ID 0, got nil")
	}

	traitA, _ := traitSvc.Create("personality", "Bold")
	traitB, _ := traitSvc.Create("personality", "Cautious")

	created, err := charSvc.Create(models.Character{
		Firstname:         "Eowyn",
		Surname:           "of Rohan",
		PersonalityTraits: []models.PersonalityTrait{{ID: traitA.ID, Label: traitA.Label}},
	})
	if err != nil {
		t.Fatalf("seed failed: %v", err)
	}

	updatedInput := *created
	updatedInput.Nickname = "Shieldmaiden"
	updatedInput.PersonalityTraits = []models.PersonalityTrait{{ID: traitB.ID, Label: traitB.Label}}

	updated, err := charSvc.Update(updatedInput)
	if err != nil {
		t.Fatalf("Update failed: %v", err)
	}
	if updated.Nickname != "Shieldmaiden" {
		t.Errorf("expected nickname 'Shieldmaiden', got %q", updated.Nickname)
	}
	if len(updated.PersonalityTraits) != 1 || updated.PersonalityTraits[0].ID != traitB.ID {
		t.Errorf("expected PersonalityTraits replaced with traitB (%d), got %+v", traitB.ID, updated.PersonalityTraits)
	}
}

func TestCharacterService_Delete_SuccessAndGuards(t *testing.T) {
	database := db.SetupTestDB(t)
	charSvc := services.NewCharacterService(database)

	if err := charSvc.Delete(0); err == nil {
		t.Error("expected error deleting character with ID 0, got nil")
	}

	created, err := charSvc.Create(models.Character{
		Firstname: "Gimli",
		Surname:   "Son of Gloin",
	})
	if err != nil {
		t.Fatalf("seed failed: %v", err)
	}

	if err := charSvc.Delete(created.ID); err != nil {
		t.Fatalf("Delete failed: %v", err)
	}
	if _, err := charSvc.GetByID(created.ID); err == nil {
		t.Error("expected error fetching deleted character, got nil")
	}
	if err := charSvc.Delete(9999); err == nil {
		t.Error("expected error deleting non-existent character, got nil")
	}
}
