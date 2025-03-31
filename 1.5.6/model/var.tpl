var (
//TODO 定义仅 {{.lowerStartCamelObject}} 表所需字段
{{.lowerStartCamelObject}}Fiedls = []string{"`字段名1`", "`字段名2`", "`字段名3`"}
{{.lowerStartCamelObject}}FieldNames = append({{.lowerStartCamelObject}}Fiedls, sqls.BaseFields...)
{{.lowerStartCamelObject}}Rows = strings.Join({{.lowerStartCamelObject}}FieldNames, ",")
{{.lowerStartCamelObject}}RowsExpectAutoSet = {{if .postgreSql}}strings.Join(stringx.Remove({{.lowerStartCamelObject}}FieldNames, {{if .autoIncrement}}"{{.originalPrimaryKey}}", {{end}} {{.ignoreColumns}}), ","){{else}}strings.Join(stringx.Remove({{.lowerStartCamelObject}}FieldNames, {{if .autoIncrement}}"{{.originalPrimaryKey}}", {{end}} {{.ignoreColumns}}), ","){{end}}
{{.lowerStartCamelObject}}RowsWithPlaceHolder = {{if .postgreSql}}builder.PostgreSqlJoin(stringx.Remove({{.lowerStartCamelObject}}FieldNames, "{{.originalPrimaryKey}}", {{.ignoreColumns}})){{else}}strings.Join(stringx.Remove({{.lowerStartCamelObject}}FieldNames, "{{.originalPrimaryKey}}", {{.ignoreColumns}}), "=?,") + "=?"{{end}}

{{if .withCache}}{{.cacheKeys}}{{end}}

{{.lowerStartCamelObject}}Query = sqls.Query[{{.upperStartCamelObject}}]{ Rows:{{.lowerStartCamelObject}}Rows }
)
