package model

import "time"

type ScheduleTemplate struct {
	ID         uint64       `gorm:"primaryKey;autoIncrement"`
	WeekDay    time.Weekday `gorm:"not null"`
	WeekParity string       `gorm:"not null"`
	TimeStart  string       `gorm:"not null"`
	TimeEnd    string       `gorm:"not null"`

	ItemTypeID uint32 `gorm:"not_null"`
	SubjectID  uint32 `gorm:"not null"`
	LocationID uint32 `gorm:"not null"`
	TeacherID  uint32 `gorm:"not null"`
	GroupID    uint32 `gorm:"not null"`

	ItemType *ScheduleItemType `gorm:"foreignKey:ItemTypeID;references:ID"`
	Subject  *Subject          `gorm:"foreignKey:SubjectID;references:ID;constraint:OnDelete:CASCADE"`
	Location *Location         `gorm:"foreignKey:LocationID;references:ID;constraint:OnDelete:CASCADE"`
	Group    *EduGroup         `gorm:"foreignKey:GroupID;references:ID;constraint:OnDelete:CASCADE"`
	Teacher  *Teacher          `gorm:"foreignKey:TeacherID;references:ID;constraint:OnDelete:CASCADE"`
}

type ScheduleOverride struct {
	ID         uint64    `gorm:"primaryKey;autoIncrement"`
	TemplateID uint64    `gorm:"not null"`
	Date       time.Time `gorm:"not null"`

	SubjectID  *uint32
	LocationID *uint32
	TeacherID  *uint32
	Note       *string
	Canceled   bool `gorm:"default:false"`

	Template *ScheduleTemplate `gorm:"foreignKey:TemplateID;references:ID;constraint:OnDelete:CASCADE"`
}

type ScheduleItemType struct {
	ID   uint32 `gorm:"primaryKey;type:bigint;autoIncrement"`
	Name string `gorm:"uniqueIndex;not null"`
}

type Subject struct {
	ID   uint32 `gorm:"primaryKey;type:bigint;autoIncrement"`
	Name string `gorm:"uniqueIndex;not null"`
}
