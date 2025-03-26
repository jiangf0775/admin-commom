package web

type LogicResult struct {
}

func (l *LogicResult) assert(ok *Result, fail *Result, err error) (*Result, error) {
	if err != nil {
		return fail, err
	}
	return ok, nil
}

// @msg 接口返回的msg的内容
// @resrr 接口返回的error的内容
// @sysErr 程序运行时的产生的错误
func (l *LogicResult) ResultMsg(msg string, resrr *Error, sysErr error) (resp *Result, err error) {

	return l.assert(SuccessMsg(msg), Fail(resrr), sysErr)
}

// @data 接口返回的data的内容
// @resrr 接口返回的error的内容
// @sysErr 程序运行时的产生的错误
func (l *LogicResult) ResultData(data any, resrr *Error, sysErr error) (resp *Result, err error) {

	return l.assert(SuccessData(data), Fail(resrr), sysErr)
}
