package liteconv

import (
	"strconv"
)



//Float64ToString - Get string from float64
func Float64ToString(value float64) string {
	return strconv.FormatFloat(value, _FMT, _PREC, _BIT_SIZE_64)
}

/*//Float32ToString - Get string from float32
func Float32ToString(value float32) string {

}*/


//BoolToString - get string from bool
func BoolToString(value bool) string  {
	if value{
		return _VAL_TRUE_STR
	}
	return _VAL_FALSE_STR
}

/*
//Uint32ToString - get string from uint32
func Uint32ToString(value uint32) string {

}
*/

//Uint64ToString - get string from uint64
func Uint64ToString(value uint64) string {
	return strconv.FormatUint(value, _BASE_64)
}