package web

// 统一错误编码
var (
	DBError = NewError(1000, "数据库错误")
	SQError = NewError(1100, "SQ的Sql语法错误")
)
