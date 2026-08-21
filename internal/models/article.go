package models

import (
	"errors"
	"strings"
	"time"

	"gorm.io/gorm"
)

type Article struct {
	ID          int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	Title       string    `gorm:"size:150;not null" json:"title"`
	Description string    `gorm:"type:text;not null;default:''" json:"description"`
	Category    string    `gorm:"size:50;not null;default:'General'" json:"category"`
	Tags        string    `gorm:"type:text" json:"tags,omitempty"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func (a *Article) Validate() error {
	if strings.TrimSpace(a.Title) == "" {
		return errors.New("Article title cannot be empty.")
	}
	if len(a.Title) > 150 {
		return errors.New("Article title cannot exceed 150 characters.")
	}
	if strings.TrimSpace(a.Category) == "" {
		a.Category = "General"
	}

	return nil
}

func (a *Article) BeforeSave(tx *gorm.DB) error {
	a.Title = strings.TrimSpace(a.Title)
	a.Description = strings.TrimSpace(a.Description)
	a.UpdatedAt = time.Now()
	return nil
}
