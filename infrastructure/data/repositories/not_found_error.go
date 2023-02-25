package repositories

import "net/http"

const DefaultNotFoundMsg string = "Không tìm thấy bản ghi tương ứng"

// NotFoundError
type NotFoundError struct {
	Msg           string
	ErrMsg        string
	OriginalError error
}

func (ne *NotFoundError) Message() []string {
	return []string{ne.Msg}
}

func (ne *NotFoundError) Code() int {
	return http.StatusNotFound
}

func (ne *NotFoundError) Error() string {
	if ne.OriginalError != nil {
		return ne.ErrMsg + " , " + ne.OriginalError.Error()
	}
	return ne.ErrMsg
}
