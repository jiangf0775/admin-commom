package web

// 统一错误编码
var (
	DBQueryError  = NewError(1001, "数据查询时发生错误")
	DBInsertError = NewError(1002, "数据写入时发生错误")
	DBUpdateError = NewError(1003, "数据更新时发生错误")
	DBDeleteError = NewError(1004, "数据删除时发生错误")

	NotifyNotFundData = NewError(OK, "没有获取到数据")
	NotifySuccess     = NewError(OK, "success")

	BizError = NewError(1200, "业务错误")

	ValitateRangeError = NewError(FAIL, "数据校验数不通过，字范围设置错误")   // wrong number range setting
	ValitateTypeError  = NewError(FAIL, "数据校验数不通过，字段值类型设置错误") // unsupported type on setting field value

	JwtGetUserInfoError         = NewError(FAIL, "用户信息获取失败")
	JwtAuthenticationError      = NewError(FAIL, "token认证失败")
	RequestArgumentConvertError = NewError(FAIL, "请求参数类型转换错误")
	DefaultError                = NewError(FAIL, "执行时发生错误")
)
