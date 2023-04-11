package dtos

import "english_exam_go/infrastructure/data/entities"

type AddTakerToExam struct {
	ExamID  int   `json:"exam_id"`
	UserIds []int `json:"user_ids"`
}

func (req AddTakerToExam) ToListTakerEntity() []*entities.ExamTasker {
	var examTakerEntities []*entities.ExamTasker
	for _, userId := range req.UserIds {
		examTakerEntities = append(examTakerEntities, &entities.ExamTasker{
			ExamID: req.ExamID,
			UserID: userId,
		})
	}
	return examTakerEntities
}
