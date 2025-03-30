package sql

import (
	sq "github.com/Masterminds/squirrel"
)

/**
 * 用于扩展公共方法
 */

type QueryBuilder struct {
	builder sq.SelectBuilder
}

func (m *QueryBuilder) SetBuilder(builder sq.SelectBuilder) QueryBuilder {
	m.builder = builder
	return QueryBuilder{builder: builder}
}

func (m QueryBuilder) GetBuilder() sq.SelectBuilder {
	return m.builder
}

func (m QueryBuilder) SkipDel() QueryBuilder {
	m.builder = m.builder.Where(" deleted = ?", DELETE_NO)
	return m
}

func (m QueryBuilder) BaseColumns() QueryBuilder {
	m.builder = m.builder.Columns(BaseFields...)
	return QueryBuilder{builder: m.builder}
}
