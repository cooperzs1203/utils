/**
* @Author: Cooper
* @Date: 2019/10/15 16:05
 */

package utils

import (
	"math/rand"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

const (
	randomChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

// get random number from 0 to limit
// limit int : limit max number
func GetRandomNumberWithIn(limit int) int {
	return GetRandomNumberFromMinToMax(0 , limit)
}

// get random number from min to max
func GetRandomNumberFromMinToMax(min, max int) int {
	return min + rand.Intn(max - min)
}

// get a random string of length equal to [length]
// length int : string length
func GetRandomString(length int) string {
	randomBytes := []byte(randomChars)
	strBytes := make([]byte , 0 , length)
	for i := 0; i < length; i++ {
		index := GetRandomNumberWithIn(len(randomBytes))
		strBytes = append(strBytes , randomBytes[index])
	}
	return string(strBytes)
}

// get a random capital string of length equal to [length]
// length int : string length
func GetRandomCapitalString(length int) string {
	return strings.ToUpper(GetRandomString(length))
}

// get a random lower case string of length equal to [length]
// length int : string length
func GetRandomLowerCaseString(length int) string {
	return strings.ToLower(GetRandomString(length))
}