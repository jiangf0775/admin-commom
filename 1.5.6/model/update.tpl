func (m *default{{.upperStartCamelObject}}Model) Update(ctx context.Context, {{if .containsIndexCache}}newData{{else}}data{{end}} *{{.upperStartCamelObject}}) error {
	{{if .withCache}}{{if .containsIndexCache}}data, err:=m.FindOne(ctx, newData.{{.upperStartCamelPrimaryKey}})
	if err!=nil{
		return err
	}

{{end}}	{{.keys}}
    _, {{if .containsIndexCache}}err{{else}}err:{{end}}= m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		//TODO 1.仅需要更新的字段
		var setStr = "`field1`=?,`field2`=?,`field3`=?,`modified_date`=?, `modified_user_name`=?, `modified_user_id`=?"
		query := fmt.Sprintf("update %s set %s where {{.originalPrimaryKey}} = {{if .postgreSql}}$1{{else}}?{{end}}", m.table, setStr)
		//TODO 2.删去不需要跟新的字段值 ！！注意顺序
		return conn.ExecCtx(ctx, query, {{.expressionValues}})
	}, {{.keyValues}}){{else}}query := fmt.Sprintf("update %s set %s where {{.originalPrimaryKey}} = {{if .postgreSql}}$1{{else}}?{{end}}", m.table, {{.lowerStartCamelObject}}RowsWithPlaceHolder)
    _,err:=m.conn.ExecCtx(ctx, query, {{.expressionValues}}){{end}}
	return err
}

//TODO 自定义方法 手动加上 _sqlQuery_
func (m *default{{.upperStartCamelObject}}Model) UpdateByBuild(ctx context.Context,builder sq.UpdateBuilder) (int64,error) {
    return _sqlQuery_.UpdateByBuild(ctx, builder)
}

func(m *default{{.upperStartCamelObject}}Model)  SelectBuilder() {{.upperStartCamelObject}}SelectBuilder {
	return New{{.upperStartCamelObject}}SelectBuilder(m.table)
}

func(m *default{{.upperStartCamelObject}}Model)  UpdateBuilder() {{.upperStartCamelObject}}UpdateBuilder {
	return New{{.upperStartCamelObject}}UpdateBuilder(m.table)
}

func(m *default{{.upperStartCamelObject}}Model)  InsertBuilder() sq.InsertBuilder {
	return sq.Insert(m.table)
}

func(m *default{{.upperStartCamelObject}}Model)  DeleteBuilder() sq.DeleteBuilder {
	return sq.Delete(m.table)
}