package web

type LogicResult struct{}

func (l *LogicResult) assert(ok *Result, fail *Result, err error) (*Result, error) {
	if err != nil {
		return fail, err
	}
	return ok, nil
}

// @msg 接口返回的msg的内容
// @resrr 接口返回的error的内容
// @sysErr 程序运行时的产生的错误
func (l *LogicResult) Result(resrr *Error, sysErr error) (*Result, error) {

	return l.assert(SuccessNo(), Fail(resrr), sysErr)
}

// @data 接口返回的data的内容
// @resrr 接口返回的error的内容
// @sysErr 程序运行时的产生的错误
func (l *LogicResult) ResultData(data any, resrr *Error, sysErr error) (*Result, error) {

	return l.assert(SuccessData(data), Fail(resrr), sysErr)
}

func (l *LogicResult) Fail(resrr *Error, sysErr error) (*Result, error) {
	return Fail(resrr), sysErr
}

func (l *LogicResult) SuccessMsg(msg string) (*Result, error) {
	return SuccessMsg(msg), nil
}

func (l *LogicResult) SuccessData(data any) (*Result, error) {
	return SuccessData(data), nil
}

func (l *LogicResult) SuccessCode(code *Error) (*Result, error) {
	return SuccessCode(code), nil
}

// @返回格式为带分页数据
func (l *LogicResult) PageCode(code *Error) (*Result, error) {
	return SuccessCodeData(code, PageResult{}), nil
}
