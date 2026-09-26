package services

import (
	"errors"
	"fmt"
	"world-builder/internal/logger"
	"world-builder/internal/models"

	"gorm.io/gorm"
)

type CharacterService struct {
	db *gorm.DB
}

func NewCharacterService(db *gorm.DB) *CharacterService {
	return &CharacterService{db: db}
}

func (s *CharacterService) preloadCharacter(db *gorm.DB) *gorm.DB {
	return db.
		Preload("Gender").
		Preload("Race").
		Preload("Occupation").
		Preload("Alignment").
		Preload("Status").
		Preload("Affiliation").
		Preload("Faction").
		Preload("Birthplace").
		Preload("CurrentLocation").
		Preload("PersonalityTraits").
		Preload("Strengths").
		Preload("Flaws").
		Preload("Weaknesses").
		Preload("Fears").
		Preload("Relationships").
		Preload("Relationships.TargetCharacter").
		Preload("Backstories").
		Preload("Backstories.LoreCategory")
}

func (s *CharacterService) Create(char models.Character) (*models.Character, error) {
	if err := s.db.Create(&char).Error; err != nil {
		logger.LogError("CHARACTER DB:Create", err)
		return nil, err
	}

	s.db.Model(&char).Association("PersonalityTraits").Replace(char.PersonalityTraits)
	s.db.Model(&char).Association("Strengths").Replace(char.Strengths)
	s.db.Model(&char).Association("Flaws").Replace(char.Flaws)
	s.db.Model(&char).Association("Weaknesses").Replace(char.Weaknesses)
	s.db.Model(&char).Association("Fears").Replace(char.Fears)

	if err := s.db.Model(&char).Association("Relationships").Replace(char.Relationships); err != nil {
		logger.LogError("CHARACTER DB:Replace Relationships", err)
	}

	if err := s.db.Model(&char).Association("Backstories").Replace(char.Backstories); err != nil {
		logger.LogError("CHARACTER DB:Replace Backstories", err)
	}

	return s.GetByID(char.ID)
}

func (s *CharacterService) GetByID(id int64) (*models.Character, error) {
	var character models.Character

	err := s.preloadCharacter(s.db).First(&character, id).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logger.LogError("CHARACTER DB:Get", gorm.ErrRecordNotFound)
			return nil, fmt.Errorf("Character #%d not found.", id)
		}
		logger.LogError("CHARACTER DB:Get", err)
		return nil, err
	}

	return &character, nil
}

func (s *CharacterService) GetAllCharacters(limit int) ([]models.Character, error) {
	var characters []models.Character

	if limit <= 0 {
		limit = 200
	}

	if err := s.preloadCharacter(s.db).Limit(limit).Find(&characters).Error; err != nil {
		logger.LogError("CHARACTER DB:Get", err)
		return nil, err
	}

	return characters, nil
}

func (s *CharacterService) GetAllCharacterAsLookup() ([]LookupItem, error) {
	var characters []models.Character

	err := s.db.Model(&models.Character{}).
		Select("id", "firstname", "surname", "nickname").
		Order("firstname ASC, surname ASC").
		Find(&characters).Error

	if err != nil {
		return nil, err
	}

	lookups := make([]LookupItem, 0, len(characters))
	for _, c := range characters {
		displayName := c.Firstname + " " + c.Surname
		if c.Nickname != "" {
			displayName += " (" + c.Nickname + ")"
		}

		lookups = append(lookups, LookupItem{
			ID:    c.ID,
			Label: displayName,
		})
	}

	return lookups, nil
}

func (s *CharacterService) Update(char models.Character) (*models.Character, error) {
	if char.ID == 0 {
		logger.LogError("CHARACTER DB:Update", fmt.Errorf("Invalid ID: %d", char.ID))
		return nil, fmt.Errorf("Cannot update character with an invalid id")
	}

	result := s.db.Save(&char)
	if result.Error != nil {
		logger.LogError("CHARACTER DB:Update", result.Error)
		return nil, result.Error
	}

	s.db.Model(&char).Association("PersonalityTraits").Replace(char.PersonalityTraits)
	s.db.Model(&char).Association("Strengths").Replace(char.Strengths)
	s.db.Model(&char).Association("Flaws").Replace(char.Flaws)
	s.db.Model(&char).Association("Weaknesses").Replace(char.Weaknesses)
	s.db.Model(&char).Association("Fears").Replace(char.Fears)

	if result.RowsAffected == 0 {
		logger.LogError("CHARACTER DB:Update", fmt.Errorf("ID Not found: %d", char.ID))
		return nil, fmt.Errorf("Character #%d not found.", char.ID)
	}

	return s.GetByID(char.ID)
}

func (s *CharacterService) Delete(id int64) error {
	if id == 0 {
		logger.LogError("CHARACTER DB:Delete", fmt.Errorf("Invalid ID: %d", id))
		return fmt.Errorf("Invalid ID: #%d.", id)
	}
	result := s.db.Delete(&models.Character{}, id)
	if result.Error != nil {
		logger.LogError("CHARACTER DB:Delete", result.Error)
		return result.Error
	}

	if result.RowsAffected == 0 {
		logger.LogError("CHARACTER DB:Delete", fmt.Errorf("ID Not Found: %d", id))
		return fmt.Errorf("Unable to delete character. Character #%d not found.", id)
	}

	return nil
}
