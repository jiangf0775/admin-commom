package tool

import (
	"database/sql"
)

/*
 *  不直接返回
 */
type (
	InsertResult struct {
		//最后写入的id
		LastId uint64 `json:"lastId"`

		//受影响行数
		Rows uint64 `json:"rows"`
	}

	UpdateResult struct {
		//受影响行数
		Rows uint64 `json:"rows"`
	}

	DeleteResult struct {
		//受影响行数
		Rows uint64 `json:"rows"`
	}
)

func SqlResultToInsertResult(value sql.Result) InsertResult {
	var result InsertResult
	//mysql的resut.go文件下的 LastInsertId 方法err永远都是nil
	id, _ := value.LastInsertId()
	result.LastId = uint64(id)

	//mysql的resut.go文件下的 RowsAffected 方法err永远都是nil
	affected, _ := value.RowsAffected()
	result.Rows = uint64(affected)
	return result
}

func SqlResultToUpdateResult(value sql.Result) UpdateResult {
	var result UpdateResult
	affected, _ := value.RowsAffected()
	result.Rows = uint64(affected)
	return result
}

func SqlResultToDeleteResult(value sql.Result) DeleteResult {
	var result DeleteResult
	affected, _ := value.RowsAffected()
	result.Rows = uint64(affected)
	return result
}
