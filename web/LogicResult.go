package web

type LogicResult struct{}

func (l *LogicResult) assert(ok *Result, fail *Result, err error) *Result {
	if err == nil {
		return ok
	}
	return fail
}

// @msg 接口返回的msg的内容
// @resErr 如果sysErr错误时需要返回的错误
// @sysErr 程序运行时的产生的错误
func (l *LogicResult) ResultIfErr(resErr *Error, sysErr error) *Result {

	return l.assert(SuccessNo(), Fail(resErr), sysErr)
}

// @data 接口返回的data的内容
// @resrr 接口返回的error的内容
// @sysErr 程序运行时的产生的错误
func (l *LogicResult) ResultData(data any, resErr *Error, sysErr error) *Result {

	return l.assert(SuccessData(data), Fail(resErr), sysErr)
}

func (l *LogicResult) Fail(resrr *Error) *Result {
	return Fail(resrr)
}

func (l *LogicResult) SuccessMsg(msg string) *Result {
	return SuccessMsg(msg)
}

func (l *LogicResult) SuccessData(data any) *Result {
	return SuccessData(data)
}

func (l *LogicResult) SuccessCode(code *Error) *Result {
	return SuccessCode(code)
}

// @返回格式为带分页数据
func (l *LogicResult) PageCode(code *Error) *Result {
	return SuccessCodeData(code, PageResult{})
}

func (l *LogicResult) PageData(data any, total uint64) *Result {
	return SuccessCodeData(NotifySuccess, PageResult{Total: total, Results: data})
}
