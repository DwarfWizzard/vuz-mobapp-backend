package model

type Location struct {
	ID       uint64  `gorm:"primaryKey;type:bigint;autoIncrement" json:"id"`
	Address  string  `json:"address"`
	Building *string `json:"building,omitempty"`
	Level    *string `json:"level,omitempty"`
	Room     *string `json:"room,omitempty"`
}
