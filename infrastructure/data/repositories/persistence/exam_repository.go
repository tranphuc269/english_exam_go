package persistence

import (
	"context"
	"english_exam_go/infrastructure/data/entities"
	"english_exam_go/infrastructure/data/repositories"
	"fmt"
	"gorm.io/gorm"
)

type IExamRepository interface {
	CreateExam(context.Context, *entities.Exam) error
	UpdateExam(context.Context, *entities.Exam) error
	FindExamById(context.Context, uint) (*entities.Exam, error)
	FindAllExams(ctx context.Context) ([]*entities.Exam, error)
	FindExamsByCreatorId(context.Context, uint) ([]*entities.Exam, error)
	FindExamsByTaskerId(context.Context, uint) ([]*entities.Exam, error)
}

type ExamRepository struct {
}

func (er ExamRepository) FindAllExams(ctx context.Context) ([]*entities.Exam, error) {
	//TODO implement me
	db := repositories.GetConn()
	var exams []*entities.Exam
	result := db.Order("created_at").Find(&exams)
	if result.Error != nil || result.RowsAffected == 0 {
		return nil, &repositories.NotFoundError{
			Msg:           repositories.DefaultNotFoundMsg,
			ErrMsg:        fmt.Sprintf("[infrastructure.data.repositories.persistence.FindAllExams] failed to find exams from rdb"),
			OriginalError: nil,
		}
	}
	return exams, nil
}

func (er ExamRepository) CreateExam(ctx context.Context, exam *entities.Exam) error {
	//TODO implement me
	db := repositories.GetConn()
	result := db.Create(exam)
	fmt.Printf("result.Error : %s\n", result.Error)
	if result.Error != nil {
		return &repositories.RdbRuntimeError{
			ErrMsg:        fmt.Sprintf("[infrastructure.data.repositories.persistence.CreateExam] fail to insert Exam to Database"),
			OriginalError: result.Error,
		}
	}
	return nil
}

func (er ExamRepository) UpdateExam(ctx context.Context, exam *entities.Exam) error {
	//TODO implement me
	db := repositories.GetConn()
	//err := db.Model(&exam).Association("ExamQuestions").Replace(exam.ExamQuestions)
	//if err != nil {
	//	return err
	//}

	for _, ques := range exam.ExamQuestions {
		ques.ExamId = int(exam.ID)
	}
	result := db.Updates(exam.ExamQuestions)
	_ = db.Updates(exam)
	fmt.Printf("result.Error : %s\n", result.Error)
	if result.Error != nil {
		return &repositories.RdbRuntimeError{
			ErrMsg:        fmt.Sprintf("[infrastructure.data.repositories.persistence.UpdateExam] fail to update Exam to Database"),
			OriginalError: result.Error,
		}
	}
	return nil
}

func (er ExamRepository) FindExamById(ctx context.Context, ID uint) (*entities.Exam, error) {
	//TODO implement me
	db := repositories.GetConn()
	examEnt := &entities.Exam{}
	err := db.Preload("ExamQuestions", func(db *gorm.DB) *gorm.DB {
		return db.Preload("Answers")
	}).First(&examEnt, ID)
	//fmt.Println(err.Error)
	if err.Error != nil {
		return nil, &repositories.NotFoundError{
			Msg:           repositories.DefaultNotFoundMsg,
			ErrMsg:        fmt.Sprintf("[infrastructure.data.repositories.persistence.FindExamById] failed to find ExamEntity from rdb. ID : %d", ID),
			OriginalError: err.Error,
		}
	}
	return examEnt, err.Error
}

func (er ExamRepository) FindExamsByCreatorId(ctx context.Context, UserID uint) ([]*entities.Exam, error) {
	//TODO implement me
	db := repositories.GetConn()
	var exams []*entities.Exam
	result := db.Where("creator_id=?", UserID).Order("created_at").Find(&exams)
	return exams, result.Error
}

func (er ExamRepository) FindExamsByTaskerId(ctx context.Context, UserID uint) ([]*entities.Exam, error) {
	//TODO implement me
	db := repositories.GetConn()
	var exams []*entities.Exam
	result := db.Table("exams").Select("exam.*").
		Joins("JOIN exam_takers ON exams.id = exam_takers.exam_id").
		Where("exam_takers.user_id = ?", UserID).
		Find(&exams)
	return exams, result.Error
}

func CreateExamRepository() IExamRepository {
	return &ExamRepository{}
}
