package tool

import "strconv"

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

// #region string
func StringToUint64s(value []string) []uint64 {
	if len(value) == 0 {
		return nil
	}
	var result []uint64 = make([]uint64, len(value))
	for i, v := range value {
		v2, _ := strconv.ParseUint(v, 10, 64)
		result[i] = v2
	}
	return result
}

//#endregion

// #region uint64
func Uint64ToStrings(value []uint64) []string {
	if len(value) == 0 {
		return nil
	}
	var result []string = make([]string, len(value))
	for i, v := range value {
		result[i] = strconv.FormatUint(v, 10)
	}
	return result
}

//#endregion
