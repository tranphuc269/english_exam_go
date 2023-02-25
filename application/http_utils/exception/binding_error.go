package exception

import "net/http"

type BindingError struct {
	ErrMsg        string
	OriginalError error
}

func (be *BindingError) Code() string {
	return "binding_error"
}

func (be *BindingError) Message() []string {
	return []string{"Request chứa các giá trị đầu vào không hợp lệ. Vui lòng kiểm tra lại đầu vào của bạn.\n"}
}

func (be *BindingError) HTTPStatus() int {
	return http.StatusBadRequest
}

func (be *BindingError) Error() string {
	return be.ErrMsg + " , " + be.OriginalError.Error()
}
