package liteconv

import (
	"strconv"
//	"Kitana/dbcon"
)

const _FMT = 'f'

//Float64ToString - Get string from float64
func Float64ToString(value float64) string {
	return strconv.FormatFloat(value, _FMT, 0, 64)
}

/*//Float32ToString - Get string from float32
func Float32ToString(value float32) string {

}*/

//Int32ToString - Get string from int32
func Int32ToString(value int) string  {
	return 	strconv.Itoa(value)
}

//Int64ToString - Get string from int64
func Int64ToString(value int64) string {
	return strconv.FormatInt(value, 64)
}

//BoolToString - get string from bool
func BoolToString(value bool) string  {
	if value{
		return "true"
	}
	return "false"
}

/*
//Uint32ToString - get string from uint32
func Uint32ToString(value uint32) string {

}
*/

//Uint64ToString - get string from uint64
func Uint64ToString(value uint64) string {
	return strconv.FormatUint(value, 64)
}

// ToStringMap interface{} to map[string]string
func InterfaceToStringMap(i1 interface{}) map[string]string {
	switch i2 := i1.(type) {
	case map[string]interface{}:
		m1 := map[string]string{}
		for k, v := range i2 {
			m1[k] = InterfaceToString(v)
		}
		return m1
	case map[string]string:
		return i2
	default:
		return map[string]string{}
	}
}


//InterfaceToString - get string from any value (not for array)
func InterfaceToString(i interface{}) string {
	switch value := i.(type) {
	case string:
		return string(value)
	case int:
		return strconv.Itoa(value)
	case int64:
		return strconv.FormatInt(value, 64)
	//case uint32:
	case uint64:
		return strconv.FormatUint(value, 64)
	//case float32:
	case float64:
		return strconv.FormatFloat(value, _FMT, 0, 64)
	case bool:
		if value {
			return "true"
		}
		return "false"
	default:
		return "";

	}
}