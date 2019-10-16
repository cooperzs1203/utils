/**
* @Author: Cooper
* @Date: 2019/10/15 16:06
 */

package utils

import (
	"fmt"
	"strconv"
)

// transform a string to a int value , return zero if there is an error
// str string : pending conversion string
func StringToInt64(str string) int64 {
	return StringToInt64WithDefaultValue(str , 0)
}

// transform a string to a int value , return default value if there is an error
// str string : pending conversion string
// defaultValue int64 : return value if make an error
func StringToInt64WithDefaultValue(str string, defaultValue int64) int64 {
	number , err := strconv.ParseInt(str , 10 , 64)
	if err != nil {
		return defaultValue
	}
	return number
}

// transform a string to a float value , return 0.00 if these is an error
// str string : pending conversion string
func StringToFloat64(str string) float64 {
	return StringToFloat64WithDefaultValue(str , 0.00)
}

// transform a string to a float value , return default value if there is an error
// str string : pending conversion string
// defaultValue float64 : return value if make an error
func StringToFloat64WithDefaultValue(str string , defaultValue float64) float64 {
	number , err := strconv.ParseFloat(str , 64)
	if err != nil {
		return defaultValue
	}
	return number
}

// transform a float to a string value
// number float64 : pending conversion number
func Float64ToString(number float64) string {
	return fmt.Sprintf("%f" , number)
}