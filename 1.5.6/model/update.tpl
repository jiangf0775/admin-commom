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
func (m *default{{.upperStartCamelObject}}Model) UpdateByField(ctx context.Context,clauses sqls.BaseModelFieldMap) (int64,error) {
    return _sqlQuery_.UpdateByField(ctx, clauses)
}
func (m *default{{.upperStartCamelObject}}Model) UpdateByWhere(ctx context.Context,builder sq.UpdateBuilder) (int64,error) {
    return _sqlQuery_.UpdateByWhere(ctx, builder)
}


func (m *default{{.upperStartCamelObject}}Model) FindSum(ctx context.Context,builder sq.SelectBuilder, field string) (float64,error) {
    return _sqlQuery_.FindSum(ctx, builder, field)
}


func (m *default{{.upperStartCamelObject}}Model) FindCount(ctx context.Context, builder sq.SelectBuilder, field string) (int64,error) {
    return _sqlQuery_.FindCount(ctx, builder, field)
}

func (m *default{{.upperStartCamelObject}}Model) FindAll(ctx context.Context,builder sq.SelectBuilder,orderBy string) ([]*{{.upperStartCamelObject}},error) {
    return _sqlQuery_.FindAll(ctx, builder, orderBy)
}

func (m *default{{.upperStartCamelObject}}Model) FindPageListByPage(ctx context.Context,builder sq.SelectBuilder,page ,pageSize int,orderBy string) ([]*{{.upperStartCamelObject}},error) {
    return _sqlQuery_.FindPageListByPage(ctx, builder, page, pageSize, orderBy)
}

func (m *default{{.upperStartCamelObject}}Model) FindPageListByPageWithTotal(ctx context.Context,builder sq.SelectBuilder,page ,pageSize int,orderBy string) ([]*{{.upperStartCamelObject}},int64,error) {
    return _sqlQuery_.FindPageListByPageWithTotal(ctx, builder, page, pageSize, orderBy)
}

func (m *default{{.upperStartCamelObject}}Model) FindPageListByIdDESC(ctx context.Context,builder sq.SelectBuilder ,preMinId ,pageSize int) ([]*{{.upperStartCamelObject}},error) {
    return _sqlQuery_.FindPageListByIdDESC(ctx, builder, preMinId, pageSize)
}

func (m *default{{.upperStartCamelObject}}Model) FindPageListByIdASC(ctx context.Context,builder sq.SelectBuilder,preMaxId ,pageSize int) ([]*{{.upperStartCamelObject}},error)  {
    return _sqlQuery_.FindPageListByIdASC(ctx, builder, preMaxId, pageSize)
}

func (m *default{{.upperStartCamelObject}}Model) Trans(ctx context.Context,fn func(ctx context.Context,session sqlx.Session) error) error {
    return _sqlQuery_.Trans(ctx, fn)
}

func(m *default{{.upperStartCamelObject}}Model)  SqlBuilder() sq.SelectBuilder {
	return sq.Select().From(m.table)
}
func(m *default{{.upperStartCamelObject}}Model)  UpdateBuilder() sq.UpdateBuilder {
	return sq.Update(m.table)
}
