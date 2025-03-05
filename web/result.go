package web

type Result struct {
	Code int    `json:"message"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func SuccessNo() *Result {
	return &Result{
		Code: OK,
		Msg:  "success",
		Data: nil,
	}
}

func Success(data any) *Result {
	return &Result{
		Code: OK,
		Msg:  "success",
		Data: data,
	}
}

func Fail(err *Error) *Result {
	return &Result{
		Code: err.Code,
		Msg:  err.Msg,
	}
}
