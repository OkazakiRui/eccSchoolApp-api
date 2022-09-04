package infrastructure

import (
	"github.com/yumekiti/eccSchoolApp-api/domain"
	"github.com/yumekiti/eccSchoolApp-api/domain/repository"
)

type TimetableRepository struct{}

func NewTimetableRepository() repository.TimetableRepository {
	return &TimetableRepository{}
}

func (r *TimetableRepository) Get(week string, user *domain.User) (*domain.Timetable, error) {
	// c, token := config.FalconLogin(user)

	// 返す値の初期化
	// var Date string
	// var Weekday string
	// var Period []string
	// var SubjectTitle []string
	// var Classroom []string
	// var Teacher []string

	return &domain.Timetable{
		Date:    "2021-01-01",
		Weekday: "月",
		Timetable: []domain.TimetableDetail{
			{
				Period:       "1",
				SubjectTitle: "国語",
				Classroom:    "1",
				Teacher:      "田中",
			},
			{
				Period:       "2",
				SubjectTitle: "数学",
				Classroom:    "2",
				Teacher:      "鈴木",
			},
		},
	}, nil
}