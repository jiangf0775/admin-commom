type (
	{{.lowerStartCamelObject}}Model interface{
		{{.method}}
	}

	default{{.upperStartCamelObject}}Model struct {
		{{if .withCache}}sqlc.CachedConn{{else}}conn sqlx.SqlConn{{end}}
		table string
	}

	{{.upperStartCamelObject}} struct {
	    //数据库表映射的结构体
		{{.fields}}
	}
)


func (m *{{.upperStartCamelObject}} ) SetInsert(name string, id uint64, date time.Time) {
	m.CreatedDate = date
	m.CreatedUserName = name
	m.CreatedUserId = id

}

func (m *{{.upperStartCamelObject}} ) SetEdit(name string, id int64, date time.Time) {
	m.ModifiedDate = sql.NullTime{Time: date, Valid: true}
	m.ModifiedUserName = sql.NullString{String: name, Valid: true}
	m.ModifiedUserId = sql.NullInt64{Int64: id, Valid: true}
}