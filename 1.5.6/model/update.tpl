func (m *default{{.upperStartCamelObject}}Model) Update(ctx context.Context, {{if .containsIndexCache}}newData{{else}}data{{end}} *{{.upperStartCamelObject}}) error {
	{{if .withCache}}{{if .containsIndexCache}}data, err:=m.FindOne(ctx, newData.{{.upperStartCamelPrimaryKey}})
	if err!=nil{
		return err
	}

{{end}}	{{.keys}}
    _, {{if .containsIndexCache}}err{{else}}err:{{end}}= m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where {{.originalPrimaryKey}} = {{if .postgreSql}}$1{{else}}?{{end}}", m.table, {{.lowerStartCamelObject}}RowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, {{.expressionValues}})
	}, {{.keyValues}}){{else}}query := fmt.Sprintf("update %s set %s where {{.originalPrimaryKey}} = {{if .postgreSql}}$1{{else}}?{{end}}", m.table, {{.lowerStartCamelObject}}RowsWithPlaceHolder)
    _,err:=m.conn.ExecCtx(ctx, query, {{.expressionValues}}){{end}}
	return err
}

//TODO 自定义方法 手动加上 _sqlQuery_
func (m *default{{.upperStartCamelObject}}Model) UpdateByBuild(ctx context.Context,builder sq.UpdateBuilder) (int64,error) {
    return _sqlQuery_.UpdateByBuild(ctx, builder)
}

func(m *default{{.upperStartCamelObject}}Model)  SelectBuilder() sq.SelectBuilder {
	return sq.Select().From(m.table)
}

func(m *default{{.upperStartCamelObject}}Model)  InsertBuilder() sq.InsertBuilder {
	return sq.Insert(m.table)
}

func(m *default{{.upperStartCamelObject}}Model)  UpdateBuilder() *{{.upperStartCamelObject}}UpdateBuilder {
	return NewBaseOptionUpdateBuilder(m.table)
}

func(m *default{{.upperStartCamelObject}}Model)  DeleteBuilder() sq.DeleteBuilder {
	return sq.Delete(m.table)
}