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
	    //***所有字段中可空类型[ sql.Nullxxx ]的统一使用【get set】
	    //【get set】在 {{.upperStartCamelObject}}Model.go 文件中增加
		{{.fields}}
	}
)