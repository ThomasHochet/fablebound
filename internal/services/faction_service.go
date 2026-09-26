package services

import (
	"errors"
	"fmt"
	"world-builder/internal/logger"
	"world-builder/internal/models"

	"gorm.io/gorm"
)

type FactionService struct {
	db *gorm.DB
}

type FactionDetails struct {
	models.Faction
	Affiliations []models.Affiliation `json:"affiliations"`
}

func NewFactionService(db *gorm.DB) *FactionService {
	return &FactionService{db: db}
}

func (s *FactionService) CreateFaction(faction models.Faction) (*models.Faction, error) {
	if err := s.db.Create(&faction).Error; err != nil {
		logger.LogError("FACTION DB:Save", err)
		return nil, err
	}

	return s.GetFaction(faction.ID)

	// return &faction, nil
}

func (s *FactionService) GetFaction(id int64) (*models.Faction, error) {
	var faction models.Faction

	if err := s.db.First(&faction, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logger.LogError("FACTION DB:Get", gorm.ErrRecordNotFound)
			return nil, fmt.Errorf("Faction #%d not found.", id)
		}
		logger.LogError("FACTION DB:Get", err)
		return nil, err
	}

	return &faction, nil
}

func (s *FactionService) GetFactionAffiliations(id int64) (*FactionDetails, error) {
	var faction models.Faction

	if err := s.db.First(&faction, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logger.LogError("FACTION DB:Get", gorm.ErrRecordNotFound)
			return nil, fmt.Errorf("Faction #%d not found.", id)
		}
		logger.LogError("FACTION DB:GetFactionAffiliations (Faction)", err)
		return nil, err
	}

	var affiliations []models.Affiliation
	if err := s.db.Where("faction_id = ?", id).Find(&affiliations).Error; err != nil {
		logger.LogError("FACTION DB:GetFactionAffiliations (Affiliations)", err)
		return nil, err
	}

	return &FactionDetails{
		Faction:      faction,
		Affiliations: affiliations,
	}, nil
}

func (s *FactionService) GetAllFactions() ([]models.Faction, error) {
	var factions []models.Faction

	if err := s.db.Order("name ASC").Find(&factions).Error; err != nil {
		logger.LogError("FACTION DB:Get", err)
		return nil, err
	}

	return factions, nil
}

func (s *FactionService) GetFactionsWithAffiliations() ([]models.Faction, error) {
	var factions []models.Faction

	if err := s.db.Where("id IN (?)",
		s.db.Model(&models.Affiliation{}).Select("faction_id").Where("faction_id IS NOT NULL"),
	).Order("name ASC").Find(&factions).Error; err != nil {
		logger.LogError("FACTION DB:Get", err)
		return nil, err
	}

	return factions, nil
}

func (s *FactionService) UpdateFaction(faction models.Faction) (*models.Faction, error) {
	if faction.ID == 0 {
		logger.LogError("FACTION DB:Update", fmt.Errorf("Cannot update faction without a valid ID: %d", faction.ID))
		return nil, fmt.Errorf("Cannot update faction without a valid ID.")
	}

	result := s.db.Save(&faction)
	if result.Error != nil {
		logger.LogError("FACTION DB:Update", result.Error)
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		logger.LogError("FACTION DB:Update", fmt.Errorf("ID Not Found: %d", faction.ID))
		return nil, fmt.Errorf("Faction #%d not found:", faction.ID)
	}

	return s.GetFaction(faction.ID)

	// return &faction, nil
}

func (s *FactionService) DeleteFaction(id int64) error {
	result := s.db.Delete(&models.Faction{}, id)
	if result.Error != nil {
		logger.LogError("FACTION DB:Delete", result.Error)
		return result.Error
	}

	if result.RowsAffected == 0 {
		logger.LogError("FACTION DB:Delete", fmt.Errorf("ID Not Found: %d", id))
		return fmt.Errorf("Unable to delete. Faction #%d not found:", id)
	}

	return nil
}

// Affialiations

func (s *FactionService) CreateAffiliation(affiliation models.Affiliation) (*models.Affiliation, error) {
	if err := s.db.Create(&affiliation).Error; err != nil {
		logger.LogError("AFFILIATION DB:Save", err)
		return nil, err
	}

	return s.GetAffiliation(affiliation.ID)

	// return &affiliation, nil
}

func (s *FactionService) GetAffiliation(id int64) (*models.Affiliation, error) {
	var affiliation models.Affiliation

	if err := s.db.Preload("Faction").First(&affiliation, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logger.LogError("AFFILIATION DB:Get", gorm.ErrRecordNotFound)
			return nil, fmt.Errorf("Affiliation #%d not found", id)
		}
		logger.LogError("AFFILIATION DB:Get", err)
		return nil, err
	}

	return &affiliation, nil
}

func (s *FactionService) GetAffiliationsFaction(id int64) ([]models.Affiliation, error) {
	var affiliations []models.Affiliation

	if err := s.db.Preload("Faction").Where("faction_id = ?", id).Find(&affiliations).Error; err != nil {
		logger.LogError("AFFILIATION DB:Get", err)
		return nil, err
	}

	return affiliations, nil
}

func (s *FactionService) GetAllAffiliations() ([]models.Affiliation, error) {
	var affiliations []models.Affiliation

	if err := s.db.Preload("Faction").Order("name ASC").Find(&affiliations).Error; err != nil {
		logger.LogError("AFFILIATION DB:Get", err)
		return nil, err
	}

	return affiliations, nil
}

func (s *FactionService) UpdateAffiliation(affiliation models.Affiliation) (*models.Affiliation, error) {
	if affiliation.ID == 0 {
		logger.LogError("AFFILIATION DB:Update", fmt.Errorf("Cannot update affiliation without a valid ID: %d", affiliation.ID))
		return nil, fmt.Errorf("Cannot update affiliation without a valid ID.")
	}

	result := s.db.Save(&affiliation)
	if result.Error != nil {
		logger.LogError("AFFILIATION DB:Update", result.Error)
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		logger.LogError("AFFILIATION DB:Update", fmt.Errorf("ID Not Found: %d", affiliation.ID))
		return nil, fmt.Errorf("Affiliation #%d not found.", affiliation.ID)
	}

	return s.GetAffiliation(affiliation.ID)

	// return &affiliation, nil
}

func (s *FactionService) DeleteAffiliation(id int64) error {
	result := s.db.Delete(&models.Affiliation{}, id)
	if result.Error != nil {
		logger.LogError("AFFILIATION DB:Delete", result.Error)
		return result.Error
	}

	if result.RowsAffected == 0 {
		logger.LogError("AFFILIATION DB:Delete", fmt.Errorf("ID Not Found: %d", id))
		return fmt.Errorf("Unable to delete affiliation. Affiliation #%d not found.", id)
	}

	return nil
}
