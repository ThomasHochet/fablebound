package models

type Character struct {
	ID         int64  `gorm:"primaryKey;autoIncrement" json:"id"`
	Firstname  string `gorm:"size:150;not null" json:"firstname"`
	Middlename string `gorm:"size:200" json:"middlename,omitempty"`
	Surname    string `gorm:"size:150;not null" json:"surname"`
	Nickname   string `gorm:"size:150" json:"nickname,omitempty"` // e.g. "The Iron Hound"

	// Basic Identity & Status
	Age         string `gorm:"size:50;default:'Unknown'" json:"age"` // String for "Unknown", "240", "Ageless"
	Description string `gorm:"type:text" json:"description,omitempty"`
	Portrait    []byte `gorm:"type:blob" json:"portrait,omitempty"` // Image path or URL
	Goals       string `gorm:"type:text" json:"goals,omitempty"`

	// Foreign Keys for Single-Choice Lookups (Nullable pointers if optional)
	// Example for me and my smooth brain : GenderID = (in SQL) gender_id BIGINT NOT_NULL FOREIGN KEY
	GenderID     *int64 `gorm:"index" json:"gender_id,omitempty"`
	RaceID       *int64 `gorm:"index" json:"race_id,omitempty"`
	OccupationID *int64 `gorm:"index" json:"occupation_id,omitempty"`
	AlignmentID  *int64 `gorm:"index" json:"alignment_id,omitempty"`
	StatusID     *int64 `gorm:"index" json:"status_id,omitempty"`

	// Foreign Keys for World Locations & Factions
	AffiliationID        *int64 `gorm:"index" json:"affiliation_id,omitempty"`
	FactionID            *int64 `gorm:"index" json:"faction_id,omitempty"`
	BirthplaceLocationID *int64 `gorm:"index" json:"birthplace_location_id,omitempty"`
	CurrentLocationID    *int64 `gorm:"index" json:"current_location_id,omitempty"`

	// GORM Association
	// Other example for me and the smooth brain : Gender *Gender = (in SQL) LEFT JOIN genders on genders.id = characters.gender_id
	Gender     *Gender     `gorm:"foreignKey:GenderID;references:ID" json:"gender,omitempty"`
	Race       *Race       `gorm:"foreignKey:RaceID;references:ID" json:"race,omitempty"`
	Occupation *Occupation `gorm:"foreignKey:OccupationID;references:ID" json:"occupation,omitempty"`
	Alignment  *Alignment  `gorm:"foreignKey:AlignmentID;references:ID" json:"alignment,omitempty"`
	Status     *Status     `gorm:"foreignKey:StatusID;references:ID" json:"status,omitempty"`

	Affiliation     *Affiliation `gorm:"foreignKey:AffiliationID;references:ID" json:"affiliation,omitempty"`
	Faction         *Faction     `gorm:"foreignKey:FactionID;references:ID" json:"faction,omitempty"`
	Birthplace      *Location    `gorm:"foreignKey:BirthplaceLocationID;references:ID" json:"birthplace,omitempty"`
	CurrentLocation *Location    `gorm:"foreignKey:CurrentLocationID;references:ID" json:"current_location,omitempty"`

	// Separate Many-to-Many Categories
	PersonalityTraits []PersonalityTrait `gorm:"many2many:character_personality_traits" json:"personality_traits,omitempty"`
	Strengths         []Strength         `gorm:"many2many:character_strengths" json:"strengths,omitempty"`
	Flaws             []Flaw             `gorm:"many2many:character_flaws" json:"flaws,omitempty"`
	Weaknesses        []Weakness         `gorm:"many2many:character_weaknesses" json:"weaknesses,omitempty"`
	Fears             []Fear             `gorm:"many2many:character_fears" json:"fears,omitempty"`

	// Lore & Character Relationships (Populated via Queries)
	Relationships []Relationship `gorm:"foreignKey:SourceCharID" json:"relationships,omitempty"`
	Backstories   []Lore         `gorm:"foreignKey:CharacterID" json:"backstories,omitempty"`
}
