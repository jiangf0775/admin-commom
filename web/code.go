package web

const OK = 0

// 统一错误编码
var (
	DBError = NewError(1000, "数据库错误")
)
