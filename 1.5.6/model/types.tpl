type (
	{{.lowerStartCamelObject}}Model interface{
		{{.method}}
	}

	default{{.upperStartCamelObject}}Model struct {
		{{if .withCache}}sqlc.CachedConn{{else}}conn sqlx.SqlConn{{end}}
		table string
	}

	{{.upperStartCamelObject}} struct {
	    //数据库表映射的结构体
		{{.fields}}
		//TODO 手动删除与【 sqls.BaseModel 】 重复的字段
		sqls.BaseModel
	}
)
