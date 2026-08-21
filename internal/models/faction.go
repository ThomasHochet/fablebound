package models

type Affiliation struct {
	ID          int64  `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string `gorm:"size:150;not null" json:"name"`
	Description string `gorm:"type:text" json:"description,omitempty"`
	Type        string `gorm:"size:50" json:"type,omitempty"` // Free to call themselves however they want, no 'type' table necessary, for ease of storytelling
	FactionID   *int64 `gorm:"index" json:"faction_id,omitempty"`

	Faction *Faction `gorm:"foreignKey:FactionID;references:ID" json:"faction,omitempty"`
}

type Faction struct {
	ID          int64  `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string `gorm:"size:150;not null" json:"name"`
	Description string `gorm:"type:text" json:"description,omitempty"`
}
