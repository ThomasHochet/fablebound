package services

import (
	"errors"
	"fmt"
	"world-builder/internal/models"

	"gorm.io/gorm"
)

type ArticleService struct {
	db *gorm.DB
}

func NewArticleService(db *gorm.DB) *ArticleService {
	return &ArticleService{db: db}
}

func (s *ArticleService) Create(title, description, category string) (*models.Article, error) {
	article := &models.Article{
		Title:       title,
		Description: description,
		Category:    category,
	}

	if err := article.Validate(); err != nil {
		return nil, err
	}

	if err := s.db.Create(article).Error; err != nil {
		return nil, fmt.Errorf("Failed to create article: %w", err)
	}

	return article, nil
}

func (s *ArticleService) GetArticle(id int64) (*models.Article, error) {
	var article models.Article

	if err := s.db.First(&article, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("Article #%d not found.", id)
		}
		return nil, err
	}

	return &article, nil
}

func (s *ArticleService) GetAll() ([]models.Article, error) {
	var articles []models.Article

	if err := s.db.Order("updated_at desc").Find(&articles).Error; err != nil {
		return nil, err
	}

	return articles, nil
}

// Get all already stored categories
func (s *ArticleService) GetCategories() ([]string, error) {
	var categories []string

	err := s.db.Table("articles").
		Select("DISTINCT category").
		Where("category IS NOT NULL AND category != ''").
		Order("category ASC").
		Find(&categories).Error

	return categories, err
}

func (s *ArticleService) Update(id int64, title, description string) (*models.Article, error) {
	article, err := s.GetArticle(id)
	if err != nil {
		return nil, err
	}

	article.Title = title
	article.Description = description

	if err := article.Validate(); err != nil {
		return nil, err
	}

	if err := s.db.Save(article).Error; err != nil {
		return nil, fmt.Errorf("Couldn't update article: %w", err)
	}

	return article, nil
}

func (s *ArticleService) Delete(id int64) error {
	result := s.db.Delete(&models.Article{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("Article #%d not found to delete", id)
	}

	return nil
}
