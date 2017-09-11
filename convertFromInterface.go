package liteconv

import (
	"strconv"
	"encoding/json"
)

//InterfaceToString - get string from any value (not for array)
func InterfaceToString(i interface{}) string {
	switch value := i.(type) {
	case string:
		return string(value)
	case int:
		return strconv.Itoa(value)
	case *int:
		return strconv.Itoa(*value)
	case int64:
		return strconv.FormatInt(value, 32)
	case *int64:
		return strconv.FormatInt(*value, 32)
	case uint64:
		return strconv.FormatUint(value, 32)
	case *uint64:
		return strconv.FormatUint(*value, 32)
	case float32:
		return strconv.FormatFloat(float64(value), _FMT, _PREC, _BIT_SIZE_64)
	case *float32:
		return strconv.FormatFloat(float64(*value), _FMT, _PREC, _BIT_SIZE_64)
	case float64:
		return strconv.FormatFloat(value, _FMT, _PREC, _BIT_SIZE_64)
	case *float64:
		return strconv.FormatFloat(*value, _FMT, _PREC, _BIT_SIZE_64)
	case bool:
		if value {
			return _VAL_TRUE_STR
		}
		return _VAL_FALSE_STR
	case *bool:
		if *value {
			return _VAL_TRUE_STR
		}
		return _VAL_FALSE_STR
	default:
		return _VAL_NULL_STR

	}
}
/*

func InterfaceToFloat64(i interface{}) float64 {
	switch value := i.(type) {
	default:
		res, _ := strconv.ParseFloat(InterfaceToString(), _BIT_SIZE_64)
		return res
	case *json.Number:
		res, _ := value.Float64()
		return res
	case json.Number:
		res, _ := value.Float64()
		return res
	case int64:
		return float64(value)
	case float64:
		return value
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

	case *bool:
		if *value {
			return _VAL_TRUE_FLOAT64
		} else {
			return _VAL_FALSE_FLOAT64
		}

		return _VAL_NULL_FLOAT64
	}
}
*/

//InterfaceToInt
func InterfaceToInt(i interface{}) int {

	switch value := i.(type) {
	case string:
		result , err := strconv.Atoi(value)
		if err!= nil{
			return 0
		}
		return result
	case int:
		return value
	case int64:
		return int(value)
	case uint32:
		return int(value)
	case uint64:
		return int(value)
	case float32:
		return int(value)
	case float64:
		return int(value)
	case bool:
		if value {
			return _VAL_TRUE_INT
		}
		return _VAL_FALSE_INT

	case *string:
		result , err := strconv.Atoi(*value)
		if err!= nil{
			return 0
		}
		return result
	case *int:
		return *value
	case *int64:
		return int(*value)
	case *uint32:
		return int(*value)
	case *uint64:
		return int(*value)
	case *float32:
		return int(*value)
	case *float64:
		return int(*value)
	case *bool:
		if *value {
			return _VAL_TRUE_INT
		}
		return _VAL_FALSE_INT

	case *json.Number:
		result, err := value.Int64()
		if err != nil {
			return _VAL_NULL_INT
		}
		return int(result)
	case json.Number:
		result, err := value.Int64()
		if err != nil {
			return _VAL_NULL_INT
		}
		return int(result)
	default:
		return _VAL_NULL_INT
	}
}

// InterfaceToStringMap
func InterfaceToStringMap(some interface{}) map[string]string {
	switch value := some.(type) {
	case map[string]interface{}:
		result := map[string]string{}
		for k, v := range value {
			result[k] = InterfaceToString(v)
		}
		return result
	case map[string]string:
		return value
	default:
		return map[string]string{}
	}
}

func InterfaceToArray(some interface{}) []interface{}  {
	//:TODO
	return *new([]interface{})
}