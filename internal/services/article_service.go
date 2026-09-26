package services

import (
	"errors"
	"fmt"
	"world-builder/internal/logger"
	"world-builder/internal/models"

	"gorm.io/gorm"
)

type ArticleService struct {
	db *gorm.DB
}

func NewArticleService(db *gorm.DB) *ArticleService {
	return &ArticleService{db: db}
}

func (s *ArticleService) Create(article models.Article) (*models.Article, error) {

	if err := article.Validate(); err != nil {
		logger.LogError("Article:Validate", err)
		return nil, err
	}

	if err := s.db.Create(&article).Error; err != nil {
		logger.LogError("ARTICLE DB:Save", err)
		return nil, fmt.Errorf("Failed to create article: %w", err)
	}

	return &article, nil
}

func (s *ArticleService) GetArticle(id int64) (*models.Article, error) {
	var article models.Article

	if err := s.db.First(&article, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logger.LogError("ARTICLE DB:Get", err)
			return nil, fmt.Errorf("Article #%d not found.", id)
		}
		return nil, err
	}

	return &article, nil
}

func (s *ArticleService) GetAll() ([]models.Article, error) {
	var articles []models.Article

	if err := s.db.Order("updated_at desc").Find(&articles).Error; err != nil {
		logger.LogError("ARTICLE DB:Get", err)
		return nil, err
	}

	return articles, nil
}

// Get all already stored categories
func (s *ArticleService) GetCategories() ([]string, error) {
	var categories []string

	if err := s.db.Table("articles").
		Select("DISTINCT category").
		Where("category IS NOT NULL AND category != ''").
		Order("category ASC").
		Find(&categories).Error; err != nil {
		logger.LogError("ARTICLE CATEGORIES DB:Get", err)
		return nil, err
	}

	return categories, nil
}

func (s *ArticleService) Update(article models.Article) (*models.Article, error) {
	if article.ID == 0 {
		logger.LogError("ARTICLE DB:Update", fmt.Errorf("Invalid ID"))
		return nil, fmt.Errorf("Cannot update character with an invalid ID")
	}

	// if err := article.Validate(); err != nil {
	// 	return nil, err
	// }

	if err := s.db.Save(&article).Error; err != nil {
		logger.LogError("ARTICLE DB:Update", err)
		return nil, fmt.Errorf("Couldn't update article: %w", err)
	}

	return &article, nil
}

func (s *ArticleService) Delete(id int64) error {
	result := s.db.Delete(&models.Article{}, id)
	if result.Error != nil {
		logger.LogError("ARTICLE DB:Delete", result.Error)
		return result.Error
	}
	if result.RowsAffected == 0 {
		logger.LogError("ARTICLE DB:Delete", fmt.Errorf("#%d not found", id))
		return fmt.Errorf("Article #%d not found to delete", id)
	}

	return nil
}
