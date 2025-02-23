package sql

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type SqlQuery[TEntity comparable] struct {
	conn  sqlx.SqlConn
	rows  string
	table string
}

func (m *SqlQuery[TEntity]) FindSum(ctx context.Context, builder SqlBuilder, field string) (float64, error) {

	if len(field) == 0 {
		return 0, errors.Wrapf(errors.New("FindSum Least One Field"), "FindSum Least One Field")
	}

	builder = builder.Columns("IFNULL(SUM(" + field + "),0)")

	query, values, err := builder.Where("deleted = ?", DELETE_NO).ToSql()
	if err != nil {
		return 0, err
	}

	var resp float64

	err = m.conn.QueryRowCtx(ctx, &resp, query, values...)

	switch err {
	case nil:
		return resp, nil
	default:
		return 0, err
	}
}

func (m *SqlQuery[TEntity]) FindCount(ctx context.Context, builder SqlBuilder, field string) (int64, error) {

	return getCount(ctx, m.conn, builder, field)
}

func (m *SqlQuery[TEntity]) FindAll(ctx context.Context, builder SqlBuilder, orderBy string) ([]*TEntity, error) {

	builder = builder.Columns(m.rows)

	if orderBy == "" {
		builder = builder.OrderBy("id DESC")
	} else {
		builder = builder.OrderBy(orderBy)
	}

	query, values, err := builder.Where("deleted = ?", DELETE_NO).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*TEntity
	err = m.conn.QueryRowsCtx(ctx, &resp, query, values...)

	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *SqlQuery[TEntity]) FindPageListByPage(ctx context.Context, builder SqlBuilder, page, pageSize int64, orderBy string) ([]*TEntity, error) {

	builder = builder.Columns(m.rows)

	if orderBy == "" {
		builder = builder.OrderBy("id DESC")
	} else {
		builder = builder.OrderBy(orderBy)
	}

	if page < 1 {
		page = 1
	}
	offset := (page - 1) * pageSize

	query, values, err := builder.Where("deleted = ?", DELETE_NO).Offset(uint64(offset)).Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*TEntity

	err = m.conn.QueryRowsCtx(ctx, &resp, query, values...)

	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *SqlQuery[TEntity]) FindPageListByPageWithTotal(ctx context.Context, builder SqlBuilder, page, pageSize int64, orderBy string) ([]*TEntity, int64, error) {

	total, err := getCount(ctx, m.conn, builder, "id")
	if err != nil {
		return nil, 0, err
	}

	builder = builder.Columns(m.rows)

	if orderBy == "" {
		builder = builder.OrderBy("id DESC")
	} else {
		builder = builder.OrderBy(orderBy)
	}

	if page < 1 {
		page = 1
	}
	offset := (page - 1) * pageSize

	query, values, err := builder.Where("deleted = ?", DELETE_NO).Offset(uint64(offset)).Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, total, err
	}

	var resp []*TEntity

	err = m.conn.QueryRowsCtx(ctx, &resp, query, values...)

	switch err {
	case nil:
		return resp, total, nil
	default:
		return nil, total, err
	}
}

func (m *SqlQuery[TEntity]) FindPageListByIdDESC(ctx context.Context, builder SqlBuilder, preMinId, pageSize int64) ([]*TEntity, error) {

	builder = builder.Columns(m.rows)

	if preMinId > 0 {
		builder = builder.Where(" id < ? ", preMinId)
	}

	query, values, err := builder.Where("deleted = ?", DELETE_NO).OrderBy("id DESC").Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*TEntity

	err = m.conn.QueryRowsCtx(ctx, &resp, query, values...)

	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *SqlQuery[TEntity]) FindPageListByIdASC(ctx context.Context, builder SqlBuilder, preMaxId, pageSize int64) ([]*TEntity, error) {

	builder = builder.Columns(m.rows)

	if preMaxId > 0 {
		builder = builder.Where(" id > ? ", preMaxId)
	}

	query, values, err := builder.Where("deleted = ?", DELETE_NO).OrderBy("id ASC").Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*TEntity

	err = m.conn.QueryRowsCtx(ctx, &resp, query, values...)

	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *SqlQuery[TEntity]) Trans(ctx context.Context, fn func(ctx context.Context, session sqlx.Session) error) error {
	return m.conn.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		return fn(ctx, session)
	})
}

func getCount(ctx context.Context, conn sqlx.SqlConn, builder SqlBuilder, field string) (int64, error) {

	if len(field) == 0 {
		return 0, errors.Wrapf(errors.New("FindCount Least One Field"), "FindCount Least One Field")
	}

	builder = builder.Columns("COUNT(" + field + ")")

	query, values, err := builder.Where("deleted = ?", DELETE_NO).ToSql()
	if err != nil {
		return 0, err
	}

	var resp int64

	err = conn.QueryRowCtx(ctx, &resp, query, values...)

	switch err {
	case nil:
		return resp, nil
	default:
		return 0, err
	}
}
