package services

import (
	"fmt"
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
		return nil, fmt.Errorf("Invalid trait category: %s", category)
	}
}

func (s *TraitService) Create(category, label string) (*TraitItem, error) {
	model, err := s.getModel(category)
	if err != nil {
		return nil, err
	}

	data := map[string]interface{}{"label": label}
	if err := s.db.Model(model).Create(&data).Error; err != nil {
		return nil, err
	}

	return s.Fetch(category, data["id"].(int64))

	// return &TraitItem{
	// 	ID:    data["id"].(int64),
	// 	Label: label,
	// }, err
}

func (s *TraitService) Fetch(category string, id int64) (*TraitItem, error) {
	model, err := s.getModel(category)
	if err != nil {
		return nil, err
	}

	var item TraitItem
	if err = s.db.Model(model).First(&item, id).Error; err != nil {
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

	return results, err
}

func (s *TraitService) Update(id int64, label, category string) (*TraitItem, error) {
	model, err := s.getModel(category)
	if err != nil {
		return nil, err
	}

	result := s.db.Model(model).Where("id = ?", id).Update("label", label)
	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("No rows affected for id %d in category %s", id, category)
	}

	return s.Fetch(category, id)

}

func (s *TraitService) Delete(category string, id int64) error {
	model, err := s.getModel(category)
	if err != nil {
		return err
	}
	return s.db.Model(model).Where("id = ?", id).Delete(model).Error
}
