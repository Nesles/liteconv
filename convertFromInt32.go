package liteconv

import "strconv"

//Int32ToString - Get string from int32
func Int32ToString(value int) string  {
	return 	strconv.Itoa(value)
}


//Int64ToString - Get string from int64
func Int64ToString(value int64) string {
	return strconv.FormatInt(value, _BASE_64)
}

func Int32ToBool(val int) bool  {
	return val!=_VAL_FALSE_INT
}

