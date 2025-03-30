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
	    //***所有字段中可空类型[ sql.Nullxxx ]的统一使用【get set】且字段名改为小写
	    //以减少取值赋值时的if判断【get set】在 {{.upperStartCamelObject}}Model.go中增加
		{{.fields}}
	}
)