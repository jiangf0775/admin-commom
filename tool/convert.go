package tool

// #region 1.interface
func Interface2Uint64s(value []uint64) []interface{} {
	if len(value) == 0 {
		return []interface{}{}
	}
	result := make([]interface{}, len(value))
	for i, v := range value {
		result[i] = v
	}
	return result
}

//#endregion

// #region 2.uint64
func Uint64ToInterfaces(value []uint64) []interface{} {
	if len(value) == 0 {
		return []interface{}{}
	}
	result := make([]interface{}, len(value))
	for i, v := range value {
		result[i] = v
	}
	return result
}

//#endregion
