package models

type Gender struct {
	ID    int64  `gorm:"primaryKey;autoIncrement" json:"id"`
	Label string `gorm:"size:100;not null;unique" json:"label"`
}

type Race struct {
	ID    int64  `gorm:"primaryKey;autoIncrement" json:"id"`
	Label string `gorm:"size:100;not null;unique" json:"label"`
}

type Occupation struct {
	ID    int64  `gorm:"primaryKey;autoIncrement" json:"id"`
	Label string `gorm:"size:100;not null;unique" json:"label"`
}

type Alignment struct {
	ID    int64  `gorm:"primaryKey;autoIncrement" json:"id"`
	Label string `gorm:"size:100;not null;unique" json:"label"`
}

type Status struct {
	ID    int64  `gorm:"primaryKey;autoIncrement" json:"id"`
	Label string `gorm:"size:100;not null;unique" json:"label"`
}
