package model

type Department struct {
	ID        uint32 `gorm:"primaryKey;type:int;autoIncrement" json:"id"`
	FacultyID uint32 `gorm:"not null" json:"faculty_id"`
	Name      string `gorm:"not null" json:"name"`

	Faculty *Faculty `gorm:"foreignKey:FacultyID" json:"faculty,omitempty"`
}
