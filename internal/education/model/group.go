package model

import "time"

type EduGroup struct {
	ID            uint32    `gorm:"primaryKey;type:int;autoIncrement" json:"id"`
	FacultyID     uint32    `gorm:"not null" json:"faculty_id"`
	DirectionID   uint32    `gorm:"not null" json:"direction_id"`
	GradeID       uint32    `gorm:"not null" json:"grade_id"`
	DepartmentID  uint32    `gorm:"not null" json:"department_id"`
	Name          string    `gorm:"not null" json:"name"`
	Number        string    `gorm:"uniqueIndex" json:"number"`
	Semester      int8      `gorm:"not null" json:"semester"`
	SemesterStart time.Time `gorm:"not null" json:"semester_start"`

	Faculty    *Faculty    `gorm:"foreignKey:FacultyID;references:ID;constraint:OnDelete:CASCADE" json:"faculty,omitempty"`
	Direction  *Direction  `gorm:"foreignKey:DirectionID;references:ID;constraint:OnDelete:CASCADE" json:"direction,omitempty"`
	Grade      *Grade      `gorm:"foreignKey:GradeID;references:ID" json:"grade,omitempty"`
	Department *Department `gorm:"foreignKey:DepartmentID;references:ID;constraint:OnDelete:CASCADE" json:"department,omitempty"`
}

type UserGroup struct {
	UsersID     uint32 `gorm:"primaryKey;constraint:OnDelete:CASCADE" json:"users_id"`
	EduGroupsID uint32 `gorm:"primaryKey;constraint:OnDelete:CASCADE" json:"edu_groups_id"`
}

func (g *EduGroup) IsEvenWeek(date time.Time) bool {
	date = time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
	start := time.Date(g.SemesterStart.Year(), g.SemesterStart.Month(), g.SemesterStart.Day(), 0, 0, 0, 0, g.SemesterStart.Location())

	offset := int(start.Weekday() - time.Monday)
	if offset < 0 {
		offset += 7
	}
	start = start.AddDate(0, 0, -offset)

	daysPassed := int(date.Sub(start).Hours() / 24)
	weekNum := daysPassed/7 + 1

	return weekNum%2 == 0
}
