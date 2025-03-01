package web

import "time"

type (
	BaseEntity struct {
		Id               int64     `json:"id"`
		CreatedDate      time.Time `json:"createdDate"`
		CreatedUserName  string    `json:"createdUserName"`
		ModifiedDate     time.Time `json:"modifiedDate,optional"`
		ModifiedUserName string    `json:"modifiedUserName,optional"`
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
