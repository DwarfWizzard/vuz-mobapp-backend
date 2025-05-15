package educore

type Department struct {
	ID        uint32       `gorm:"primaryKey;type:serial"`
	Faculty   *DictFaculty `gorm:"foreignKey:FacultyId"`
	FacultyId uint32
	Name      string
}
