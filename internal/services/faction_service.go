package services

import (
	"errors"
	"fmt"
	"world-builder/internal/models"

	"gorm.io/gorm"
)

type FactionService struct {
	db *gorm.DB
}

func NewFactionService(db *gorm.DB) *FactionService {
	return &FactionService{db: db}
}

func (s *FactionService) CreateFaction(faction models.Faction) (*models.Faction, error) {
	if err := s.db.Create(&faction).Error; err != nil {
		return nil, err
	}

	return s.GetFaction(faction.ID)

	// return &faction, nil
}

func (s *FactionService) GetFaction(id int64) (*models.Faction, error) {
	var faction models.Faction

	if err := s.db.First(&faction, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("Faction #%d not found.", id)
		}
		return nil, err
	}

	return &faction, nil
}

func (s *FactionService) GetAllFactions() ([]models.Faction, error) {
	var factions []models.Faction

	if err := s.db.Order("name ASC").Find(&factions).Error; err != nil {
		return nil, err
	}

	return factions, nil
}

func (s *FactionService) GetFactionsWithAffiliations() ([]models.Faction, error) {
	var factions []models.Faction

	if err := s.db.Where("id IN (?)",
		s.db.Model(&models.Affiliation{}).Select("faction_id").Where("faction_id IS NOT NULL"),
	).Order("name ASC").Find(&factions).Error; err != nil {
		return nil, err
	}

	return factions, nil
}

func (s *FactionService) UpdateFaction(faction models.Faction) (*models.Faction, error) {
	if faction.ID == 0 {
		return nil, fmt.Errorf("Cannot update faction without a valid ID.")
	}

	result := s.db.Save(&faction)
	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("Faction #%d not found:", faction.ID)
	}

	return s.GetFaction(faction.ID)

	// return &faction, nil
}

func (s *FactionService) DeleteFaction(id int64) error {
	result := s.db.Delete(&models.Faction{}, id)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("Unable to delete. Faction #%d not found:", id)
	}

	return nil
}

// Affialiations

func (s *FactionService) CreateAffiliation(affiliation models.Affiliation) (*models.Affiliation, error) {
	if err := s.db.Create(&affiliation).Error; err != nil {
		return nil, err
	}

	return s.GetAffiliation(affiliation.ID)

	// return &affiliation, nil
}

func (s *FactionService) GetAffiliation(id int64) (*models.Affiliation, error) {
	var affiliation models.Affiliation

	if err := s.db.Preload("Faction").First(&affiliation, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("Affiliation #%d not found", id)
		}
		return nil, err
	}

	return &affiliation, nil
}

func (s *FactionService) GetFactionAffiliations(id int64) ([]models.Affiliation, error) {
	var affiliations []models.Affiliation

	if err := s.db.Preload("Faction").Where("faction_id = ?", id).Find(&affiliations).Error; err != nil {
		return nil, err
	}

	return affiliations, nil
}

func (s *FactionService) GetAllAffiliations() ([]models.Affiliation, error) {
	var affiliations []models.Affiliation

	if err := s.db.Preload("Faction").Order("name ASC").Find(&affiliations).Error; err != nil {
		return nil, err
	}

	return affiliations, nil
}

func (s *FactionService) UpdateAffiliation(affiliation models.Affiliation) (*models.Affiliation, error) {
	if affiliation.ID == 0 {
		return nil, fmt.Errorf("Cannot update affiliation without a valid ID.")
	}

	result := s.db.Save(&affiliation)
	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("Affiliation #%d not found.", affiliation.ID)
	}

	return s.GetAffiliation(affiliation.ID)

	// return &affiliation, nil
}

func (s *FactionService) DeleteAffiliation(id int64) error {
	result := s.db.Delete(&models.Affiliation{}, id)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("Unable to delete affiliation. Affiliation #%d not found.", id)
	}

	return nil
}
