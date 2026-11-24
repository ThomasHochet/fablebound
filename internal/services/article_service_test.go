package services_test

import (
	"errors"
	"testing"
	"world-builder/internal/db"
	"world-builder/internal/models"
	"world-builder/internal/services"

	"gorm.io/gorm"
)

func TestCreate(t *testing.T) {
	database := db.SetupTestDB(t)
	articleService := services.NewArticleService(database)

	newArticle := models.Article{
		Title:       "The Fall of Gondolin",
		Description: "Turgon built the hidden city...",
	}

	created, err := articleService.Create(newArticle)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if created == nil {
		t.Fatalf("Expected a returned article, got nil")
	}
	if created.Title != "The Fall of Gondolin" {
		t.Errorf("Expected title 'The Fall of Gondolin', got '%s'", created.Title)
	}

	var count int64
	database.Model(&models.Article{}).Where("title = ?", "The Fall of Gondolin").Count(&count)
	if count != 1 {
		t.Errorf("Expected 1 record in database, but found %d", count)
	}
}

func TestCreate_ValidationError(t *testing.T) {
	database := db.SetupTestDB(t)
	articleService := services.NewArticleService(database)

	newArticle := models.Article{
		Title:       "",
		Description: "Missing Title Content.",
	}

	created, err := articleService.Create(newArticle)

	if err == nil {
		t.Fatal("Expected a validation error, got nil")
	}
	if created != nil {
		t.Errorf("Expected nil article on validation failure, got %+v", created)
	}

	var count int64
	database.Model(&models.Article{}).Count(&count)
	if count != 0 {
		t.Errorf("Expected 0 records in database due to validation failure, got %d", count)
	}
}

func TestUpdate_Success(t *testing.T) {
	database := db.SetupTestDB(t)
	articleService := services.NewArticleService(database)

	newArticle := models.Article{
		Title:       "The Fall of Gondolin",
		Description: "Turgon built the hidden city...",
	}

	created, err := articleService.Create(newArticle)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	created.Title = "Updated Title"
	created.Description = "Updated Description."

	updated, err := articleService.Update(*created)
	if err != nil {
		t.Fatalf("Expected no error on update, got %v", err)
	}

	if updated.Title != "Updated Title" {
		t.Errorf("Expected title 'Updated Title', got %s", updated.Title)
	}

	var fromDB models.Article
	if err := database.First(&fromDB, created.ID).Error; err != nil {
		t.Fatalf("Failed to fetch updated article from DB: %v", err)
	}
	if fromDB.Description != "Updated Description." {
		t.Errorf("expected description 'Updated Description.' in DB, got '%s'", fromDB.Description)
	}
}

func TestUpdate_ValidationError(t *testing.T) {
	database := db.SetupTestDB(t)
	articleService := services.NewArticleService(database)

	newArticle := models.Article{
		Title:       "The Fall of Gondolin",
		Description: "Turgon built the hidden city...",
	}

	created, err := articleService.Create(newArticle)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	created.Title = ""

	updated, err := articleService.Update(*created)
	if err == nil {
		t.Fatalf("Expected a validation Error, got nil")
	}

	if updated != nil {
		t.Errorf("Expected nil article update on validation, got %+v", updated)
	}

	var fromDB models.Article
	if err := database.First(&fromDB, created.ID).Error; err != nil {
		t.Fatalf("Failed to fetch article from DB: %v", err)
	}
	if fromDB.Title != "The Fall of Gondolin" {
		t.Errorf("Expected DB title to remain 'The Fall of Gondolin', got %s", fromDB.Title)
	}
}

func TestDelete_Success(t *testing.T) {
	database := db.SetupTestDB(t)
	articleService := services.NewArticleService(database)

	newArticle := models.Article{
		Title:       "The Fall of Gondolin",
		Description: "Turgon built the hidden city...",
	}

	created, err := articleService.Create(newArticle)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if err := articleService.Delete(created.ID); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	var fromDB models.Article
	err = database.First(&fromDB, created.ID).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		t.Errorf("Expected gorm.ErrRecordNotFound after delete, got %v", err)
	}
}

func TestDelete_NotFound(t *testing.T) {
	database := db.SetupTestDB(t)
	articleService := services.NewArticleService(database)

	err := articleService.Delete(999)

	// Take 1 : J'suis debile.
	// if err != nil {
	// 	t.Fatalf("Expected no error when deleting non-existent ID, got %v", err)
	// }

	if err == nil {
		t.Fatal("Expected an error when deleting non-existent ID, got nil")
	}
}

func TestFetch_Success(t *testing.T) {
	database := db.SetupTestDB(t)
	articleService := services.NewArticleService(database)

	newArticle := models.Article{
		Title:       "The Fall of Gondolin",
		Description: "Turgon built the hidden city...",
	}

	created, err := articleService.Create(newArticle)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	fetched, err := articleService.GetArticle(created.ID)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if fetched.ID != created.ID {
		t.Errorf("Expected ID %d, got %d", created.ID, fetched.ID)
	}
	if fetched.Title != created.Title {
		t.Errorf("Expected title 'The Fall of Gondolin', got '%s'", fetched.Title)
	}
	if fetched.Description != created.Description {
		t.Errorf("Expected description 'Turgon built the hidden city...', got '%s'", fetched.Description)
	}
}

func TestFetch_NotFound(t *testing.T) {
	database := db.SetupTestDB(t)
	articleService := services.NewArticleService(database)

	fetched, err := articleService.GetArticle(999)

	if err == nil {
		t.Fatal("Expected an error when fetching non-existent article, got nil")
	}
	if fetched != nil {
		t.Errorf("Expected nil article when not found, got %+v", fetched)
	}
}

func TestGetAll_Empty(t *testing.T) {
	database := db.SetupTestDB(t)
	articleService := services.NewArticleService(database)

	articles, err := articleService.GetAll()
	if err != nil {
		t.Fatalf("Expected no error on empty DB, got %v", err)
	}
	if len(articles) != 0 {
		t.Errorf("Expected 0 articles, got %d", len(articles))
	}
}

func TestGetAll_Success(t *testing.T) {
	database := db.SetupTestDB(t)
	articleService := services.NewArticleService(database)

	_, err := articleService.Create(models.Article{
		Title:       "The Fall of Gondolin",
		Description: "Turgon built the hidden city...",
	})
	if err != nil {
		t.Fatalf("Failed to seed first article: %v", err)
	}

	_, err = articleService.Create(models.Article{
		Title:       "Oath of Fëanor",
		Description: "A fateful vow among the Noldor...",
	})
	if err != nil {
		t.Fatalf("Failed to seed second article: %v", err)
	}

	articles, err := articleService.GetAll()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if len(articles) != 2 {
		t.Fatalf("Expected 2 articles, got %d", len(articles))
	}

	titles := make(map[string]bool, len(articles))
	for _, a := range articles {
		titles[a.Title] = true
	}
	if !titles["The Fall of Gondolin"] || !titles["Oath of Fëanor"] {
		t.Errorf("Expected both seeded titles in result, got %+v", articles)
	}
}
