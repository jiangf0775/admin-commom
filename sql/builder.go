package sql

import (
	"database/sql"
	"github.com/Masterminds/squirrel"
)

// 封装【 squirrel.SelectBuilder 】
type SqlBuilder struct {
	builder squirrel.SelectBuilder
}

func (b SqlBuilder) Page(index, page int) SqlBuilder {
	if index < 1 {
		index = 1
	}
	offset := (index - 1) * page
	b.builder = b.builder.Offset(uint64(offset)).Limit(uint64(page))
	return b
}

func init() {

}

func (b SqlBuilder) Query() (*sql.Rows, error) {
	//squirrel.Or{}
	return b.builder.Query()
}

func (b SqlBuilder) QueryRow() squirrel.RowScanner {
	return b.builder.QueryRow()
}

func (b SqlBuilder) MustSql() (string, []interface{}) {
	return b.builder.MustSql()
}

func (b SqlBuilder) Prefix(sql string, args ...interface{}) SqlBuilder {
	b.builder = b.builder.Prefix(sql, args)
	return b
}

func (b SqlBuilder) PrefixExpr(expr squirrel.Sqlizer) SqlBuilder {
	b.builder = b.builder.PrefixExpr(expr)
	return b
}

func (b SqlBuilder) Distinct() SqlBuilder {
	b.builder = b.builder.Distinct()
	return b
}

func (b SqlBuilder) Options(options ...string) SqlBuilder {
	b.builder = b.builder.Options(options...)
	return b
}

// Columns adds result columns to the query.
func (b SqlBuilder) Columns(columns ...string) SqlBuilder {
	b.builder = b.builder.Columns(columns...)
	return b
}

func (b SqlBuilder) RemoveColumns() SqlBuilder {
	b.builder = b.builder.RemoveColumns()
	return b
}

func (b SqlBuilder) Column(column interface{}, args ...interface{}) SqlBuilder {
	b.builder = b.builder.Column(column, args...)
	return b
}

func (b SqlBuilder) From(from string) SqlBuilder {
	b.builder = b.builder.From(from)
	return b
}

func (b SqlBuilder) FromSelect(from SqlBuilder, alias string) SqlBuilder {
	b.builder = b.builder.FromSelect(from.builder, alias)
	return b
}

func (b SqlBuilder) JoinClause(pred interface{}, args ...interface{}) SqlBuilder {
	b.builder = b.builder.JoinClause(pred, args...)
	return b
}

func (b SqlBuilder) Join(join string, rest ...interface{}) SqlBuilder {
	b.builder = b.builder.Join(join, rest...)
	return b
}

func (b SqlBuilder) LeftJoin(join string, rest ...interface{}) SqlBuilder {
	b.builder = b.builder.LeftJoin(join, rest...)
	return b
}

func (b SqlBuilder) RightJoin(join string, rest ...interface{}) SqlBuilder {
	return b.RightJoin(join, rest...)
}

func (b SqlBuilder) InnerJoin(join string, rest ...interface{}) SqlBuilder {
	return b.InnerJoin(join, rest...)
}

func (b SqlBuilder) CrossJoin(join string, rest ...interface{}) SqlBuilder {
	return b.CrossJoin(join, rest...)
}

func (b SqlBuilder) Where(pred interface{}, args ...interface{}) SqlBuilder {
	b.builder = b.builder.Where(pred, args...)
	return b
}

func (b SqlBuilder) OrderBy(orderBys ...string) SqlBuilder {
	if orderBys == nil || len(orderBys) == 0 {
		b.builder = b.builder.OrderBy("id DESC")
	} else {
		b.builder = b.builder.OrderBy(orderBys...)
	}
	return b
}

func (b SqlBuilder) ToSql(orderBys ...string) (string, []interface{}, error) {
	return b.builder.ToSql()
}

func GetSqlBuilder(table string) SqlBuilder {
	selectB := squirrel.Select().From(table)
	sb := SqlBuilder{
		builder: selectB,
	}
	return sb
}

func (b SqlBuilder) Offset(offset uint64) SqlBuilder {
	b.builder = b.builder.Offset(offset)
	return b
}

func (b SqlBuilder) Limit(limit uint64) SqlBuilder {
	b.builder = b.builder.Limit(limit)
	return b
}
