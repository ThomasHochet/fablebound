package services

import (
	"fmt"
	"reflect"
	"world-builder/internal/logger"
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
		logger.LogError("LOOKUP Model", fmt.Errorf("Invalid category: %s", category))
		return nil, fmt.Errorf("Invalid category: %s", category)
	}
}

func (s *LookupService) Create(category, label string) (*LookupItem, error) {
	model, err := s.getModel(category)
	if err != nil {
		logger.LogError("LOOKUP DB:Save", err)
		return nil, err
	}

	val := reflect.ValueOf(model).Elem()
	val.FieldByName("Label").SetString(label)

	if err = s.db.Create(model).Error; err != nil {
		logger.LogError("LOOKUP DB:Save", err)
		return nil, err
	}

	id := val.FieldByName("ID").Int()

	// return s.Fetch(data["id"].(int64), category)

	return &LookupItem{
		ID:    id,
		Label: label,
	}, nil
}

func (s *LookupService) Fetch(id int64, category string) (*LookupItem, error) {
	model, err := s.getModel(category)
	if err != nil {
		logger.LogError("LOOKUP DB:Get", err)
		return nil, err
	}

	var item LookupItem
	if err = s.db.Model(model).First(&item, id).Error; err != nil {
		logger.LogError("LOOKUP DB:Get", err)
		return nil, err
	}

	return &item, err
}

func (s *LookupService) FetchAll(category string) ([]LookupItem, error) {
	model, err := s.getModel(category)
	if err != nil {
		logger.LogError("LOOKUP DB:Get", err)
		return nil, err
	}

	var results []LookupItem
	if err = s.db.Model(model).Order("label ASC").Find(&results).Error; err != nil {
		logger.LogError("LOOKUP DB:Get", err)
		return nil, err
	}

	return results, err
}

func (s *LookupService) Update(id int64, label, category string) (*LookupItem, error) {
	model, err := s.getModel(category)
	if err != nil {
		logger.LogError("LOOKUP DB:Update", err)
		return nil, err
	}

	if err = s.db.Model(model).Where("id = ?", id).Update("label", label).Error; err != nil {
		logger.LogError("LOOKUP DB:Update", err)
		return nil, err
	}

	return s.Fetch(id, category)
}

func (s *LookupService) Delete(id int64, category string) error {
	model, err := s.getModel(category)
	if err != nil {
		logger.LogError("LOOKUP DB:Delete", err)
		return err
	}

	err = s.db.Model(model).Where("id = ?", id).Delete(model).Error
	if err != nil {
		logger.LogError("LOOKUP DB:Delete", err)
		return err
	}
	return nil
}
