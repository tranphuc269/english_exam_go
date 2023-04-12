package dtos

type CreateExamSubmit struct {
	ExamId            int                `json:"exam_id"`
	TakerID           int                `json:"taker_id"`
	TabSwitchCount    int                `json:"tab_switch_count"`
	SubmissionResults []SubmissionResult `json:"submission_results"`
}

type SubmissionResult struct {
	QuestionId int `json:"question_id"`
	AnswerId   int `json:"answer_id"`
}
