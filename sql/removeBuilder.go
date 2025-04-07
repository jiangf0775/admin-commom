package sql

import (
	sq "github.com/Masterminds/squirrel"
)

/**
 * 用于扩展公共方法
 */

type RemoveBuilder struct {
	builder sq.UpdateBuilder
}

func (m *RemoveBuilder) SetBuilder(builder sq.UpdateBuilder) RemoveBuilder {
	m.builder = builder
	return RemoveBuilder{builder: builder}
}
func (m RemoveBuilder) GetBuilder() sq.UpdateBuilder {
	return m.builder
}
