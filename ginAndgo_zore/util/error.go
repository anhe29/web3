package util

type Error struct {
	Code int
	Msg  string
}

func NewError(code int, msg string) *Error {
	return &Error{
		Code: code,
		Msg:  msg,
	}
}

func (e *Error) Error() string {
	return e.Msg
}
