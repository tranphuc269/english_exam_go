package main

import (
	"english_exam_go/infrastructure/data/entities"
	"english_exam_go/infrastructure/data/repositories"
	"english_exam_go/utils/app_logger"
	"github.com/joho/godotenv"
)

// RDBのauto migrateを行う
func main() {
	if err := godotenv.Load(".env"); err != nil {
		panic("Can not loading .env" + err.Error())
	}

	app_logger.Init()

	repositories.OpenDatabase()

	db := repositories.GetConn()

	if db != nil {
		db.Debug().AutoMigrate(
			&entities.BookEnt{},
		)
	}
}
