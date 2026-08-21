package models

import "time"

type Lore struct {
	ID          int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	Title       string    `gorm:"size:150;not null" json:"title"`
	Content     string    `gorm:"type:text;not null;default:'The start of a new story'" json:"content"`
	CategoryID  int64     `gorm:"index;not null" json:"category_id"`   // "Historical Event", "Myth", "Backstory"
	CharacterID *int64    `gorm:"index" json:"character_id,omitempty"` // Null = General Lore | Non-Null = Character Backstory
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`

	LoreCategory *LoreCategory `gorm:"foreignKey:CategoryID;references:ID" json:"lore_category,omitempty"`
	Character    *Character    `gorm:"foreignKey:CharacterID;references:ID" json:"character,omitempty"`
}

type LoreCategory struct {
	ID    int64  `gorm:"primaryKey;autoIncrement" json:"id"`
	Label string `gorm:"size:100;not null;unique" json:"label"`
}
