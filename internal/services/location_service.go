package services

import (
	"errors"
	"fmt"
	"world-builder/internal/logger"
	"world-builder/internal/models"

	"gorm.io/gorm"
)

type LocationService struct {
	db *gorm.DB
}

func NewLocationService(db *gorm.DB) *LocationService {
	return &LocationService{db: db}
}

// Create inserts a new location, ensuring it doesn't parent itself
func (s *LocationService) Create(location models.Location) (*models.Location, error) {
	if location.ParentID != nil && *location.ParentID == location.ID {
		logger.LogError("LOCATION DB:Save", fmt.Errorf("Location cannot be it's own parent."))
		return nil, fmt.Errorf("a location cannot be its own parent")
	}

	if err := s.db.Create(&location).Error; err != nil {
		logger.LogError("LOCATION DB:Save", err)
		return nil, err
	}

	return s.GetByID(location.ID)
}

// GetByID fetches a location, preloading its parent and its immediate children
func (s *LocationService) GetByID(id int64) (*models.Location, error) {
	var location models.Location

	err := s.db.
		Preload("Parent").
		Preload("Children").
		First(&location, id).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logger.LogError("LOCATION DB:Get", gorm.ErrRecordNotFound)
			return nil, fmt.Errorf("location #%d not found", id)
		}
		logger.LogError("LOCATION DB:Get", err)
		return nil, err
	}

	return &location, nil
}

// GetRoots fetches top-level locations (continents, planets, etc.) where ParentID is nil
func (s *LocationService) GetRoots() ([]models.Location, error) {
	var locations []models.Location

	err := s.db.
		Preload("Children").
		Where("parent_id IS NULL").
		Order("name ASC").
		Find(&locations).Error

	if err != nil {
		logger.LogError("LOCATION DB:Get", err)
		return nil, err
	}

	return locations, nil
}

// GetAll fetches all locations flatly, preloading only the parent link
func (s *LocationService) GetAll() ([]models.Location, error) {
	var locations []models.Location

	err := s.db.
		Preload("Parent").
		Order("name ASC").
		Find(&locations).Error

	if err != nil {
		logger.LogError("LOCATION DB:Get", err)
		return nil, err
	}

	return locations, nil
}

// Get to only get locations types, and subtypes
func (s *LocationService) GetTypes() ([]string, error) {
	var types []string

	err := s.db.Table("locations").
		Select("DISTINCT type").
		Where("type IS NOT NULL AND type != ''").
		Order("type ASC").
		Find(&types).Error

	if err != nil {
		logger.LogError("LOCATION DB:Get", err)
		return nil, err
	}

	return types, nil
}

func (s *LocationService) GetSubtypes() ([]string, error) {
	var subtypes []string

	err := s.db.Table("locations").
		Select("DISTINCT subtype").
		Where("subtype IS NOT NULL AND subtype != ''").
		Order("subtype ASC").
		Find(&subtypes).Error

	if err != nil {
		logger.LogError("LOCATION DB:Get", err)
		return nil, err
	}

	return subtypes, nil
}

// Update saves changes and prevents a location from becoming its own parent
func (s *LocationService) Update(location models.Location) (*models.Location, error) {
	if location.ID == 0 {
		logger.LogError("LOCATION DB:Update", fmt.Errorf("Invalid ID: %d", location.ID))
		return nil, fmt.Errorf("cannot update location without a valid ID")
	}

	if location.ParentID != nil && *location.ParentID == location.ID {
		logger.LogError("LOCATION DB:Update", fmt.Errorf("a location cannot be its own parent"))
		return nil, fmt.Errorf("a location cannot be its own parent")
	}

	result := s.db.Save(&location)
	if result.Error != nil {
		logger.LogError("LOCATION DB:Update", result.Error)
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		logger.LogError("LOCATION DB:Update", fmt.Errorf("ID Not Found: %d", location.ID))
		return nil, fmt.Errorf("location #%d not found", location.ID)
	}

	return s.GetByID(location.ID)
}

// Delete removes a location. Note: You may need to handle what happens to its children!
func (s *LocationService) Delete(id int64) error {
	result := s.db.Delete(&models.Location{}, id)
	if result.Error != nil {
		logger.LogError("LOCATION DB:Delete", result.Error)
		return result.Error
	}

	if result.RowsAffected == 0 {
		logger.LogError("LOCATION DB:Delete", fmt.Errorf("ID Not Found: %d", id))
		return fmt.Errorf("location #%d not found", id)
	}

	return nil
}
