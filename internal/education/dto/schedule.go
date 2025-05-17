package dto

import "time"

type Schedule struct {
	TemplateID uint64       `json:"-"`
	Date       string       `json:"date"`
	TimeStart  string       `json:"time_start"`
	TimeEnd    string       `json:"time_end"`
	WeekDay    time.Weekday `json:"-"`
	Canceled   bool         `json:"is_canceled"`

	GroupID     uint32 `json:"group_id"`
	GroupNumber string `json:"group_number"`

	SubjectID   *uint32 `json:"subject_id"`
	SubjectName string `json:"subject_name"`

	TeacherID   *uint32 `json:"teacher_id"`
	TeacherName string `json:"teacher_name"`

	LocationID *uint32 `json:"location_id"`
	Address    string `json:"address"`

	Note *string `json:"note,omitempty"`
}
