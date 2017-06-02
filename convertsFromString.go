package liteconv

import "strconv"


func StringToFloat64(str string) float64 {

	result, err :=strconv.ParseFloat(str, 64)
	if err!=nil{
		return 0
	}
	return  result

}