package sql

import (
	sq "github.com/Masterminds/squirrel"
	"time"
)

/**
 * 用于扩展公共方法
 */

type ModifyBuilder struct {
	builder sq.UpdateBuilder
}

func (m *ModifyBuilder) SetBuilder(builder sq.UpdateBuilder) ModifyBuilder {
	m.builder = builder
	return ModifyBuilder{builder: builder}
}
func (m ModifyBuilder) GetBuilder() sq.UpdateBuilder {
	return m.builder
}

func (m ModifyBuilder) SetEditField(name string, id uint64, date time.Time) ModifyBuilder {
	fields := make(map[string]interface{})
	fields["modified_user_id"] = id
	fields["modified_date"] = date
	fields["modified_user_name"] = name
	m.builder = m.builder.SetMap(fields)
	return ModifyBuilder{builder: m.builder}
}
