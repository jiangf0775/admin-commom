package web

type DataSourceRequest struct {
	Page     int                 `json:"page"`
	PageSize int                 `json:"pageSize"`
	Sorts    string              `json:"sorts"`
	Filters  []FilterDescription `json:"filters"`
}

/*
*

	  Or{ "deleted_at": nil }  WHERE deleted_at IS NULL

	  Or{
		  Eq{"col1": 1, "col2": 2},
		  Eq{"col1": 3, "col2": 4}
		}
	   WHERE (col1 = 1 AND col2 = 2) OR (col1 = 3 AND col2 = 4)
*/
type FilterDescription struct {
	Opt   string      `json:"opt,options=or|and"`
	Descr Description `json:"descr"`
}

type Description struct {
	Field string `json:"field"`
	value any    `json:"value"`
}
