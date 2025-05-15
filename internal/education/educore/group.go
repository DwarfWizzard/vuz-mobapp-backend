package educore

type Group struct {
	ID           uint32       `gorm:"primaryKey;type:serial"`
	Faculty      *DictFaculty `gorm:"foreignKey:FacultyId"`
	FacultyId    uint32
	Direction    *DictDirection `gorm:"foreignKey:DirectionId"`
	DirectionId  uint32
	Grade        *DictGrade `gorm:"foreignKey:GradeId"`
	GradeId      uint32
	Department   *Department `gorm:"foreignKey:DepartmentId"`
	DepartmentId uint32
	Name         string
	Number       string `gorm:"uniqueIndex"`
	Course       int8
}
