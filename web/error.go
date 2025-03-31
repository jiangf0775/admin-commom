package web

type Error struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// errors.error 接口方法
func (e *Error) Error() string {
	return e.Msg
}

func NewError(code int, msg string) *Error {
	return &Error{Code: code, Msg: msg}
}
