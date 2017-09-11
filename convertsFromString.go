//convertsFromString
package liteconv

import (
	"strconv"
)


func StringToInt32(str string) int  {
	i, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}

	return i
}