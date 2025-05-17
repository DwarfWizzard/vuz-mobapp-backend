package migration

import (
	"time"

	"github.com/DwarfWizzard/vuz-mobapp-backend/internal/education/model"
	"github.com/DwarfWizzard/vuz-mobapp-backend/pkg/pggorm"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func Seed(dbClient *pggorm.Db) error {
	loc, err := time.LoadLocation("Europe/Moscow")
	if err != nil {
		return err
	}

	db := dbClient.DB().Clauses(clause.OnConflict{DoNothing: true})

	err = db.Transaction(func(tx *gorm.DB) error {
		direcions := []model.Direction{
			{ID: 1, Name: "Прикладная математика и информатика"},
		}

		if err := tx.Create(direcions).Error; err != nil {
			return err
		}

		faculties := []model.Faculty{
			{ID: 1, Name: "Факультет математики и компьютерных наук"},
		}

		if err := tx.Create(faculties).Error; err != nil {
			return err
		}

		grades := []model.Grade{
			{ID: 1, Name: "Бакалавриат"},
		}

		if err := tx.Create(grades).Error; err != nil {
			return err
		}

		depratments := []model.Department{
			{ID: 1, FacultyID: 1, Name: "Прикладная математика"},
		}

		if err := tx.Create(depratments).Error; err != nil {
			return err
		}

		groups := []model.EduGroup{
			{ID: 1, FacultyID: 1, DirectionID: 1, GradeID: 1, DepartmentID: 1, Name: "11", Number: "123321", Semester: 2, SemesterStart: time.Date(2025, time.February, 12, 0, 0, 0, 0, loc)},
		}

		if err := tx.Create(groups).Error; err != nil {
			return err
		}

		userGroups := []model.UserGroup{
			{UsersID: 1, EduGroupsID: 1},
		}

		if err := tx.Create(userGroups).Error; err != nil {
			return err
		}

		subjects := []model.Subject{
			{ID: 1, Name: "Математический анализ"},
			{ID: 2, Name: "Алгебра"},
			{ID: 3, Name: "Теория чисел"},
			{ID: 4, Name: "Аглийский язык"},
		}

		if err := tx.Create(subjects).Error; err != nil {
			return err
		}

		itemTypes := []model.ScheduleItemType{
			{ID: 1, Name: "Лекция"},
			{ID: 2, Name: "Практика"},
		}

		if err := tx.Create(itemTypes).Error; err != nil {
			return err
		}

		locations := []model.Location{
			{ID: 1, Address: "ул. Церетели 16", Building: nil, Level: toStringPointer("5"), Room: toStringPointer("512")},
			{ID: 2, Address: "ул. Церетели 16", Building: nil, Level: toStringPointer("5"), Room: toStringPointer("505")},
			{ID: 3, Address: "ул. Церетели 16", Building: nil, Level: toStringPointer("6"), Room: toStringPointer("612")},
			{ID: 4, Address: "ул. Церетели 16", Building: nil, Level: toStringPointer("6"), Room: toStringPointer("609")},
		}

		if err := tx.Create(locations).Error; err != nil {
			return err
		}

		degree := []model.Degree{
			{ID: 1, Name: "канд. физ.-мат. наук"},
			{ID: 2, Name: "д-р. физ.-мат. наук"},
		}

		if err := tx.Create(degree).Error; err != nil {
			return err
		}

		teachers := []model.Teacher{
			{ID: 1, DegreeID: toIntPointer(2), DepartmentID: 1, Name: "Иванов Иван Иванович"},
			{ID: 2, DegreeID: toIntPointer(2), DepartmentID: 1, Name: "Василенко Игорь Иванович"},
			{ID: 3, DegreeID: toIntPointer(2), DepartmentID: 1, Name: "Албегова Жанна Борисовна"},
			{ID: 4, DegreeID: toIntPointer(1), DepartmentID: 1, Name: "Вознесенская Анна Борисовна"},
		}

		if err := tx.Create(teachers).Error; err != nil {
			return err
		}

		now := time.Now().In(loc).Truncate(24 * time.Hour)

		scheduleTemplates := []model.ScheduleTemplate{
			{ID: 1, WeekDay: 1, WeekParity: "both", TimeStart: time.Date(now.Year(), now.Month(), now.Day(), 9, 0, 0, 0, loc).Format(time.TimeOnly), TimeEnd: time.Date(now.Year(), now.Month(), now.Day(), 10, 30, 0, 0, loc).Format(time.TimeOnly), SubjectID: 1, ItemTypeID: 1, TeacherID: 1, LocationID: 1, GroupID: 1},
			{ID: 2, WeekDay: 1, WeekParity: "even", TimeStart: time.Date(now.Year(), now.Month(), now.Day(), 10, 40, 0, 0, loc).Format(time.TimeOnly), TimeEnd: time.Date(now.Year(), now.Month(), now.Day(), 12, 10, 0, 0, loc).Format(time.TimeOnly), SubjectID: 1, ItemTypeID: 1, TeacherID: 1, LocationID: 1, GroupID: 1},
			{ID: 3, WeekDay: 1, WeekParity: "odd", TimeStart: time.Date(now.Year(), now.Month(), now.Day(), 10, 40, 0, 0, loc).Format(time.TimeOnly), TimeEnd: time.Date(now.Year(), now.Month(), now.Day(), 12, 10, 0, 0, loc).Format(time.TimeOnly), SubjectID: 1, ItemTypeID: 2, TeacherID: 1, LocationID: 1, GroupID: 1},
			{ID: 4, WeekDay: 1, WeekParity: "both", TimeStart: time.Date(now.Year(), now.Month(), now.Day(), 12, 50, 0, 0, loc).Format(time.TimeOnly), TimeEnd: time.Date(now.Year(), now.Month(), now.Day(), 14, 20, 0, 0, loc).Format(time.TimeOnly), SubjectID: 3, ItemTypeID: 1, TeacherID: 3, LocationID: 1, GroupID: 1},
			{ID: 5, WeekDay: 1, WeekParity: "both", TimeStart: time.Date(now.Year(), now.Month(), now.Day(), 14, 30, 0, 0, loc).Format(time.TimeOnly), TimeEnd: time.Date(now.Year(), now.Month(), now.Day(), 16, 0, 0, 0, loc).Format(time.TimeOnly), SubjectID: 3, ItemTypeID: 2, TeacherID: 3, LocationID: 2, GroupID: 1},

			{ID: 6, WeekDay: 2, WeekParity: "even", TimeStart: time.Date(now.Year(), now.Month(), now.Day(), 9, 0, 0, 0, loc).Format(time.TimeOnly), TimeEnd: time.Date(now.Year(), now.Month(), now.Day(), 10, 30, 0, 0, loc).Format(time.TimeOnly), SubjectID: 2, ItemTypeID: 1, TeacherID: 2, LocationID: 1, GroupID: 1},
			{ID: 7, WeekDay: 2, WeekParity: "odd", TimeStart: time.Date(now.Year(), now.Month(), now.Day(), 9, 0, 0, 0, loc).Format(time.TimeOnly), TimeEnd: time.Date(now.Year(), now.Month(), now.Day(), 10, 30, 0, 0, loc).Format(time.TimeOnly), SubjectID: 1, ItemTypeID: 1, TeacherID: 1, LocationID: 1, GroupID: 1},
			{ID: 8, WeekDay: 2, WeekParity: "both", TimeStart: time.Date(now.Year(), now.Month(), now.Day(), 10, 40, 0, 0, loc).Format(time.TimeOnly), TimeEnd: time.Date(now.Year(), now.Month(), now.Day(), 12, 10, 0, 0, loc).Format(time.TimeOnly), SubjectID: 2, ItemTypeID: 2, TeacherID: 2, LocationID: 1, GroupID: 1},
			{ID: 9, WeekDay: 2, WeekParity: "odd", TimeStart: time.Date(now.Year(), now.Month(), now.Day(), 12, 50, 0, 0, loc).Format(time.TimeOnly), TimeEnd: time.Date(now.Year(), now.Month(), now.Day(), 14, 20, 0, 0, loc).Format(time.TimeOnly), SubjectID: 4, ItemTypeID: 1, TeacherID: 4, LocationID: 3, GroupID: 1},

			{ID: 10, WeekDay: 3, WeekParity: "both", TimeStart: time.Date(now.Year(), now.Month(), now.Day(), 9, 0, 0, 0, loc).Format(time.TimeOnly), TimeEnd: time.Date(now.Year(), now.Month(), now.Day(), 10, 30, 0, 0, loc).Format(time.TimeOnly), SubjectID: 3, ItemTypeID: 1, TeacherID: 3, LocationID: 1, GroupID: 1},
			{ID: 11, WeekDay: 3, WeekParity: "both", TimeStart: time.Date(now.Year(), now.Month(), now.Day(), 10, 40, 0, 0, loc).Format(time.TimeOnly), TimeEnd: time.Date(now.Year(), now.Month(), now.Day(), 12, 10, 0, 0, loc).Format(time.TimeOnly), SubjectID: 3, ItemTypeID: 2, TeacherID: 3, LocationID: 2, GroupID: 1},
			{ID: 12, WeekDay: 3, WeekParity: "even", TimeStart: time.Date(now.Year(), now.Month(), now.Day(), 12, 50, 0, 0, loc).Format(time.TimeOnly), TimeEnd: time.Date(now.Year(), now.Month(), now.Day(), 14, 20, 0, 0, loc).Format(time.TimeOnly), SubjectID: 1, ItemTypeID: 1, TeacherID: 1, LocationID: 1, GroupID: 1},
			{ID: 13, WeekDay: 3, WeekParity: "odd", TimeStart: time.Date(now.Year(), now.Month(), now.Day(), 12, 50, 0, 0, loc).Format(time.TimeOnly), TimeEnd: time.Date(now.Year(), now.Month(), now.Day(), 14, 20, 0, 0, loc).Format(time.TimeOnly), SubjectID: 2, ItemTypeID: 1, TeacherID: 2, LocationID: 1, GroupID: 1},
			{ID: 14, WeekDay: 3, WeekParity: "both", TimeStart: time.Date(now.Year(), now.Month(), now.Day(), 14, 30, 0, 0, loc).Format(time.TimeOnly), TimeEnd: time.Date(now.Year(), now.Month(), now.Day(), 16, 0, 0, 0, loc).Format(time.TimeOnly), SubjectID: 4, ItemTypeID: 2, TeacherID: 4, LocationID: 3, GroupID: 1},

			{ID: 15, WeekDay: 4, WeekParity: "both", TimeStart: time.Date(now.Year(), now.Month(), now.Day(), 9, 0, 0, 0, loc).Format(time.TimeOnly), TimeEnd: time.Date(now.Year(), now.Month(), now.Day(), 10, 30, 0, 0, loc).Format(time.TimeOnly), SubjectID: 2, ItemTypeID: 1, TeacherID: 2, LocationID: 1, GroupID: 1},
			{ID: 16, WeekDay: 4, WeekParity: "even", TimeStart: time.Date(now.Year(), now.Month(), now.Day(), 10, 40, 0, 0, loc).Format(time.TimeOnly), TimeEnd: time.Date(now.Year(), now.Month(), now.Day(), 12, 10, 0, 0, loc).Format(time.TimeOnly), SubjectID: 2, ItemTypeID: 2, TeacherID: 2, LocationID: 1, GroupID: 1},
			{ID: 17, WeekDay: 4, WeekParity: "odd", TimeStart: time.Date(now.Year(), now.Month(), now.Day(), 10, 40, 0, 0, loc).Format(time.TimeOnly), TimeEnd: time.Date(now.Year(), now.Month(), now.Day(), 12, 10, 0, 0, loc).Format(time.TimeOnly), SubjectID: 4, ItemTypeID: 2, TeacherID: 4, LocationID: 3, GroupID: 1},
			{ID: 18, WeekDay: 4, WeekParity: "both", TimeStart: time.Date(now.Year(), now.Month(), now.Day(), 12, 50, 0, 0, loc).Format(time.TimeOnly), TimeEnd: time.Date(now.Year(), now.Month(), now.Day(), 14, 20, 0, 0, loc).Format(time.TimeOnly), SubjectID: 1, ItemTypeID: 2, TeacherID: 1, LocationID: 1, GroupID: 1},
			{ID: 19, WeekDay: 4, WeekParity: "odd", TimeStart: time.Date(now.Year(), now.Month(), now.Day(), 14, 30, 0, 0, loc).Format(time.TimeOnly), TimeEnd: time.Date(now.Year(), now.Month(), now.Day(), 16, 0, 0, 0, loc).Format(time.TimeOnly), SubjectID: 3, ItemTypeID: 1, TeacherID: 3, LocationID: 1, GroupID: 1},

			{ID: 20, WeekDay: 5, WeekParity: "even", TimeStart: time.Date(now.Year(), now.Month(), now.Day(), 9, 0, 0, 0, loc).Format(time.TimeOnly), TimeEnd: time.Date(now.Year(), now.Month(), now.Day(), 10, 30, 0, 0, loc).Format(time.TimeOnly), SubjectID: 4, ItemTypeID: 1, TeacherID: 4, LocationID: 3, GroupID: 1},
			{ID: 21, WeekDay: 5, WeekParity: "odd", TimeStart: time.Date(now.Year(), now.Month(), now.Day(), 9, 0, 0, 0, loc).Format(time.TimeOnly), TimeEnd: time.Date(now.Year(), now.Month(), now.Day(), 10, 30, 0, 0, loc).Format(time.TimeOnly), SubjectID: 1, ItemTypeID: 1, TeacherID: 1, LocationID: 1, GroupID: 1},
			{ID: 22, WeekDay: 5, WeekParity: "both", TimeStart: time.Date(now.Year(), now.Month(), now.Day(), 10, 40, 0, 0, loc).Format(time.TimeOnly), TimeEnd: time.Date(now.Year(), now.Month(), now.Day(), 12, 10, 0, 0, loc).Format(time.TimeOnly), SubjectID: 2, ItemTypeID: 2, TeacherID: 2, LocationID: 2, GroupID: 1},
		}

		if err := tx.Create(scheduleTemplates).Error; err != nil {
			return err
		}

		speakers := []model.EventSpeaker{
			{ID: 1, Name: "Петров Василий Иванович", Bio: "директор SoftwareLab"},
			{ID: 2, Name: "Иванов Иван Иванович", Bio: "Доктор физико-математических наук, професор ФМКН"},
		}

		if err := tx.Create(speakers).Error; err != nil {
			return err
		}

		events := []model.Event{
			{ID: 1, LocationID: 4, FacultyID: 1, Title: "Будущее в IT", Description: "Лекция от директора SoftwareLab посвященная прогнозам, которые ожидает IT-индустрия", Date: "2025-06-23", RegistrationEnabled: false, TimeStart: "13:00", TimeEnd: "14:55"},
			{ID: 2, LocationID: 1, FacultyID: 1, Title: "Современные проблемы математического анализа", Description: "Лекция професора ФМКН посвященная современным проблемам математического анализа", Date: "2025-07-25", RegistrationEnabled: false, TimeStart: "13:00", TimeEnd: "16:30"},
		}

		if err := tx.Create(events).Error; err != nil {
			return err
		}

		speakersEvent := []model.SpeakerEvent{
			{SpeakerID: 1, EventID: 1},
			{SpeakerID: 2, EventID: 2},
		}

		if err := tx.Create(speakersEvent).Error; err != nil {
			return err
		}

		activities := []model.EventActivity{
			{ID: 1, EventID: 1, Time: "13:00", Activity: "Первая часть"},
			{ID: 2, EventID: 1, Time: "13:45", Activity: "Кофебрейк"},
			{ID: 3, EventID: 1, Time: "14:00", Activity: "Вторая часть"},
			{ID: 4, EventID: 1, Time: "14:40", Activity: "Вопросы-ответы"},

			{ID: 5, EventID: 2, Time: "13:00", Activity: "Вступление"},
			{ID: 6, EventID: 2, Time: "13:15", Activity: "Первая часть"},
			{ID: 7, EventID: 2, Time: "14:45", Activity: "Кофебрейк"},
			{ID: 8, EventID: 2, Time: "15:00", Activity: "Вторая часть"},
		}

		if err := tx.Create(activities).Error; err != nil {
			return err
		}

		comments := []model.EventComment{
			{ID: 1, UsersID: 1, EventID: 1, Text: "Интересно"},
			{ID: 2, UsersID: 1, EventID: 2, Text: "Тоже интересно"},
		}

		if err := tx.Create(comments).Error; err != nil {
			return err
		}

		return nil
	})

	return err
}

func toStringPointer(v string) *string {
	return &v
}

func toIntPointer(v uint32) *uint32 {
	return &v
}

// scheduleTemplates := []model.ScheduleTemplate{
// 			{ID: 1, WeekDay: 0, WeekParity: "both", TimeStart: time.Date(now.Year(), now.Month(), now.Day(), 9, 0, 0, 0, loc).Format(time.TimeOnly), TimeEnd: time.Date(now.Year(), now.Month(), now.Day(), 10, 30, 0, 0, loc).Format(time.TimeOnly), SubjectID: 1, ItemTypeID: 1, TeacherID: 1, LocationID: 1, GroupID: 1},
// 			{ID: 2, WeekDay: 0, WeekParity: "even", TimeStart: time.Date(now.Year(), now.Month(), now.Day(), 10, 40, 0, 0, loc).Format(time.TimeOnly), TimeEnd: time.Date(now.Year(), now.Month(), now.Day(), 12, 10, 0, 0, loc).Format(time.TimeOnly), SubjectID: 1, ItemTypeID: 1, TeacherID: 1, LocationID: 1, GroupID: 1},
// 			{ID: 3, WeekDay: 0, WeekParity: "odd", TimeStart: time.Date(now.Year(), now.Month(), now.Day(), 10, 40, 0, 0, loc).Format(time.TimeOnly), TimeEnd: time.Date(now.Year(), now.Month(), now.Day(), 12, 10, 0, 0, loc).Format(time.TimeOnly), SubjectID: 1, ItemTypeID: 2, TeacherID: 1, LocationID: 1, GroupID: 1},
// 			{ID: 4, WeekDay: 0, WeekParity: "both", TimeStart: time.Date(now.Year(), now.Month(), now.Day(), 12, 50, 0, 0, loc).Format(time.TimeOnly), TimeEnd: time.Date(now.Year(), now.Month(), now.Day(), 14, 20, 0, 0, loc).Format(time.TimeOnly), SubjectID: 3, ItemTypeID: 1, TeacherID: 3, LocationID: 1, GroupID: 1},
// 			{ID: 5, WeekDay: 0, WeekParity: "both", TimeStart: time.Date(now.Year(), now.Month(), now.Day(), 14, 30, 0, 0, loc).Format(time.TimeOnly), TimeEnd: time.Date(now.Year(), now.Month(), now.Day(), 16, 0, 0, 0, loc).Format(time.TimeOnly), SubjectID: 3, ItemTypeID: 2, TeacherID: 3, LocationID: 2, GroupID: 1},

// 			{ID: 6, WeekDay: 1, WeekParity: "even", TimeStart: time.Date(now.Year(), now.Month(), now.Day(), 9, 0, 0, 0, loc).Format(time.TimeOnly), TimeEnd: time.Date(now.Year(), now.Month(), now.Day(), 10, 30, 0, 0, loc).Format(time.TimeOnly), SubjectID: 2, ItemTypeID: 1, TeacherID: 2, LocationID: 1, GroupID: 1},
// 			{ID: 7, WeekDay: 1, WeekParity: "odd", TimeStart: time.Date(now.Year(), now.Month(), now.Day(), 9, 0, 0, 0, loc).Format(time.TimeOnly), TimeEnd: time.Date(now.Year(), now.Month(), now.Day(), 10, 30, 0, 0, loc).Format(time.TimeOnly), SubjectID: 1, ItemTypeID: 1, TeacherID: 1, LocationID: 1, GroupID: 1},
// 			{ID: 8, WeekDay: 1, WeekParity: "both", TimeStart: time.Date(now.Year(), now.Month(), now.Day(), 10, 40, 0, 0, loc).Format(time.TimeOnly), TimeEnd: time.Date(now.Year(), now.Month(), now.Day(), 12, 10, 0, 0, loc).Format(time.TimeOnly), SubjectID: 2, ItemTypeID: 2, TeacherID: 2, LocationID: 1, GroupID: 1},
// 			{ID: 9, WeekDay: 1, WeekParity: "odd", TimeStart: time.Date(now.Year(), now.Month(), now.Day(), 12, 50, 0, 0, loc).Format(time.TimeOnly), TimeEnd: time.Date(now.Year(), now.Month(), now.Day(), 14, 20, 0, 0, loc).Format(time.TimeOnly), SubjectID: 4, ItemTypeID: 1, TeacherID: 4, LocationID: 3, GroupID: 1},

// 			{ID: 10, WeekDay: 2, WeekParity: "both", TimeStart: time.Date(now.Year(), now.Month(), now.Day(), 9, 0, 0, 0, loc).Format(time.TimeOnly), TimeEnd: time.Date(now.Year(), now.Month(), now.Day(), 10, 30, 0, 0, loc).Format(time.TimeOnly), SubjectID: 3, ItemTypeID: 1, TeacherID: 3, LocationID: 1, GroupID: 1},
// 			{ID: 11, WeekDay: 2, WeekParity: "both", TimeStart: time.Date(now.Year(), now.Month(), now.Day(), 10, 40, 0, 0, loc).Format(time.TimeOnly), TimeEnd: time.Date(now.Year(), now.Month(), now.Day(), 12, 10, 0, 0, loc).Format(time.TimeOnly), SubjectID: 3, ItemTypeID: 2, TeacherID: 3, LocationID: 2, GroupID: 1},
// 			{ID: 12, WeekDay: 2, WeekParity: "even", TimeStart: time.Date(now.Year(), now.Month(), now.Day(), 12, 50, 0, 0, loc).Format(time.TimeOnly), TimeEnd: time.Date(now.Year(), now.Month(), now.Day(), 14, 20, 0, 0, loc).Format(time.TimeOnly), SubjectID: 1, ItemTypeID: 1, TeacherID: 1, LocationID: 1, GroupID: 1},
// 			{ID: 13, WeekDay: 2, WeekParity: "odd", TimeStart: time.Date(now.Year(), now.Month(), now.Day(), 12, 50, 0, 0, loc).Format(time.TimeOnly), TimeEnd: time.Date(now.Year(), now.Month(), now.Day(), 14, 20, 0, 0, loc).Format(time.TimeOnly), SubjectID: 2, ItemTypeID: 1, TeacherID: 2, LocationID: 1, GroupID: 1},
// 			{ID: 14, WeekDay: 2, WeekParity: "both", TimeStart: time.Date(now.Year(), now.Month(), now.Day(), 14, 30, 0, 0, loc).Format(time.TimeOnly), TimeEnd: time.Date(now.Year(), now.Month(), now.Day(), 16, 0, 0, 0, loc).Format(time.TimeOnly), SubjectID: 4, ItemTypeID: 2, TeacherID: 4, LocationID: 3, GroupID: 1},

// 			{ID: 15, WeekDay: 3, WeekParity: "both", TimeStart: time.Date(now.Year(), now.Month(), now.Day(), 9, 0, 0, 0, loc).Format(time.TimeOnly), TimeEnd: time.Date(now.Year(), now.Month(), now.Day(), 10, 30, 0, 0, loc).Format(time.TimeOnly), SubjectID: 2, ItemTypeID: 1, TeacherID: 2, LocationID: 1, GroupID: 1},
// 			{ID: 16, WeekDay: 3, WeekParity: "even", TimeStart: time.Date(now.Year(), now.Month(), now.Day(), 10, 40, 0, 0, loc).Format(time.TimeOnly), TimeEnd: time.Date(now.Year(), now.Month(), now.Day(), 12, 10, 0, 0, loc).Format(time.TimeOnly), SubjectID: 2, ItemTypeID: 2, TeacherID: 2, LocationID: 1, GroupID: 1},
// 			{ID: 17, WeekDay: 3, WeekParity: "odd", TimeStart: time.Date(now.Year(), now.Month(), now.Day(), 10, 40, 0, 0, loc).Format(time.TimeOnly), TimeEnd: time.Date(now.Year(), now.Month(), now.Day(), 12, 10, 0, 0, loc).Format(time.TimeOnly), SubjectID: 4, ItemTypeID: 2, TeacherID: 4, LocationID: 3, GroupID: 1},
// 			{ID: 18, WeekDay: 3, WeekParity: "both", TimeStart: time.Date(now.Year(), now.Month(), now.Day(), 12, 50, 0, 0, loc).Format(time.TimeOnly), TimeEnd: time.Date(now.Year(), now.Month(), now.Day(), 14, 20, 0, 0, loc).Format(time.TimeOnly), SubjectID: 1, ItemTypeID: 2, TeacherID: 1, LocationID: 1, GroupID: 1},
// 			{ID: 19, WeekDay: 3, WeekParity: "odd", TimeStart: time.Date(now.Year(), now.Month(), now.Day(), 14, 30, 0, 0, loc).Format(time.TimeOnly), TimeEnd: time.Date(now.Year(), now.Month(), now.Day(), 16, 0, 0, 0, loc).Format(time.TimeOnly), SubjectID: 3, ItemTypeID: 1, TeacherID: 3, LocationID: 1, GroupID: 1},

// 			{ID: 20, WeekDay: 4, WeekParity: "even", TimeStart: time.Date(now.Year(), now.Month(), now.Day(), 9, 0, 0, 0, loc).Format(time.TimeOnly), TimeEnd: time.Date(now.Year(), now.Month(), now.Day(), 10, 30, 0, 0, loc).Format(time.TimeOnly), SubjectID: 4, ItemTypeID: 1, TeacherID: 4, LocationID: 3, GroupID: 1},
// 			{ID: 21, WeekDay: 4, WeekParity: "odd", TimeStart: time.Date(now.Year(), now.Month(), now.Day(), 9, 0, 0, 0, loc).Format(time.TimeOnly), TimeEnd: time.Date(now.Year(), now.Month(), now.Day(), 10, 30, 0, 0, loc).Format(time.TimeOnly), SubjectID: 1, ItemTypeID: 1, TeacherID: 1, LocationID: 1, GroupID: 1},
// 			{ID: 22, WeekDay: 4, WeekParity: "both", TimeStart: time.Date(now.Year(), now.Month(), now.Day(), 10, 40, 0, 0, loc).Format(time.TimeOnly), TimeEnd: time.Date(now.Year(), now.Month(), now.Day(), 12, 10, 0, 0, loc).Format(time.TimeOnly), SubjectID: 2, ItemTypeID: 2, TeacherID: 2, LocationID: 2, GroupID: 1},
// 		}
