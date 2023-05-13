package domain_exception

import (
	"english_exam_go/application/http_utils/exception"
)

type AuthException struct {
}

func (a AuthException) Code() int {
	//TODO implement me
	return 401
}

func (a AuthException) Message() []string {
	//TODO implement me
	return []string{"Not authentication"}
}

func (a AuthException) HTTPStatus() int {
	//TODO implement me
	return 401
}

func CreateAuthException() exception.ApplicationError {
	return AuthException{}
}
