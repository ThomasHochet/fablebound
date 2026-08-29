package main

import (
	"context"
	"fmt"
	"world-builder/internal/models"
	"world-builder/internal/services"

	"github.com/wailsapp/wails/v3/pkg/application"
	"gorm.io/gorm"
)

// App struct
type App struct {
	wailsApp            *application.App
	db                  *gorm.DB
	articleService      *services.ArticleService
	characterService    *services.CharacterService
	factionService      *services.FactionService
	locationService     *services.LocationService
	lookupService       *services.LookupService
	loreService         *services.LoreService
	relationshipService *services.RelationshipService
	traitService        *services.TraitService
}

// NewApp creates a new App application struct
func NewApp(database *gorm.DB) *App {
	return &App{
		db:                  database,
		articleService:      services.NewArticleService(database),
		characterService:    services.NewCharacterService(database),
		factionService:      services.NewFactionService(database),
		locationService:     services.NewLocationService(database),
		lookupService:       services.NewLookupService(database),
		loreService:         services.NewLoreService(database),
		relationshipService: services.NewRelationshipService(database),
		traitService:        services.NewTraitService(database),
	}
}

// ServiceStartup replaces the v2 startup method
func (a *App) ServiceStartup(ctx context.Context, options application.ServiceOptions) error {
	// Keep a reference to the main application instance
	a.wailsApp = application.Get()
	return nil
}

// Open window editor based on the section
func (a *App) OpenEditorWindow(section string) {
	if a.wailsApp == nil {
		a.wailsApp = application.Get()
	}

	a.wailsApp.Window.NewWithOptions(application.WebviewWindowOptions{
		Title:  fmt.Sprintf("World Builder - %s Editor", section),
		Width:  1920,
		Height: 1080,
		URL:    fmt.Sprintf("#/editor/%s", section),
	})
}

// Articles

func (a *App) CreateArticle(title string, description string) (*models.Article, error) {
	return a.articleService.Create(title, description, "General")
}

func (a *App) FetchAllArticles() ([]models.Article, error) {
	return a.articleService.GetAll()
}

func (a *App) FetchArticle(id int64) (*models.Article, error) {
	return a.articleService.GetArticle(id)
}

func (a *App) UpdateArticle(id int64, title, description string) (*models.Article, error) {
	return a.articleService.Update(id, title, description)
}

func (a *App) DeleteArticle(id int64) error {
	return a.articleService.Delete(id)
}

// Characters

func (a *App) SaveCharacter(char models.Character) (*models.Character, error) {
	if char.ID == 0 {
		return a.characterService.Create(char)
	}
	return a.characterService.Update(char)
}

func (a *App) FetchCharacter(id int64) (*models.Character, error) {
	return a.characterService.GetByID(id)
}

func (a *App) FetchAllCharacters(limit int) ([]models.Character, error) {
	return a.characterService.GetAllCharacters(limit)
}

func (a *App) DeleteCharacter(id int64) error {
	return a.characterService.Delete(id)
}

// Factions

func (a *App) SaveFaction(faction models.Faction) (*models.Faction, error) {
	if faction.ID == 0 {
		return a.factionService.CreateFaction(faction)
	}
	return a.factionService.UpdateFaction(faction)
}

func (a *App) FetchFaction(id int64) (*models.Faction, error) {
	return a.factionService.GetFaction(id)
}

func (a *App) FetchAllFactions() ([]models.Faction, error) {
	return a.factionService.GetAllFactions()
}

func (a *App) DeleteFaction(id int64) error {
	return a.factionService.DeleteFaction(id)
}

// Affiliations

func (a *App) SaveAffiliation(affiliation models.Affiliation) (*models.Affiliation, error) {
	if affiliation.ID == 0 {
		return a.factionService.CreateAffiliation(affiliation)
	}
	return a.factionService.UpdateAffiliation(affiliation)
}

func (a *App) FetchAffiliation(id int64) (*models.Affiliation, error) {
	return a.factionService.GetAffiliation(id)
}

func (a *App) FetchAllAffiliations() ([]models.Affiliation, error) {
	return a.factionService.GetAllAffiliations()
}

func (a *App) DeleteAffiliation(id int64) error {
	return a.factionService.DeleteAffiliation(id)
}

// LocationsS

func (a *App) SaveLocation(location models.Location) (*models.Location, error) {
	if location.ID == 0 {
		return a.locationService.Create(location)
	}
	return a.locationService.Update(location)
}

func (a *App) FetchLocation(id int64) (*models.Location, error) {
	return a.locationService.GetByID(id)
}

func (a *App) FetchRoots() ([]models.Location, error) {
	return a.locationService.GetRoots()
}

func (a *App) FetchAllLocations() ([]models.Location, error) {
	return a.locationService.GetAll()
}

func (a *App) FetchLocationTypes() ([]string, error) {
	return a.locationService.GetTypes()
}

func (a *App) FetchLocationSubtypes() ([]string, error) {
	return a.locationService.GetSubtypes()
}

func (a *App) DeleteLocation(id int64) error {
	return a.locationService.Delete(id)
}

// Lookup

func (a *App) CreateLookup(category, label string) (*services.LookupItem, error) {
	return a.lookupService.Create(category, label)
}

func (a *App) SaveLookup(id int64, label, category string) (*services.LookupItem, error) {
	return a.lookupService.Update(id, label, category)
}

func (a *App) FetchLookup(id int64, category string) (*services.LookupItem, error) {
	return a.lookupService.Fetch(id, category)
}

func (a *App) FetchAllLookup(category string) ([]services.LookupItem, error) {
	return a.lookupService.FetchAll(category)
}

func (a *App) DeleteLookup(id int64, category string) error {
	return a.lookupService.Delete(id, category)
}

// Lore

func (a *App) SaveLore(lore models.Lore) (*models.Lore, error) {
	if lore.ID == 0 {
		if lore.CategoryID == 0 && lore.LoreCategory != nil && lore.LoreCategory.Label != "" {
			newCategory, err := a.loreService.CreateCategory(lore.LoreCategory.Label)
			if err != nil {
				return nil, err
			}
			lore.CategoryID = newCategory.ID

		}
		return a.loreService.Create(lore)
	}
	return a.loreService.Update(lore)
}

func (a *App) FetchLore(id int64) (*models.Lore, error) {
	return a.loreService.GetByID(id)
}

func (a *App) FetchAllLore() ([]models.Lore, error) {
	return a.loreService.GetAll()
}

func (a *App) FetchLoreByCategory(id int64) ([]models.Lore, error) {
	return a.loreService.GetByCategoryID(id)
}

func (a *App) FetchCharacterLore(id int64) ([]models.Lore, error) {
	return a.loreService.GetByCharacterID(id)
}

func (a *App) FetchGeneralLore() ([]models.Lore, error) {
	return a.loreService.GetGeneralLore()
}

func (a *App) DeleteLore(id int64) error {
	return a.loreService.Delete(id)
}

// Lore Category

func (a *App) SaveLoreCategory(id int64, label string) (*models.LoreCategory, error) {
	if id == 0 {
		return a.loreService.CreateCategory(label)
	}
	return a.loreService.UpdateCategory(id, label)
}

func (a *App) FetchCategories() ([]models.LoreCategory, error) {
	return a.loreService.GetAllCategories()
}

func (a *App) DeleteLoreCategory(id int64) error {
	return a.loreService.DeleteCategory(id)
}

// Relationship

func (a *App) SaveRelationships(relationship models.Relationship) (*models.Relationship, error) {
	if relationship.ID == 0 {
		return a.relationshipService.Create(relationship)
	}
	return a.relationshipService.Update(relationship)
}

func (a *App) FetchRelationship(id int64) (*models.Relationship, error) {
	return a.relationshipService.GetByID(id)
}

func (a *App) FetchCharacterRelationships(id int64) ([]models.Relationship, error) {
	return a.relationshipService.GetByCharacterID(id)
}

func (a *App) DeleteRelationship(id int64) error {
	return a.relationshipService.Delete(id)
}

// Traits

func (a *App) CreateTrait(category, label string) (*services.TraitItem, error) {
	return a.traitService.Create(category, label)
}

func (a *App) SaveTrait(id int64, label, category string) (*services.TraitItem, error) {
	return a.traitService.Update(id, label, category)
}

func (a *App) FetchTrait(category string, id int64) (*services.TraitItem, error) {
	return a.traitService.Fetch(category, id)
}

func (a *App) FetchAllTraits(category string) ([]services.TraitItem, error) {
	return a.traitService.FetchAll(category)
}

func (a *App) DeleteTrait(category string, id int64) error {
	return a.traitService.Delete(category, id)
}
