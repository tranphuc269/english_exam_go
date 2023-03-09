package dtos

type CreateExamSubmit struct {
	ExamId            int                `json:"exam_id"`
	SubmissionResults []SubmissionResult `json:"submission_results"`
}

type SubmissionResult struct {
	QuestionId int `json:"question_id"`
	AnswerId   int `json:"answer_id"`
}
