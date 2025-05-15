package educore

type DictFaculty struct {
	ID   uint32 `gorm:"primaryKey;type:serial"`
	Name string `gorm:"uniqueIndex"`
}

type DictDirection struct {
	ID   uint32 `gorm:"primaryKey;type:serial"`
	Name string `gorm:"uniqueIndex"`
}

type DictGrade struct {
	ID   uint32 `gorm:"primaryKey;type:serial"`
	Name string `gorm:"uniqueIndex"`
}
