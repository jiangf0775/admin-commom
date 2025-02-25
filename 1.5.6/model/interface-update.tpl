Update(ctx context.Context, {{if .containsIndexCache}}newData{{else}}data{{end}} *{{.upperStartCamelObject}}) error
FindSum(ctx context.Context,builder sqls.SqlBuilder, field string) (float64,error)
FindCount(ctx context.Context, builder sqls.SqlBuilder, field string) (int64,error)
FindAll(ctx context.Context,builder sqls.SqlBuilder,orderBy string) ([]*{{.upperStartCamelObject}},error)
FindPageListByPage(ctx context.Context,builder sqls.SqlBuilder,page ,pageSize int64,orderBy string) ([]*{{.upperStartCamelObject}},error)
FindPageListByPageWithTotal(ctx context.Context,builder sqls.SqlBuilder,page ,pageSize int64,orderBy string) ([]*{{.upperStartCamelObject}},int64,error)
FindPageListByIdDESC(ctx context.Context,builder sqls.SqlBuilder ,preMinId ,pageSize int64) ([]*{{.upperStartCamelObject}},error)
FindPageListByIdASC(ctx context.Context,builder sqls.SqlBuilder,preMaxId ,pageSize int64) ([]*{{.upperStartCamelObject}},error)
Trans(ctx context.Context,fn func(ctx context.Context,session sqlx.Session) error) error
SqlBuilder() sqls.SqlBuilder