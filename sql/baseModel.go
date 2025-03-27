package sql

import (
	"context"
	"database/sql"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"time"
)

type CreateModel interface {
	SetInsert(name string, id uint64, date time.Time)
}

type UpdateModel interface {
	SetEdit(name string, id uint64, date time.Time)
}

type BaseOpt interface {
	Trans(ctx context.Context, fn func(ctx context.Context, session sqlx.Session) error) error
}

// 数据库操作父类
type BaseQuery[TEntity comparable] interface {
	BaseOpt
	FindOne(ctx context.Context, id uint64) (*TEntity, error)
	FindSum(ctx context.Context, builder squirrel.SelectBuilder, field string) (float64, error)
	FindCount(ctx context.Context, builder squirrel.SelectBuilder, field string) (uint64, error)
	FindAll(ctx context.Context, builder squirrel.SelectBuilder, orderBy string) ([]*TEntity, error)
	FindPageListByPage(ctx context.Context, builder squirrel.SelectBuilder, page, pageSize uint64, orderBy string) ([]*TEntity, error)
	FindPageListByPageWithTotal(ctx context.Context, builder squirrel.SelectBuilder, page, pageSize uint64, orderBy string) ([]*TEntity, uint64, error)
	FindPageListByIdDESC(ctx context.Context, builder squirrel.SelectBuilder, preMinId, pageSize uint64) ([]*TEntity, error)
	FindPageListByIdASC(ctx context.Context, builder squirrel.SelectBuilder, preMaxId, pageSize uint64) ([]*TEntity, error)
	SelectBuilder() squirrel.SelectBuilder
}

type BaseInsert[TEntity comparable] interface {
	BaseOpt
	Insert(ctx context.Context, data *TEntity) (sql.Result, error)
	InsertByBuild(ctx context.Context, builder squirrel.InsertBuilder) (int64, error)
	InsertBuilder() squirrel.InsertBuilder
}

type BaseUpdate[TEntity comparable] interface {
	BaseOpt
	Update(ctx context.Context, data *TEntity) error
	UpdateByBuild(ctx context.Context, builder squirrel.UpdateBuilder) (int64, error)
	UpdateBuilder() squirrel.UpdateBuilder
}

type BaseDelete interface {
	BaseOpt
	Delete(ctx context.Context, id uint64) error
	DeleteByBuild(ctx context.Context, builder squirrel.UpdateBuilder) (int64, error)
	DeleteBuilder() squirrel.DeleteBuilder
}
