



func new{{.upperStartCamelObject}}Model(conn sqlx.SqlConn{{if .withCache}}, c cache.CacheConf, opts ...cache.Option{{end}}) *default{{.upperStartCamelObject}}Model {
    //该方法在xxxxModel.go文件中被调用
    //TODO 将文件中 _sqlQuery_ 全部替换为 {{.upperStartCamelObject}}Query
     _sqlQuery_.Conn = conn
     _sqlQuery_.Table = {{.table}}

	return &default{{.upperStartCamelObject}}Model{
		{{if .withCache}}CachedConn: sqlc.NewConn(conn, c, opts...){{else}}conn:conn{{end}},
		table:      {{.table}},
	}
}

func (m *default{{.upperStartCamelObject}}Model) withSession(session sqlx.Session) *default{{.upperStartCamelObject}}Model {
	return &default{{.upperStartCamelObject}}Model{
		{{if .withCache}}CachedConn:m.CachedConn.WithSession(session){{else}}conn:sqlx.NewSqlConnFromSession(session){{end}},
		table:      {{.table}},
	}
}
