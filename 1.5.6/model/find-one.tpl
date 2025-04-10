func (m *default{{.upperStartCamelObject}}Model) FindOne(ctx context.Context, {{.lowerStartCamelPrimaryKey}} u{{.dataType}}) (*{{.upperStartCamelObject}}, error) {
	{{if .withCache}}{{.cacheKey}}
	var resp {{.upperStartCamelObject}}
	err := m.QueryRowCtx(ctx, &resp, {{.cacheKeyVariable}}, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query :=  fmt.Sprintf("select %s from %s where {{.originalPrimaryKey}} = {{if .postgreSql}}$1{{else}}?{{end}} limit 1", {{.lowerStartCamelObject}}Rows, m.table)
		return conn.QueryRowCtx(ctx, v, query, {{.lowerStartCamelPrimaryKey}})
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}{{else}}query := fmt.Sprintf("select %s from %s where {{.originalPrimaryKey}} = {{if .postgreSql}}$1{{else}}?{{end}} limit 1", {{.lowerStartCamelObject}}Rows, m.table)
	var resp {{.upperStartCamelObject}}
	err := m.conn.QueryRowCtx(ctx, &resp, query, {{.lowerStartCamelPrimaryKey}})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}{{end}}
}

//TODO 自定义方法 手动加上 _sqlQuery_
func (m *Query[TEntity]) FirstDefault(ctx context.Context, builder sq.SelectBuilder) (*TEntity, error){
    return _sqlQuery_.FirstDefault(ctx, builder)
}


func (m *default{{.upperStartCamelObject}}Model) FindSum(ctx context.Context,builder sq.SelectBuilder, field string) (float64,error) {
    return _sqlQuery_.FindSum(ctx, builder, field)
}


func (m *default{{.upperStartCamelObject}}Model) FindCount(ctx context.Context, builder sq.SelectBuilder, field string) (uint64,error) {
    return _sqlQuery_.FindCount(ctx, builder, field)
}

func (m *Query[TEntity]) FindCountDefault(ctx context.Context, builder sq.SelectBuilder) (uint64, error){
    return _sqlQuery_.FindCountDefault(ctx, builder)
}

func (m *default{{.upperStartCamelObject}}Model) FindAll(ctx context.Context,builder sq.SelectBuilder,orderBy string) ([]*{{.upperStartCamelObject}},error) {
    return _sqlQuery_.FindAll(ctx, builder, orderBy)
}

func (m *default{{.upperStartCamelObject}}Model) FindAllIdASC(ctx context.Context,builder sq.SelectBuilder) ([]*{{.upperStartCamelObject}},error) {
    return m.FindAll(ctx, builder, " id Asc")
}

func (m *default{{.upperStartCamelObject}}Model) FindAllIdASC(ctx context.Context,builder sq.SelectBuilder) ([]*{{.upperStartCamelObject}},error) {
    return m.FindAll(ctx, builder, " id Desc")
}

func (m *default{{.upperStartCamelObject}}Model) FindListByPage(ctx context.Context,builder sq.SelectBuilder,page ,pageSize uint64,orderBy string) ([]*{{.upperStartCamelObject}},error) {
    return _sqlQuery_.FindPageListByPage(ctx, builder, page, pageSize, orderBy)
}

func (m *default{{.upperStartCamelObject}}Model) FindListByPageIdASC(ctx context.Context,builder sq.SelectBuilder,preMaxId ,pageSize uint64) ([]*{{.upperStartCamelObject}},error)  {
    return _sqlQuery_.FindPageListByIdASC(ctx, builder, preMaxId, pageSize)
}

func (m *default{{.upperStartCamelObject}}Model) FindListByPageIdDESC(ctx context.Context,builder sq.SelectBuilder ,preMinId ,pageSize uint64) ([]*{{.upperStartCamelObject}},error) {
    return _sqlQuery_.FindPageListByIdDESC(ctx, builder, preMinId, pageSize)
}

func (m *default{{.upperStartCamelObject}}Model) FindListByPageWithTotal(ctx context.Context,builder sq.SelectBuilder,page ,pageSize uint64,orderBy string) ([]*{{.upperStartCamelObject}},uint64,error) {
    return _sqlQuery_.FindPageListByPageWithTotal(ctx, builder, page, pageSize, orderBy)
}

func (m *default{{.upperStartCamelObject}}Model) FindListByPageWithTotalIdASC(ctx context.Context,builder sq.SelectBuilder,page ,pageSize uint64) ([]*{{.upperStartCamelObject}},uint64,error) {
    return m.FindPageListByPageWithTotal(ctx, builder, page, pageSize, " id ASC")
}

func (m *default{{.upperStartCamelObject}}Model) FindListByPageWithTotalIdDESC(ctx context.Context,builder sq.SelectBuilder,page ,pageSize uint64) ([]*{{.upperStartCamelObject}},uint64,error) {
    return m.FindPageListByPageWithTotal(ctx, builder, page, pageSize, " id DESC")
}

func (m *default{{.upperStartCamelObject}}Model) Trans(ctx context.Context,fn func(ctx context.Context,session sqlx.Session) error) error {
    return _sqlQuery_.Trans(ctx, fn)
}