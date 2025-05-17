package model

import "time"

type Event struct {
	ID                  uint32 `gorm:"primaryKey;autoIncrement" json:"id"`
	LocationID          uint32 `gorm:"not null" json:"-"`
	FacultyID           uint32 `gorm:"not null" json:"-"`
	Title               string `gorm:"not null" json:"title"`
	Description         string `json:"description,omitempty"`
	MediaUrl            string `json:"media_url,omitempty"`
	Date                string `json:"date,omitempty"`
	RegistrationEnabled bool   `json:"-"`
	TimeStart           string `json:"time_start"`
	TimeEnd             string `json:"time_end"`

	RegistrationDeadline *time.Time `json:"-"`

	Faculty    *Faculty         `gorm:"foreignKey:FacultyID;references:ID;constraint:OnDelete:CASCADE" json:"faculty,omitempty"`
	Location   *Location        `gorm:"foreignKey:LocationID;references:ID;constraint:OnDelete:CASCADE" json:"location,omitempty"`
	Speakers   []*EventSpeaker  `gorm:"many2many:speaker_events;joinForeignKey:EventID;JoinReferences:SpeakerID" json:"speakers,omitempty"`
	Activities []*EventActivity `json:"activities,omitempty"`
	Comments   []*EventComment  `json:"comments,omitempty"`
}

type SpeakerEvent struct {
	SpeakerID uint32 `gorm:"primaryKey"`
	EventID   uint32 `gorm:"primaryKey"`
}

type EventSpeaker struct {
	ID       uint32 `gorm:"primaryKey;autoIncrement" json:"id"`
	Name     string `gorm:"not null" json:"name"`
	Bio      string `json:"bio"`
	PhotoURL string `json:"photo_url"`

	Events []*Event `gorm:"many2many:speaker_events;joinForeignKey:SpeakerID;JoinReferences:EventID" json:"-"`
}

type EventActivity struct {
	ID       uint32 `gorm:"primaryKey;autoIncrement" json:"id"`
	EventID  uint32 `gorm:"not null" json:"event_id"`
	Time     string `json:"time"`
	Activity string `json:"activity"`

	Event *Event `gorm:"foreignKey:EventID;references:ID;constraint:OnDelete:CASCADE" json:"-"`
}

type EventComment struct {
	ID      uint32 `gorm:"primaryKey;autoIncrement" json:"id"`
	UsersID uint32 `gorm:"not null" json:"user_id"`
	EventID uint32 `gorm:"not null" json:"event_id"`
	Text    string `gorm:"not null" json:"text"`

	Event *Event `gorm:"foreignKey:EventID;references:ID;constraint:OnDelete:CASCADE" json:"-"`
}

// type Registration struct {
// 	UsersID          uint32 `gorm:"primaryKey"`
// 	EventID          uint32 `gorm:"primaryKey"`
// 	RegistrationDate time.Time

// 	Event *Event `gorm:"foreignKey:EventID;references:ID;constraint:OnDelete:CASCADE"`
// }
