package repositories

import "net/http"

type OptimisticLockError struct {
	Msg    string
	ErrMsg string
}

func (ole *OptimisticLockError) Code() string {
	return "optimistic_lock_error"
}

func (ole *OptimisticLockError) Message() []string {
	return []string{ole.Msg}
}

func (ole *OptimisticLockError) HttpStatus() int {
	return http.StatusConflict
}

func (ole *OptimisticLockError) Error() string {
	return ole.ErrMsg
}
