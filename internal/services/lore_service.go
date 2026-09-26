package services

import (
	"errors"
	"fmt"
	"world-builder/internal/logger"
	"world-builder/internal/models"

	"gorm.io/gorm"
)

type LoreService struct {
	db *gorm.DB
}

func NewLoreService(db *gorm.DB) *LoreService {
	return &LoreService{db: db}
}

// --- LORE CRUD & QUERIES ---

// Create inserts a new lore entry
func (s *LoreService) Create(lore models.Lore) (*models.Lore, error) {
	if lore.CategoryID == 0 {
		logger.LogError("LORE DB:Save", fmt.Errorf("Invalid ID: %d", lore.CategoryID))
		return nil, fmt.Errorf("a valid category_id is required")
	}

	if err := s.db.Create(&lore).Error; err != nil {
		logger.LogError("LORE DB:Save", err)
		return nil, err
	}

	// Re-fetch to populate LoreCategory and optional Character structs
	return s.GetByID(lore.ID)
}

// GetByID fetches a single lore entry by ID with preloaded associations
func (s *LoreService) GetByID(id int64) (*models.Lore, error) {
	var lore models.Lore

	err := s.db.
		Preload("LoreCategory").
		Preload("Character").
		First(&lore, id).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logger.LogError("LORE DB:Get", gorm.ErrRecordNotFound)
			return nil, fmt.Errorf("lore entry #%d not found", id)
		}
		logger.LogError("LORE DB:Get", err)
		return nil, err
	}

	return &lore, nil
}

// GetAll fetches all lore entries ordered by newest first
func (s *LoreService) GetAll() ([]models.Lore, error) {
	var lores []models.Lore

	err := s.db.
		Preload("LoreCategory").
		Preload("Character").
		Order("created_at DESC").
		Find(&lores).Error

	if err != nil {
		logger.LogError("LORE DB:Get", err)
		return nil, err
	}

	return lores, nil
}

// GetByCharacterID fetches all backstory/lore entries assigned to a specific character
func (s *LoreService) GetByCharacterID(characterID int64) ([]models.Lore, error) {
	var lores []models.Lore

	err := s.db.
		Preload("LoreCategory").
		Preload("Character").
		Where("character_id = ?", characterID).
		Order("created_at DESC").
		Find(&lores).Error

	if err != nil {
		logger.LogError("LORE DB:Get", err)
		return nil, err
	}

	return lores, nil
}

// GetByCategoryID fetches all lore entries under a specific category (e.g., "Historical Event")
func (s *LoreService) GetByCategoryID(categoryID int64) ([]models.Lore, error) {
	var lores []models.Lore

	err := s.db.
		Preload("LoreCategory").
		Preload("Character").
		Where("category_id = ?", categoryID).
		Order("created_at DESC").
		Find(&lores).Error

	if err != nil {
		logger.LogError("LORE DB:Get", err)
		return nil, err
	}

	return lores, nil
}

// GetGeneralLore fetches world lore entries NOT tied to any specific character (CharacterID IS NULL)
func (s *LoreService) GetGeneralLore() ([]models.Lore, error) {
	var lores []models.Lore

	err := s.db.
		Preload("LoreCategory").
		Where("character_id IS NULL").
		Order("created_at DESC").
		Find(&lores).Error

	if err != nil {
		logger.LogError("LORE DB:Get", err)
		return nil, err
	}

	return lores, nil
}

// Update saves changes to a lore record
func (s *LoreService) Update(lore models.Lore) (*models.Lore, error) {
	if lore.ID == 0 {
		logger.LogError("LORE DB:Update", fmt.Errorf("Invalid ID: %d", lore.ID))
		return nil, fmt.Errorf("cannot update lore entry without a valid ID")
	}

	result := s.db.Save(&lore)
	if result.Error != nil {
		logger.LogError("LORE DB:Update", result.Error)
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		logger.LogError("LORE DB:Update", fmt.Errorf("lore entry #%d not found", lore.ID))
		return nil, fmt.Errorf("lore entry #%d not found", lore.ID)
	}

	return s.GetByID(lore.ID)
}

// Delete removes a lore record by ID
func (s *LoreService) Delete(id int64) error {
	result := s.db.Delete(&models.Lore{}, id)
	if result.Error != nil {
		logger.LogError("LORE DB:Delete", result.Error)
		return result.Error
	}

	if result.RowsAffected == 0 {
		logger.LogError("LORE DB:Delete", fmt.Errorf("lore entry #%d not found", id))
		return fmt.Errorf("lore entry #%d not found", id)
	}

	return nil
}

// --- LORE CATEGORIES CRUD ---

func (s *LoreService) CreateCategory(label string) (*models.LoreCategory, error) {
	cat := models.LoreCategory{Label: label}
	if err := s.db.Create(&cat).Error; err != nil {
		logger.LogError("LORE Cat DB:Save", err)
		return nil, err
	}
	return &cat, nil
}

func (s *LoreService) GetAllCategories() ([]models.LoreCategory, error) {
	var categories []models.LoreCategory
	err := s.db.Order("label ASC").Find(&categories).Error

	if err != nil {
		logger.LogError("LORE Cat DB:Get", err)
		return nil, err
	}
	return categories, nil
}

func (s *LoreService) GetBackstoryLoreCategory() (*models.LoreCategory, error) {
	var category models.LoreCategory

	err := s.db.Where("label = ?", "Backstory").First(&category).Error

	if err != nil {
		logger.LogError("LORE Cat DB:Get", err)
		return nil, err
	}

	return &category, nil
}

// UpdateCategory renames a lore category by ID
func (s *LoreService) UpdateCategory(id int64, label string) (*models.LoreCategory, error) {
	if id == 0 {
		logger.LogError("LORE Cat DB:Update", fmt.Errorf("Invalid ID: %d", id))
		return nil, fmt.Errorf("cannot update category without a valid ID")
	}

	if label == "" {
		logger.LogError("LORE Cat DB:Update", fmt.Errorf("Empty Label"))
		return nil, fmt.Errorf("category label cannot be empty")
	}

	category := models.LoreCategory{
		ID:    id,
		Label: label,
	}

	// Save updates the record matching category.ID
	result := s.db.Save(&category)
	if result.Error != nil {
		logger.LogError("LORE Cat DB:Update", result.Error)
		return nil, result.Error // Returns SQLite unique constraint error if label exists
	}

	if result.RowsAffected == 0 {
		logger.LogError("LORE Cat DB:Update", fmt.Errorf("ID Not Found: %d", id))
		return nil, fmt.Errorf("lore category #%d not found", id)
	}

	return &category, nil
}

func (s *LoreService) DeleteCategory(id int64) error {
	result := s.db.Delete(&models.LoreCategory{}, id)
	if result.Error != nil {
		logger.LogError("LORE Cat DB:Delete", result.Error)
		return result.Error
	}

	if result.RowsAffected == 0 {
		logger.LogError("LORE Cat DB:Delete", fmt.Errorf("ID Not Found: %d", id))
		return fmt.Errorf("lore category #%d not found", id)
	}

	return nil
}
