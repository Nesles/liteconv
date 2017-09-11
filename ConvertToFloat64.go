package liteconv

import (
	"strconv"
	"encoding/json"
)

// ToFloat64 converts interface{} to float64
func InterfaceToFloat64(i1 interface{}) float64 {
	if i1 == nil {
		return _VAL_FALSE_FLOAT64
	}
	switch value := i1.(type) {
	default:
		result, _ := strconv.ParseFloat(InterfaceToString(value), _BIT_SIZE_64)
		return result
	case *json.Number:
		result, err := value.Float64()
		if err != nil {
			return _VAL_NULL_FLOAT64
		}
		return result
	case json.Number:
		result, err := value.Float64()
		if err != nil {
			return _VAL_NULL_FLOAT64
		}
		return result
	case int64:
		return float64(value)
	case float32:
		return float64(value)
	case uint64:
		return float64(value)
	case int:
		return float64(value)
	case uint:
		return float64(value)
	case bool:
		if value {
			return _VAL_TRUE_FLOAT64
		} else {
			return _VAL_FALSE_FLOAT64
		}
	case float64:
		return value
	case *bool:
		if value == nil {
			return _VAL_FALSE_FLOAT64
		}
		if *value {
			return _VAL_TRUE_FLOAT64
		} else {
			return _VAL_FALSE_FLOAT64
		}
	case string:
		result,err := strconv.ParseFloat(value, _BIT_SIZE_64)
		if err != nil {
			return _VAL_NULL_FLOAT64
		}
		return result
	}
	return _VAL_FALSE_FLOAT64
}

func StringToFloat64(str string) float64 {

	result, err :=strconv.ParseFloat(str, _BIT_SIZE_64)
	if err!=nil{
		return 0
	}
	return  result

}