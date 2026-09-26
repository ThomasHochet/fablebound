package services

import (
	"errors"
	"fmt"
	"world-builder/internal/logger"
	"world-builder/internal/models"

	"gorm.io/gorm"
)

type RelationshipService struct {
	db *gorm.DB
}

func NewRelationshipService(db *gorm.DB) *RelationshipService {
	return &RelationshipService{db: db}
}

// Create inserts a new relationship link between two characters
func (s *RelationshipService) Create(rel models.Relationship) (*models.Relationship, error) {
	if rel.SourceCharID == rel.TargetCharID {
		logger.LogError("RELATION DB:Save", fmt.Errorf("Self relationship forbidden."))
		return nil, fmt.Errorf("a character cannot have a relationship with themselves")
	}

	if err := s.db.Create(&rel).Error; err != nil {
		logger.LogError("RELATION DB:Save", err)
		return nil, err
	}

	// Re-fetch to populate SourceCharacter and TargetCharacter structs for Svelte
	return s.GetByID(rel.ID)
}

// GetByID fetches a single relationship record with both characters populated
func (s *RelationshipService) GetByID(id int64) (*models.Relationship, error) {
	var rel models.Relationship

	err := s.db.
		Preload("SourceCharacter").
		Preload("TargetCharacter").
		First(&rel, id).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logger.LogError("RELATION DB:Get", gorm.ErrRecordNotFound)
			return nil, fmt.Errorf("relationship #%d not found", id)
		}
		logger.LogError("RELATION DB:Get", err)
		return nil, err
	}

	return &rel, nil
}

// GetByCharacterID fetches ALL relationships involving a character (as either source OR target)
func (s *RelationshipService) GetByCharacterID(characterID int64) ([]models.Relationship, error) {
	var relationships []models.Relationship

	err := s.db.
		Preload("SourceCharacter").
		Preload("TargetCharacter").
		Where("source_char_id = ? OR target_char_id = ?", characterID, characterID).
		Find(&relationships).Error

	if err != nil {
		logger.LogError("RELATION DB:Get", err)
		return nil, err
	}

	return relationships, nil
}

// Update saves changes to relationship type or notes
func (s *RelationshipService) Update(rel models.Relationship) (*models.Relationship, error) {
	if rel.ID == 0 {
		logger.LogError("RELATION DB:Update", fmt.Errorf("Invalid ID: %d", rel.ID))
		return nil, fmt.Errorf("cannot update relationship without a valid ID")
	}

	result := s.db.Save(&rel)
	if result.Error != nil {
		logger.LogError("RELATION DB:Update", result.Error)
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		logger.LogError("RELATION DB:Update", fmt.Errorf("ID Not Found: %d", rel.ID))
		return nil, fmt.Errorf("relationship #%d not found", rel.ID)
	}

	return s.GetByID(rel.ID)
}

// Delete removes a relationship link by ID
func (s *RelationshipService) Delete(id int64) error {
	result := s.db.Delete(&models.Relationship{}, id)
	if result.Error != nil {
		logger.LogError("RELATION DB:Delete", result.Error)
		return result.Error
	}

	if result.RowsAffected == 0 {
		logger.LogError("RELATION DB:Delete", fmt.Errorf("ID Not Found: %d", id))
		return fmt.Errorf("relationship #%d not found", id)
	}

	return nil
}
