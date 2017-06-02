package liteconv

import (
	"strconv"
	"encoding/json"
)

// ToFloat64 converts interface{} to float64
func InterfaceToFloat64(i1 interface{}) float64 {
	if i1 == nil {
		return 0.0
	}
	switch i2 := i1.(type) {
	default:
		i3, _ := strconv.ParseFloat(InterfaceToString(i2), 64)
		return i3
	case *json.Number:
		i3, _ := i2.Float64()
		return i3
	case json.Number:
		i3, _ := i2.Float64()
		return i3
	case int64:
		return float64(i2)
	case float32:
		return float64(i2)
	case uint64:
		return float64(i2)
	case int:
		return float64(i2)
	case uint:
		return float64(i2)
	case bool:
		if i2 {
			return 1.0
		} else {
			return 0.0
		}
	case float64:
		return i2
	case *bool:
		if i2 == nil {
			return 0.0
		}
		if *i2 {
			return 1.0
		} else {
			return 0.0
		}
	}
	return 0.0
}

