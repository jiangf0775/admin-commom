package sql

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type Query[TEntity comparable] struct {
	Conn  sqlx.SqlConn
	Rows  string
	Table string
}

func (m *Query[TEntity]) FindSum(ctx context.Context, builder sq.SelectBuilder, field string) (float64, error) {

	if len(field) == 0 {
		return 0, errors.Wrapf(errors.New("FindSum Least One Field"), "FindSum Least One Field")
	}

	builder = builder.Columns("IFNULL(SUM(" + field + "),0)")

	query, values, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	var resp float64

	err = m.Conn.QueryRowCtx(ctx, &resp, query, values...)

	switch err {
	case nil:
		return resp, nil
	default:
		return 0, err
	}
}

func (m *Query[TEntity]) FindCount(ctx context.Context, builder sq.SelectBuilder, field string) (uint64, error) {

	return getCount(ctx, m.Conn, builder, field)
}

func (m *Query[TEntity]) FindAll(ctx context.Context, builder sq.SelectBuilder, orderBy string) ([]*TEntity, error) {

	builder = builder.Columns(m.Rows)

	if orderBy == "" {
		builder = builder.OrderBy("id DESC")
	} else {
		builder = builder.OrderBy(orderBy)
	}

	query, values, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*TEntity
	err = m.Conn.QueryRowsCtx(ctx, &resp, query, values...)

	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *Query[TEntity]) FindPageListByPage(ctx context.Context, builder sq.SelectBuilder, page, pageSize uint64, orderBy string) ([]*TEntity, error) {

	builder = builder.Columns(m.Rows)

	if orderBy == "" {
		builder = builder.OrderBy("id DESC")
	} else {
		builder = builder.OrderBy(orderBy)
	}

	if page < 1 {
		page = 1
	}
	offset := (page - 1) * pageSize

	query, values, err := builder.Offset(uint64(offset)).Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*TEntity

	err = m.Conn.QueryRowsCtx(ctx, &resp, query, values...)

	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *Query[TEntity]) FindPageListByPageWithTotal(ctx context.Context, builder sq.SelectBuilder, page, pageSize uint64, orderBy string) ([]*TEntity, uint64, error) {

	total, err := getCount(ctx, m.Conn, builder, "id")
	if err != nil {
		return nil, 0, err
	}

	builder = builder.Columns(m.Rows)

	if orderBy == "" {
		builder = builder.OrderBy("id DESC")
	} else {
		builder = builder.OrderBy(orderBy)
	}

	if page < 1 {
		page = 1
	}
	offset := (page - 1) * pageSize

	query, values, err := builder.Offset(uint64(offset)).Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, total, err
	}

	var resp []*TEntity

	err = m.Conn.QueryRowsCtx(ctx, &resp, query, values...)

	switch err {
	case nil:
		return resp, total, nil
	default:
		return nil, total, err
	}
}

func (m *Query[TEntity]) FindPageListByIdDESC(ctx context.Context, builder sq.SelectBuilder, preMinId, pageSize uint64) ([]*TEntity, error) {

	builder = builder.Columns(m.Rows)

	if preMinId > 0 {
		builder = builder.Where(" id < ? ", preMinId)
	}

	query, values, err := builder.OrderBy("id DESC").Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*TEntity

	err = m.Conn.QueryRowsCtx(ctx, &resp, query, values...)

	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *Query[TEntity]) FindPageListByIdASC(ctx context.Context, builder sq.SelectBuilder, preMaxId, pageSize uint64) ([]*TEntity, error) {

	builder = builder.Columns(m.Rows)

	if preMaxId > 0 {
		builder = builder.Where(" id > ? ", preMaxId)
	}

	query, values, err := builder.OrderBy("id ASC").Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*TEntity

	err = m.Conn.QueryRowsCtx(ctx, &resp, query, values...)

	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *Query[TEntity]) Trans(ctx context.Context, fn func(ctx context.Context, session sqlx.Session) error) error {
	return m.Conn.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		return fn(ctx, session)
	})
}

func getCount(ctx context.Context, conn sqlx.SqlConn, builder sq.SelectBuilder, field string) (uint64, error) {

	if len(field) == 0 {
		return 0, errors.Wrapf(errors.New("FindCount Least One Field"), "FindCount Least One Field")
	}

	/*builder = builder.Columns("COUNT(" + field + ")")

	query, values, err := builder.Where("deleted = ?", DELETE_NO).ToSql()
	if err != nil {
		return 0, err
	}*/

	query := "SELECT COUNT(id) FROM `base_option` WHERE deleted = ?"
	var values = make([]interface{}, 1)
	values[0] = 0

	var resp int64
	err := conn.QueryRowCtx(ctx, &resp, query, values...)

	switch err {
	case nil:
		return uint64(resp), nil
	default:
		return 0, err
	}
}

func (m *Query[TEntity]) UpdateByBuild(ctx context.Context, builder sq.UpdateBuilder) (int64, error) {

	sql, values, err := builder.ToSql()
	if err != nil {
		return 0, err
	}
	result, err := m.Conn.ExecCtx(ctx, sql, values...)
	if err != nil {
		return 0, err
	}
	row, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return row, nil
}

func (m *Query[TEntity]) InsertByBuild(ctx context.Context, builder sq.InsertBuilder) (int64, error) {

	sql, _, err := builder.ToSql()
	if err != nil {
		return 0, err
	}
	result, err := m.Conn.ExecCtx(ctx, sql)
	if err != nil {
		return 0, err
	}
	row, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return row, nil
}

func (m *Query[TEntity]) DeleteByBuild(ctx context.Context, builder sq.UpdateBuilder) (int64, error) {

	sql, _, err := builder.ToSql()
	if err != nil {
		return 0, err
	}
	result, err := m.Conn.ExecCtx(ctx, sql)
	if err != nil {
		return 0, err
	}
	row, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return row, nil
}
