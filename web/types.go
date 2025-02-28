package web

type (
	BaseEntity struct {
		Id               int64  `json:"id"`
		CreatedDate      string `json:"createdDate"`
		CreatedUserName  int    `json:"createdUserName"`
		ModifiedDate     int    `json:"modifiedDate,optional"`
		ModifiedUserName int    `json:"modifiedUserName,optional"`
	}

	PageReq struct {
		Index int `json:"index,default=1"`
		Page  int `json:"page,default=10"`
	}

	PathId struct {
		Id int64 `path:"id"`
	}

	Page struct {
		Total   int `json:"total"`
		Results any `json:"results"`
	}
)
