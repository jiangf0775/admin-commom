func (m *default{{.upperStartCamelObject}}Model) Delete(ctx context.Context, data {{.upperStartCamelObject}}) error {
	{{if .withCache}}{{if .containsIndexCache}}data, err:=m.FindOne(ctx, {{.lowerStartCamelPrimaryKey}})
	if err!=nil{
		return err
	}

{{end}}	{{.keys}}
    _, err {{if .containsIndexCache}}={{else}}:={{end}} m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update from %s set `deleted`=? where {{.originalPrimaryKey}} = {{if .postgreSql}}$1{{else}}?{{end}}", m.table)
		return conn.ExecCtx(ctx, query, {{.lowerStartCamelPrimaryKey}})
	}, {{.keyValues}}){{else}}query := fmt.Sprintf("update from %s set `deleted`= ?,`modified_date`=?, `modified_user_name`=?, `modified_user_id`=? where {{.originalPrimaryKey}} = {{if .postgreSql}}$1{{else}}?{{end}}", m.table)
		_,err:=m.conn.ExecCtx(ctx, query, sqls.DELETE_YES, data.ModifiedDate, data.ModifiedUserName, data.ModifiedUserId,{{.lowerStartCamelPrimaryKey}}){{end}}
	return err
}

func (m *default{{.upperStartCamelObject}}Model) DeleteByBuild(ctx context.Context,builder sq.UpdateBuilder) (int64,error) {
    return _sqlQuery_.DeleteByBuild(ctx, builder)
}