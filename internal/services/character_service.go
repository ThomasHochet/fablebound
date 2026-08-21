package services

import (
	"errors"
	"fmt"
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
		Preload("Relationships.TargetCharacter").
		Preload("Backstories.LoreCategory")
}

func (s *CharacterService) Create(char models.Character) (*models.Character, error) {
	if err := s.db.Create(&char).Error; err != nil {
		return nil, err
	}

	return s.GetByID(char.ID)
}

func (s *CharacterService) GetByID(id int64) (*models.Character, error) {
	var character models.Character

	err := s.preloadCharacter(s.db).First(&character, id).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("Character #%d not found.", id)
		}
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
		return nil, err
	}

	return characters, nil
}

func (s *CharacterService) Update(char models.Character) (*models.Character, error) {
	if char.ID == 0 {
		return nil, fmt.Errorf("Cannot update character with an invalid id")
	}

	result := s.db.Save(&char)
	if result.Error != nil {
		return nil, result.Error
	}

	s.db.Model(&char).Association("PersonalityTraits").Replace(char.PersonalityTraits)
	s.db.Model(&char).Association("Strengths").Replace(char.Strengths)
	s.db.Model(&char).Association("Flaws").Replace(char.Flaws)
	s.db.Model(&char).Association("Weaknesses").Replace(char.Weaknesses)
	s.db.Model(&char).Association("Fears").Replace(char.Fears)

	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("Character #%d not found.", char.ID)
	}

	return s.GetByID(char.ID)
}

func (s *CharacterService) Delete(id int64) error {
	if id == 0 {
		return fmt.Errorf("Invalid ID: #%d.", id)
	}
	result := s.db.Delete(&models.Character{}, id)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("Unable to delete character. Character #%d not found.", id)
	}

	return nil
}
