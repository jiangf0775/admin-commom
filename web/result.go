package web

type Result[T any] struct {
	Code int    `json:"message"`
	Msg  string `json:"msg"`
	Data T      `json:"data"`
}

func Success[T any](data T) *Result[T] {
	return &Result[T]{
		Code: OK,
		Msg:  "success",
		Data: data,
	}
}

func Fail[T any](err *Error) *Result[T] {
	return &Result[T]{
		Code: err.Code,
		Msg:  err.Msg,
	}
}
