package model

type Faculty struct {
	ID   uint32 `gorm:"primaryKey;type:int;autoIncrement" json:"id"`
	Name string `gorm:"uniqueIndex;not null" json:"name"`
}

type Direction struct {
	ID   uint32 `gorm:"primaryKey;type:int;autoIncrement" json:"id"`
	Name string `gorm:"uniqueIndex;not null" json:"name"`
}

type Grade struct {
	ID   uint32 `gorm:"primaryKey;type:int;autoIncrement" json:"id"`
	Name string `gorm:"uniqueIndex;not null" json:"name"`
}
