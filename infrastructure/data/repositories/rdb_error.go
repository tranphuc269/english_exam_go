package repositories

type RdbRuntimeError struct {
	ErrMsg        string
	OriginalError error
}

func (rre *RdbRuntimeError) Code() string {
	return "rdb_runtime_error"
}

func (rre *RdbRuntimeError) IsInternal() bool {
	return true
}

func (rre *RdbRuntimeError) Error() string {
	return rre.ErrMsg + " , " + rre.OriginalError.Error()
}
