package exception

type systemError interface {
	IsInternal() bool
	Code() string
}
