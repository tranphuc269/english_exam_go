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
	UpdateQuestion(context.Context, *entities.ExamQuestion) error
	FindExamById(context.Context, uint) (*entities.Exam, error)
	FindAllExams(ctx context.Context, offset int, limit int) ([]*entities.Exam, error)
	FindExamsByCreatorId(context.Context, int, int, uint) ([]*entities.Exam, error)
	FindExamsByTaskerId(context.Context, int, int, uint) ([]*entities.Exam, error)
}

type ExamRepository struct {
}

func (er ExamRepository) FindAllExams(ctx context.Context, offset int, limit int) ([]*entities.Exam, error) {
	//TODO implement me
	db := repositories.GetConn()
	var exams []*entities.Exam
	result := db.Offset(offset).Limit(limit).Order("created_at").Find(&exams)
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

	updateData := map[string]interface{}{
		"ExamName":        exam.ExamName,
		"ExamDescription": exam.ExamDescription,
		"ListenFile":      exam.ListenFile,
		"ExamStartTime":   exam.ExamStartTime,
		"ExamEndTime":     exam.ExamEndTime,
	}

	if err := db.Model(&exam).Updates(updateData).Error; err != nil {
		return err
	}

	return nil
}

func (er ExamRepository) UpdateQuestion(ctx context.Context, question *entities.ExamQuestion) error {
	//TODO implement me
	db := repositories.GetConn()

	updateData := map[string]interface{}{
		"QuestionCase": question.QuestionCase,
		"QuestionText": question.QuestionText,
		"File":         question.File,
	}

	for _, ans := range question.Answers {
		updateAns := map[string]interface{}{
			"Content":   ans.Content,
			"IsCorrect": ans.IsCorrect,
		}
		db.Model(&ans).Updates(updateAns)
	}

	if err := db.Model(&question).Updates(updateData).Error; err != nil {
		return err
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

func (er ExamRepository) FindExamsByCreatorId(ctx context.Context, offset int, limit int, UserID uint) ([]*entities.Exam, error) {
	//TODO implement me
	db := repositories.GetConn()
	var exams []*entities.Exam
	result := db.Offset(offset).Limit(limit).Where("creator_id=?", UserID).Order("created_at").Find(&exams)
	return exams, result.Error
}

func (er ExamRepository) FindExamsByTaskerId(ctx context.Context, offset int, limit int, UserID uint) ([]*entities.Exam, error) {
	//TODO implement me
	db := repositories.GetConn()
	var exams []*entities.Exam
	result := db.Offset(offset).Limit(limit).Table("exams").Select("exam.*").
		Joins("JOIN exam_takers ON exams.id = exam_takers.exam_id").
		Where("exam_takers.user_id = ?", UserID).
		Find(&exams)
	return exams, result.Error
}

func CreateExamRepository() IExamRepository {
	return &ExamRepository{}
}
