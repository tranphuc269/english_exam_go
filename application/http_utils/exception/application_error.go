package exception

type applicationError interface {
	Code() string
	Message() []string
	HTTPStatus() int
}
