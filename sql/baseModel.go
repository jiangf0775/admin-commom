package sql

import (
	"context"
	"database/sql"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"time"
)

var (
	BaseFields = []string{"`id`", "`deleted`", "`created_date`", "`created_user_name`", "`created_user_id`", "`modified_date`", "`modified_user_name`", "`modified_user_id`"}
)

type BaseModel struct {
	Id               uint64         `db:"id"`
	Deleted          uint64         `db:"deleted"`
	CreatedDate      time.Time      `db:"created_date"`
	CreatedUserName  string         `db:"created_user_name"`
	CreatedUserId    uint64         `db:"created_user_id"`
	ModifiedDate     sql.NullTime   `db:"modified_date"`
	ModifiedUserName sql.NullString `db:"modified_user_name"`
	ModifiedUserId   sql.NullInt64  `db:"modified_user_id"`
}

func (m *BaseModel) SetInsertField(name string, id uint64, date time.Time) {
	m.CreatedDate = date
	m.CreatedUserName = name
	m.CreatedUserId = id
}

func (m *BaseModel) SetEditField(name string, id int64, date time.Time) {
	m.ModifiedDate = sql.NullTime{Time: date, Valid: true}
	m.ModifiedUserName = sql.NullString{String: name, Valid: true}
	m.ModifiedUserId = sql.NullInt64{Int64: id, Valid: true}
}

func (m *BaseModel) GetEditField() (name string, date time.Time) {
	if m.ModifiedDate.Valid {
		date = m.ModifiedDate.Time
	}
	if m.ModifiedUserName.Valid {
		name = m.ModifiedUserName.String
	}
	return name, date
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
	FindListByPage(ctx context.Context, builder squirrel.SelectBuilder, page, pageSize uint64, orderBy string) ([]*TEntity, error)
	FindListByPageWithTotal(ctx context.Context, builder squirrel.SelectBuilder, page, pageSize uint64, orderBy string) ([]*TEntity, uint64, error)
	FindListByPageIdDESC(ctx context.Context, builder squirrel.SelectBuilder, preMinId, pageSize uint64) ([]*TEntity, error)
	FindListByPageIdASC(ctx context.Context, builder squirrel.SelectBuilder, preMaxId, pageSize uint64) ([]*TEntity, error)
}

type BaseInsert[TEntity comparable] interface {
	BaseOpt
	Insert(ctx context.Context, data TEntity) (sql.Result, error)
	InsertByBuild(ctx context.Context, builder squirrel.InsertBuilder) (int64, error)
	InsertBuilder() squirrel.InsertBuilder
}

type BaseUpdate[TEntity comparable] interface {
	BaseOpt
	Update(ctx context.Context, data TEntity) error
	UpdateByBuild(ctx context.Context, builder squirrel.UpdateBuilder) (int64, error)
	//UpdateBuilder() squirrel.UpdateBuilder
}

type BaseDelete interface {
	BaseOpt
	Delete(ctx context.Context, id uint64) error
	DeleteByBuild(ctx context.Context, builder squirrel.UpdateBuilder) (int64, error)
	DeleteBuilder() squirrel.DeleteBuilder
}
