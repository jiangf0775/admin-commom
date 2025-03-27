package web

// 统一错误编码
var (
	DBError       = NewError(1100, "数据库错误")
	DBQueryError  = NewError(1101, "数据查询时发生错误")
	DBInsertError = NewError(1102, "数据写入时发生错误")
	DBUpdateError = NewError(1103, "数据更新时发生错误")
	DBDeleteError = NewError(1104, "数据删除时发生错误")

	ServerInternalError = NewError(1200, "服务内部错误")
	BizError            = NewError(1300, "业务错误")
)
