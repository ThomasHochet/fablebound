package models

type PersonalityTrait struct {
	ID    int64  `gorm:"primaryKey;autoIncrement" json:"id"`
	Label string `gorm:"size:100;not null;unique" json:"label"`
}

type Strength struct {
	ID    int64  `gorm:"primaryKey;autoIncrement" json:"id"`
	Label string `gorm:"size:100;not null;unique" json:"label"`
}

type Flaw struct {
	ID    int64  `gorm:"primaryKey;autoIncrement" json:"id"`
	Label string `gorm:"size:100;not null;unique" json:"label"`
}

type Weakness struct {
	ID    int64  `gorm:"primaryKey;autoIncrement" json:"id"`
	Label string `gorm:"size:100;not null;unique" json:"label"`
}

type Fear struct {
	ID    int64  `gorm:"primaryKey;autoIncrement" json:"id"`
	Label string `gorm:"size:100;not null;unique" json:"label"`
}
