package models

type Location struct {
	ID          int64  `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string `gorm:"size:150;not null" json:"name"`
	Type        string `gorm:"size:150;not null" json:"type"`
	Subtype     string `gorm:"size:150" json:"subtype,omitempty"`
	Description string `gorm:"type:text" json:"description,omitempty"`

	ParentID *int64     `gorm:"index" json:"parent_id,omitempty"`
	Parent   *Location  `gorm:"foreignKey:ParentID;references:ID" json:"parent,omitempty"`
	Children []Location `gorm:"foreignKey:ParentID;references:ID" json:"children,omitempty"`
}
