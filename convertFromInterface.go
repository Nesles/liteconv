package liteconv

import "strconv"

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