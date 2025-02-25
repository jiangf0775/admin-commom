package sql

import (
	"context"
	"database/sql"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

// 数据库操作父类
type BaseModel[TEntity comparable] interface {
	Insert(ctx context.Context, data *TEntity) (sql.Result, error)
	FindOne(ctx context.Context, id int64) (*TEntity, error)
	Update(ctx context.Context, data *TEntity) error
	FindSum(ctx context.Context, builder SqlBuilder, field string) (float64, error)
	FindCount(ctx context.Context, builder SqlBuilder, field string) (int64, error)
	FindAll(ctx context.Context, builder SqlBuilder, orderBy string) ([]*TEntity, error)
	FindPageListByPage(ctx context.Context, builder SqlBuilder, page, pageSize int64, orderBy string) ([]*TEntity, error)
	FindPageListByPageWithTotal(ctx context.Context, builder SqlBuilder, page, pageSize int64, orderBy string) ([]*TEntity, int64, error)
	FindPageListByIdDESC(ctx context.Context, builder SqlBuilder, preMinId, pageSize int64) ([]*TEntity, error)
	FindPageListByIdASC(ctx context.Context, builder SqlBuilder, preMaxId, pageSize int64) ([]*TEntity, error)
	Trans(ctx context.Context, fn func(ctx context.Context, session sqlx.Session) error) error
	SelectBuilder() SqlBuilder
	Delete(ctx context.Context, id int64) error
}
