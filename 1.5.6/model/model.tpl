package {{.pkg}}
{{if .withCache}}
import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)
{{else}}

import (
    sqls "common/sql"
    sq "github.com/Masterminds/squirrel"
    "github.com/zeromicro/go-zero/core/stores/sqlx"
)
{{end}}
var _ {{.upperStartCamelObject}}Model = (*custom{{.upperStartCamelObject}}Model)(nil)

type (
	// {{.upperStartCamelObject}}Model is an interface to be customized, add more methods here,
	// and implement the added methods in custom{{.upperStartCamelObject}}Model.
	{{.upperStartCamelObject}}Model interface {
		{{.lowerStartCamelObject}}Model
	}

	custom{{.upperStartCamelObject}}Model struct {
		*default{{.upperStartCamelObject}}Model
	}

    //扩展UpdateBuilder方法
    {{.upperStartCamelObject}}UpdateBuilder struct {
		sq.UpdateBuilder
		sqls.ModifyBuilder
	}
    //扩展SelectBuilder方法
    {{.upperStartCamelObject}}SelectBuilder struct {
		sq.SelectBuilder
		sqls.QueryBuilder
	}
   //todo 生成【 xxxModel.go 】

)

// New{{.upperStartCamelObject}}Model returns a model for the database table.
func New{{.upperStartCamelObject}}Model(conn sqlx.SqlConn{{if .withCache}}, c cache.CacheConf, opts ...cache.Option{{end}}) {{.upperStartCamelObject}}Model {
	return &custom{{.upperStartCamelObject}}Model{
		default{{.upperStartCamelObject}}Model: new{{.upperStartCamelObject}}Model(conn{{if .withCache}}, c, opts...{{end}}),
	}
}

//提供创建函数
func New{{.upperStartCamelObject}}UpdateBuilder(table string) {{.upperStartCamelObject}}UpdateBuilder {
	builder := {{.upperStartCamelObject}}UpdateBuilder{}
	builder.SetBuilder(sq.Update(table))
	return builder
}

//提供创建函数
func New{{.upperStartCamelObject}}SelectBuilder(table string) {{.upperStartCamelObject}}SelectBuilder {
	builder := {{.upperStartCamelObject}}SelectBuilder{}
	builder.SetBuilder(sq.Select().From(table))
	return builder
}

// #region {{.upperStartCamelObject}}的【 Getter Setter 】方法
func (m *{{.upperStartCamelObject}}) Getxxx() string {

	return ""
}
func (m *{{.upperStartCamelObject}}) Setxxx(value string) {

}
// #endregion