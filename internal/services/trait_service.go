package services

import (
	"fmt"
	"reflect"
	"world-builder/internal/logger"
	"world-builder/internal/models"

	"gorm.io/gorm"
)

type TraitService struct {
	db *gorm.DB
}

func NewTraitService(db *gorm.DB) *TraitService {
	return &TraitService{db: db}
}

type TraitItem struct {
	ID    int64  `json:"id"`
	Label string `json:"label"`
}

func (s *TraitService) getModel(category string) (interface{}, error) {
	switch category {
	case "personality":
		return &models.PersonalityTrait{}, nil
	case "strength":
		return &models.Strength{}, nil
	case "flaw":
		return &models.Flaw{}, nil
	case "weakness":
		return &models.Weakness{}, nil
	case "fear":
		return &models.Fear{}, nil
	default:
		logger.LogError("TRAIT MODEL", fmt.Errorf("Invalid category %s", category))
		return nil, fmt.Errorf("Invalid trait category: %s", category)
	}
}

func (s *TraitService) Create(category, label string) (*TraitItem, error) {
	model, err := s.getModel(category)
	if err != nil {
		return nil, err
	}

	val := reflect.ValueOf(model).Elem()
	val.FieldByName("Label").SetString(label)

	if err = s.db.Create(model).Error; err != nil {
		logger.LogError("TRAIT DB:Save", err)
		return nil, err
	}

	id := val.FieldByName("ID").Int()

	return &TraitItem{
		ID:    id,
		Label: label,
	}, err
}

func (s *TraitService) Fetch(category string, id int64) (*TraitItem, error) {
	model, err := s.getModel(category)
	if err != nil {
		return nil, err
	}

	var item TraitItem
	if err = s.db.Model(model).First(&item, id).Error; err != nil {
		logger.LogError("TRAIT DB:Get", err)
		return nil, err
	}

	return &item, err
}

func (s *TraitService) FetchAll(category string) ([]TraitItem, error) {
	model, err := s.getModel(category)
	if err != nil {
		return nil, err
	}

	var results []TraitItem
	err = s.db.Model(model).Order("label ASC").Find(&results).Error

	if err != nil {
		logger.LogError("TRAIT DB:Get", err)
		return nil, err
	}

	return results, nil
}

func (s *TraitService) Update(id int64, label, category string) (*TraitItem, error) {
	model, err := s.getModel(category)
	if err != nil {
		return nil, err
	}

	result := s.db.Model(model).Where("id = ?", id).Update("label", label)
	if result.Error != nil {
		logger.LogError("TRAIT DB:Update", result.Error)
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		logger.LogError("TRAIT DB:Update", fmt.Errorf("No rows affected for id %d in category %s", id, category))
		return nil, fmt.Errorf("No rows affected for id %d in category %s", id, category)
	}

	return s.Fetch(category, id)

}

func (s *TraitService) Delete(category string, id int64) error {
	model, err := s.getModel(category)
	if err != nil {
		return err
	}

	err = s.db.Model(model).Where("id = ?", id).Delete(model).Error
	if err != nil {
		logger.LogError("TRAIT DB:Delete", err)
		return err
	}
	return nil
}
