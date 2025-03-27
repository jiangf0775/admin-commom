InsertByBuild(ctx context.Context,builder sq.InsertBuilder) (int64,error)
DeleteByBuild(ctx context.Context,builder sq.UpdateBuilder) (int64,error)

Update(ctx context.Context, {{if .containsIndexCache}}newData{{else}}data{{end}} *{{.upperStartCamelObject}}) error
UpdateByBuild(ctx context.Context,builder sq.UpdateBuilder) (int64,error)
FindSum(ctx context.Context,builder sq.SelectBuilder, field string) (float64,error)
FindCount(ctx context.Context, builder sq.SelectBuilder, field string) (uint64,error)
FindAll(ctx context.Context,builder sq.SelectBuilder,orderBy string) ([]*{{.upperStartCamelObject}},error)
FindPageListByPage(ctx context.Context,builder sq.SelectBuilder,page ,pageSize uint64,orderBy string) ([]*{{.upperStartCamelObject}},error)
FindPageListByPageWithTotal(ctx context.Context,builder sq.SelectBuilder,page ,pageSize uint64,orderBy string) ([]*{{.upperStartCamelObject}},uint64,error)
FindPageListByIdDESC(ctx context.Context,builder sq.SelectBuilder ,preMinId ,pageSize uint64) ([]*{{.upperStartCamelObject}},error)
FindPageListByIdASC(ctx context.Context,builder sq.SelectBuilder,preMaxId ,pageSize uint64) ([]*{{.upperStartCamelObject}},error)
Trans(ctx context.Context,fn func(ctx context.Context,session sqlx.Session) error) error
SelectBuilder() sq.SelectBuilder
InsertBuilder() sq.InsertBuilder
UpdateBuilder() sq.UpdateBuilder
DeleteBuilder() sq.DeleteBuilder
