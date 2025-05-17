package model

type Teacher struct {
	ID           uint32  `gorm:"primaryKey;type:int;autoIncrement" json:"id"`
	DegreeID     *uint32 `json:"degree_id,omitempty"`
	DepartmentID uint32  `gorm:"not null" json:"department_id"`
	Name         string  `gorm:"not null" json:"name"`
	Position     string  `gorm:"not null" json:"position"`

	Degree     *Degree     `gorm:"foreignKey:DegreeID;references:ID" json:"degree,omitempty"`
	Department *Department `gorm:"foreignKey:DepartmentID;references:ID;constraint:OnDelete:CASCADE" json:"department,omitempty"`
}

type Degree struct {
	ID   uint32 `gorm:"primaryKey;type:int;autoIncrement" json:"id"`
	Name string `gorm:"uniqueIndex;not null" json:"name"`
}
