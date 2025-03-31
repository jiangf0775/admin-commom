func (m *default{{.upperStartCamelObject}}Model) Delete(ctx context.Context, {{.lowerStartCamelPrimaryKey}} {{.dataType}}) error {
	{{if .withCache}}{{if .containsIndexCache}}data, err:=m.FindOne(ctx, {{.lowerStartCamelPrimaryKey}})
	if err!=nil{
		return err
	}

{{end}}	{{.keys}}
    _, err {{if .containsIndexCache}}={{else}}:={{end}} m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update from %s set `deleted`=? where {{.originalPrimaryKey}} = {{if .postgreSql}}$1{{else}}?{{end}}", DELETE_YES, m.table)
		return conn.ExecCtx(ctx, query, {{.lowerStartCamelPrimaryKey}})
	}, {{.keyValues}}){{else}}query := fmt.Sprintf("update from %s set `deleted`= ? where {{.originalPrimaryKey}} = {{if .postgreSql}}$1{{else}}?{{end}}", sqls.DELETE_YES, m.table)
		_,err:=m.conn.ExecCtx(ctx, query, {{.lowerStartCamelPrimaryKey}}){{end}}
	return err
}

func (m *default{{.upperStartCamelObject}}Model) DeleteByBuild(ctx context.Context,builder sq.UpdateBuilder) (int64,error) {
    return _sqlQuery_.DeleteByBuild(ctx, builder)
}