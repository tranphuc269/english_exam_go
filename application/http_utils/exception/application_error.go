package exception

type ApplicationError interface {
	Code() int
	Message() []string
	HTTPStatus() int
}

