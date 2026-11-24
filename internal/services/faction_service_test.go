package services_test

import (
	"testing"
	"world-builder/internal/db"
	"world-builder/internal/models"
	"world-builder/internal/services"
)

func seedFaction(t *testing.T, svc *services.FactionService, name string) *models.Faction {
	t.Helper()
	f, err := svc.CreateFaction(models.Faction{Name: name})
	if err != nil {
		t.Fatalf("failed to seed faction %q: %v", name, err)
	}
	return f
}

// --- FACTION TESTS ---

func TestFactionService_CreateAndGetFaction(t *testing.T) {
	database := db.SetupTestDB(t)
	factionSvc := services.NewFactionService(database)

	created, err := factionSvc.CreateFaction(models.Faction{Name: "Iron Guild"})
	if err != nil {
		t.Fatalf("CreateFaction failed: %v", err)
	}
	if created == nil || created.ID == 0 {
		t.Fatalf("expected created faction with non-zero ID, got %+v", created)
	}
	if created.Name != "Iron Guild" {
		t.Errorf("expected name 'Iron Guild', got %q", created.Name)
	}

	fetched, err := factionSvc.GetFaction(created.ID)
	if err != nil {
		t.Fatalf("GetFaction failed: %v", err)
	}
	if fetched.ID != created.ID || fetched.Name != "Iron Guild" {
		t.Errorf("unexpected fetched faction: %+v", fetched)
	}

	missing, err := factionSvc.GetFaction(9999)
	if err == nil {
		t.Fatal("expected error for non-existent faction ID, got nil")
	}
	if missing != nil {
		t.Errorf("expected nil faction when not found, got %+v", missing)
	}
}

func TestFactionService_GetAllFactions_OrderedByName(t *testing.T) {
	database := db.SetupTestDB(t)
	factionSvc := services.NewFactionService(database)

	empty, err := factionSvc.GetAllFactions()
	if err != nil {
		t.Fatalf("GetAllFactions on empty DB failed: %v", err)
	}
	if len(empty) != 0 {
		t.Errorf("expected 0 factions, got %d", len(empty))
	}

	seedFaction(t, factionSvc, "Zephryn Syndicate")
	seedFaction(t, factionSvc, "Dawn Wardens")

	all, err := factionSvc.GetAllFactions()
	if err != nil {
		t.Fatalf("GetAllFactions failed: %v", err)
	}
	if len(all) != 2 {
		t.Fatalf("expected 2 factions, got %d", len(all))
	}
	if all[0].Name != "Dawn Wardens" || all[1].Name != "Zephryn Syndicate" {
		t.Errorf("expected factions ordered [Dawn Wardens, Zephryn Syndicate], got [%s, %s]", all[0].Name, all[1].Name)
	}
}

func TestFactionService_UpdateAndDeleteFaction(t *testing.T) {
	database := db.SetupTestDB(t)
	factionSvc := services.NewFactionService(database)

	if _, err := factionSvc.UpdateFaction(models.Faction{ID: 0, Name: "No ID"}); err == nil {
		t.Error("expected error updating faction with ID 0, got nil")
	}

	created := seedFaction(t, factionSvc, "Old Order")

	updatedInput := *created
	updatedInput.Name = "New Order"
	updated, err := factionSvc.UpdateFaction(updatedInput)
	if err != nil {
		t.Fatalf("UpdateFaction failed: %v", err)
	}
	if updated.Name != "New Order" {
		t.Errorf("expected updated name 'New Order', got %q", updated.Name)
	}

	if err := factionSvc.DeleteFaction(created.ID); err != nil {
		t.Fatalf("DeleteFaction failed: %v", err)
	}
	if _, err := factionSvc.GetFaction(created.ID); err == nil {
		t.Error("expected error fetching deleted faction, got nil")
	}
	if err := factionSvc.DeleteFaction(9999); err == nil {
		t.Error("expected error deleting non-existent faction, got nil")
	}
}

// --- AFFILIATION & AGGREGATE QUERY TESTS ---

func TestFactionService_AffiliationCRUDAndQueries(t *testing.T) {
	database := db.SetupTestDB(t)
	factionSvc := services.NewFactionService(database)

	f1 := seedFaction(t, factionSvc, "Vanguard")
	f2 := seedFaction(t, factionSvc, "Lonely Faction")

	f1ID := f1.ID
	createdAff, err := factionSvc.CreateAffiliation(models.Affiliation{
		Name:      "Scout Rank",
		FactionID: &f1ID,
	})
	if err != nil {
		t.Fatalf("CreateAffiliation failed: %v", err)
	}
	if createdAff == nil || createdAff.ID == 0 {
		t.Fatalf("expected created affiliation with non-zero ID, got %+v", createdAff)
	}

	// GetAffiliation preloads Faction
	fetchedAff, err := factionSvc.GetAffiliation(createdAff.ID)
	if err != nil {
		t.Fatalf("GetAffiliation failed: %v", err)
	}
	if fetchedAff.Faction == nil || fetchedAff.Faction.ID != f1.ID {
		t.Errorf("expected preloaded Faction ID %d, got %+v", f1.ID, fetchedAff.Faction)
	}

	// GetAffiliationsFaction
	f1Affs, err := factionSvc.GetAffiliationsFaction(f1.ID)
	if err != nil {
		t.Fatalf("GetAffiliationsFaction failed: %v", err)
	}
	if len(f1Affs) != 1 || f1Affs[0].Name != "Scout Rank" {
		t.Errorf("expected 1 affiliation for f1, got %+v", f1Affs)
	}

	// GetAllAffiliations
	allAffs, err := factionSvc.GetAllAffiliations()
	if err != nil {
		t.Fatalf("GetAllAffiliations failed: %v", err)
	}
	if len(allAffs) != 1 {
		t.Fatalf("expected 1 total affiliation, got %d", len(allAffs))
	}

	// GetFactionAffiliations (FactionDetails)
	details, err := factionSvc.GetFactionAffiliations(f1.ID)
	if err != nil {
		t.Fatalf("GetFactionAffiliations failed: %v", err)
	}
	if details.Faction.ID != f1.ID || len(details.Affiliations) != 1 {
		t.Errorf("unexpected FactionDetails: %+v", details)
	}
	if _, err := factionSvc.GetFactionAffiliations(9999); err == nil {
		t.Error("expected error from GetFactionAffiliations for missing faction ID, got nil")
	}

	// GetFactionsWithAffiliations should only return f1, not f2
	withAffs, err := factionSvc.GetFactionsWithAffiliations()
	if err != nil {
		t.Fatalf("GetFactionsWithAffiliations failed: %v", err)
	}
	if len(withAffs) != 1 || withAffs[0].ID != f1.ID {
		t.Errorf("expected only faction %d (%s) to have affiliations, got %+v", f1.ID, f2.Name, withAffs)
	}

	// UpdateAffiliation validation & success
	if _, err := factionSvc.UpdateAffiliation(models.Affiliation{ID: 0, Name: "Invalid"}); err == nil {
		t.Error("expected error updating affiliation with ID 0, got nil")
	}
	updatedInput := *createdAff
	updatedInput.Name = "Captain Rank"
	updatedAff, err := factionSvc.UpdateAffiliation(updatedInput)
	if err != nil {
		t.Fatalf("UpdateAffiliation failed: %v", err)
	}
	if updatedAff.Name != "Captain Rank" {
		t.Errorf("expected updated affiliation name 'Captain Rank', got %q", updatedAff.Name)
	}

	// DeleteAffiliation success & not found
	if err := factionSvc.DeleteAffiliation(createdAff.ID); err != nil {
		t.Fatalf("DeleteAffiliation failed: %v", err)
	}
	if _, err := factionSvc.GetAffiliation(createdAff.ID); err == nil {
		t.Error("expected error fetching deleted affiliation, got nil")
	}
	if err := factionSvc.DeleteAffiliation(9999); err == nil {
		t.Error("expected error deleting non-existent affiliation, got nil")
	}
}
