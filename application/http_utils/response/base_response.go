package response

type BaseResponse struct {
	data interface{}
	meta *Metadata
}

type Metadata struct {
	Code      int              `json:"code"`
	Page      int              `json:"page"`
	PageSize  int              `json:"page_size"`
	Total     int              `json:"total"`
	Errors    []FieldViolation `json:"errors"`
	Message   string           `json:"message"`
	RequestId string           `json:"request_id"`
}
