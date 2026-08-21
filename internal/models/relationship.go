package models

type Relationship struct {
	ID           int64  `gorm:"primaryKey;autoIncrement" json:"id"`
	SourceCharID int64  `gorm:"not null;uniqueIndex:idx_source_target" json:"source_char_id"`
	TargetCharID int64  `gorm:"not null;uniqueIndex:idx_source_target" json:"target_char_id"`
	Type         string `gorm:"size:50;not null;default:'Acquaintance'" json:"type"` // "Companion", "Rival", "Parent", "Lover", "Enemy"
	Notes        string `gorm:"type:text" json:"notes,omitempty"`                    // Optional context/details

	SourceCharacter *Character `gorm:"foreignKey:SourceCharID;references:ID" json:"source_character,omitempty"`
	TargetCharacter *Character `gorm:"foreignKey:TargetCharID;references:ID" json:"target_character,omitempty"`
}
