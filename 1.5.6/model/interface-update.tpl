Update(ctx context.Context, {{if .containsIndexCache}}newData{{else}}data{{end}} *{{.upperStartCamelObject}}) error
UpdateByField(ctx context.Context,clauses map[string]interface{}) error
UpdateByWhere(ctx context.Context,builder sq.UpdateBuilder) error
FindSum(ctx context.Context,builder sq.SelectBuilder, field string) (float64,error)
FindCount(ctx context.Context, builder sq.SelectBuilder, field string) (int64,error)
FindAll(ctx context.Context,builder sq.SelectBuilder,orderBy string) ([]*{{.upperStartCamelObject}},error)
FindPageListByPage(ctx context.Context,builder sq.SelectBuilder,page ,pageSize int,orderBy string) ([]*{{.upperStartCamelObject}},error)
FindPageListByPageWithTotal(ctx context.Context,builder sq.SelectBuilder,page ,pageSize int,orderBy string) ([]*{{.upperStartCamelObject}},int64,error)
FindPageListByIdDESC(ctx context.Context,builder sq.SelectBuilder ,preMinId ,pageSize int) ([]*{{.upperStartCamelObject}},error)
FindPageListByIdASC(ctx context.Context,builder sq.SelectBuilder,preMaxId ,pageSize int) ([]*{{.upperStartCamelObject}},error)
Trans(ctx context.Context,fn func(ctx context.Context,session sqlx.Session) error) error
SqlBuilder() sq.SelectBuilder
UpdateBuilder() sq.UpdateBuilder
