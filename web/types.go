package web

type (
	BaseEntity struct {
		Id               uint64 `json:"id"`
		CreatedDate      uint64 `json:"createdDate"`
		CreatedUserName  string `json:"createdUserName"`
		ModifiedDate     uint64 `json:"modifiedDate,optional"`
		ModifiedUserName string `json:"modifiedUserName,optional"`
	}

	PageReq struct {
		Index int64 `json:"index,default=1"`
		Page  int   `json:"page,default=10"`
	}

	PathId struct {
		Id int64 `path:"id"`
	}

	PageResult struct {
		Total   int64 `json:"total"`
		Results any   `json:"results"`
	}
)
