package main

import (
	"english_exam_go/infrastructure/data/entities"
	"english_exam_go/infrastructure/data/repositories"
	"english_exam_go/utils/app_logger"
	"github.com/joho/godotenv"
)

// RDB auto migrate
func main() {
	if err := godotenv.Load(".env"); err != nil {
		panic("Can not loading .env" + err.Error())
	}

	app_logger.Init()

	repositories.OpenDatabase()

	db := repositories.GetConn()

	if db != nil {
		err := db.Debug().AutoMigrate(
			&entities.User{},
			&entities.Exam{},
			&entities.ExamInvite{},
			&entities.ExamResult{},
			&entities.QuestionAnswer{},
			&entities.ExamQuestion{},
		)
		//db.Model(&entities.ExamQuestion{}).AddForeignKey("question_answer_id", "question_answers(id)", "CASCADE", "CASCADE")
		if err != nil {
			return
		}
	}
}
