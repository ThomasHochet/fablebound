package services

import (
	"fmt"
	"world-builder/internal/models"

	"gorm.io/gorm"
)

type LookupService struct {
	db *gorm.DB
}

func NewLookupService(db *gorm.DB) *LookupService {
	return &LookupService{db: db}
}

type LookupItem struct {
	ID    int64  `json:"id"`
	Label string `json:"label"`
}

func (s *LookupService) getModel(category string) (interface{}, error) {
	switch category {
	case "gender":
		return &models.Gender{}, nil
	case "race":
		return &models.Race{}, nil
	case "occupation":
		return &models.Occupation{}, nil
	case "alignment":
		return &models.Alignment{}, nil
	case "status":
		return &models.Status{}, nil
	default:
		return nil, fmt.Errorf("Invalid category: %s", category)
	}
}

func (s *LookupService) Create(category, label string) (*LookupItem, error) {
	model, err := s.getModel(category)
	if err != nil {
		return nil, err
	}

	data := map[string]interface{}{"label": label}
	if err = s.db.Model(model).Create(&data).Error; err != nil {
		return nil, err
	}

	return s.Fetch(data["id"].(int64), category)

	// return &LookupItem{
	// 	ID:    data["id"].(int64),
	// 	Label: label,
	// }, nil
}

func (s *LookupService) Fetch(id int64, category string) (*LookupItem, error) {
	model, err := s.getModel(category)
	if err != nil {
		return nil, err
	}

	var item LookupItem
	if err = s.db.Model(model).First(&item, id).Error; err != nil {
		return nil, err
	}

	return &item, err
}

func (s *LookupService) FetchAll(category string) ([]LookupItem, error) {
	model, err := s.getModel(category)
	if err != nil {
		return nil, err
	}

	var results []LookupItem
	if err = s.db.Model(model).Order("label ASC").Find(&results).Error; err != nil {
		return nil, err
	}

	return results, err
}

func (s *LookupService) Update(id int64, label, category string) (*LookupItem, error) {
	model, err := s.getModel(category)
	if err != nil {
		return nil, err
	}

	if err = s.db.Model(model).Where("id = ?", id).Update("label", label).Error; err != nil {
		return nil, err
	}

	return s.Fetch(id, category)
}

func (s *LookupService) Delete(id int64, category string) error {
	model, err := s.getModel(category)
	if err != nil {
		return err
	}

	return s.db.Model(model).Where("id = ?", id).Delete(model).Error

}
