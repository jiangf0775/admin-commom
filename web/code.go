package web

// 统一错误编码
var (
	DBError             = NewError(1000, "数据库错误")
	InternalServerError = NewError(1001, "服务内部错误错误")
)
