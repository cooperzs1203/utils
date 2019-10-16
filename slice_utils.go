/**
* @Author: Cooper
* @Date: 2019/10/15 20:38
 */

package utils

// find all indexes of the int64 element type from the array
// array []int64 ： padding conversion array
// element int64 : padding conversion element
func FindInt64ElementIndex(array []int64 , element int64) []int {
	indexs := make([]int , 0)
	for index , value := range array {
		if value == element {
			indexs = append(indexs , index)
		}
	}
	return indexs
}

// find all indexes of the float64 element type from the array
// array []float64 ： padding conversion array
// element float64 : padding conversion element
func FindFloat64ElementIndex(array []float64 , element float64) []int {
	indexs := make([]int , 0)
	for index , value := range array {
		if value == element {
			indexs = append(indexs , index)
		}
	}
	return indexs
}

// find all indexes of the string element type from the array
// array []string ： padding conversion array
// element string : padding conversion element
func FindStringElementIndex(array []string , element string) []int {
	indexs := make([]int , 0)
	for index , value := range array {
		if value == element {
			indexs = append(indexs , index)
		}
	}
	return indexs
}

// make sure element exists in array
// array []int64 ： padding conversion array
// element int64 : padding conversion element
func ContainInt64Element(array []int64 , element int64) bool {
	return len(FindInt64ElementIndex(array , element))>0
}

// make sure element exists in array
// array []float64 ： padding conversion array
// element float64 : padding conversion element
func ContainFloat64Element(array []float64 , element float64) bool {
	return len(FindFloat64ElementIndex(array , element))>0
}

// make sure element exists in array
// array []string ： padding conversion array
// element string : padding conversion element
func ContainStringElement(array []string , element string) bool {
	return len(FindStringElementIndex(array , element))>0
}

// delete the element with index [index] from the int64 array ,and return origin array when index greater than the length of array
// array []int64 : padding conversion array
// index int : index of padding conversion element
func DeleteInt64ElementOfIndex(array []int64, index int) []int64 {
	if index >= len(array) {
		return array
	}
	interfaceArray := make([]interface{} , 0)
	for _ , value := range array {
		interfaceArray = append(interfaceArray , value)
	}

	newArray := make([]int64 , 0)
	for _ , value := range DeleteElementOfIndex(interfaceArray , index) {
		newArray = append(newArray , value.(int64))
	}
	return newArray
}

// delete the element with index [index] from the float64 array ,and return origin array when index greater than the length of array
// array []float64 : padding conversion array
// index int : index of padding conversion element
func DeleteFloat64ElementOfIndex(array []float64, index int) []float64 {
	if index >= len(array) {
		return array
	}
	interfaceArray := make([]interface{} , 0)
	for _ , value := range array {
		interfaceArray = append(interfaceArray , value)
	}

	newArray := make([]float64 , 0)
	for _ , value := range DeleteElementOfIndex(interfaceArray , index) {
		newArray = append(newArray , value.(float64))
	}
	return newArray
}

// delete the element with index [index] from the string array ,and return origin array when index greater than the length of array
// array []string : padding conversion array
// index int : index of padding conversion element
func DeleteStringElementOfIndex(array []string, index int) []string {
	if index >= len(array) {
		return array
	}
	interfaceArray := make([]interface{} , 0)
	for _ , value := range array {
		interfaceArray = append(interfaceArray , value)
	}

	newArray := make([]string , 0)
	for _ , value := range DeleteElementOfIndex(interfaceArray , index) {
		newArray = append(newArray , value.(string))
	}
	return newArray
}

// delete the element with index [index] from the interface{} array ,and return origin array when index greater than the length of array
// array []interface{} : padding conversion array
// index int : index of padding conversion element
func DeleteElementOfIndex(array []interface{}, index int) []interface{} {
	newArray := make([]interface{} , 0)
	newArray = append(newArray , array[:index]...)
	newArray = append(newArray , array[index+1:]...)
	return newArray
}

// insert an int64 type element into array at [index]
// array []int64 : padding conversion array
// element int64 : padding conversion element
// index int : index of insert
func InsertInt64ElementIntoArrayAtIndex(array []int64, element int64, index int) []int64 {
	elements := []int64{element}
	return InsertInt64ElementsIntoArrayAtIndex(array , elements , index)
}

// insert an float64 type element into array at [index]
// array []float64 : padding conversion array
// element float64 : padding conversion element
// index int : index of insert
func InsertFloat64ElementIntoArrayAtIndex(array []float64, element float64, index int) []float64 {
	elements := []float64{element}
	return InsertFloat64ElementsIntoArrayAtIndex(array , elements , index)
}

// insert an string type element into array at [index]
// array []string : padding conversion array
// element string : padding conversion element
// index int : index of insert
func InsertStringElementIntoArrayAtIndex(array []string, element string, index int) []string {
	elements := []string{element}
	return InsertStringElementsIntoArrayAtIndex(array , elements , index)
}

// insert an array of type int64 into the array at [index]
// array []int64 : padding conversion array
// elements []int64 : padding conversion elements
// index int : index of insert
func InsertInt64ElementsIntoArrayAtIndex(array []int64, elements []int64, index int) []int64 {
	newArray := make([]int64 , 0)
	newArray = append(newArray , array[:index]...)
	newArray = append(newArray , elements...)
	newArray = append(newArray , array[index:]...)
	return newArray
}

// insert an array of type float64 into the array at [index]
// array []float64 : padding conversion array
// elements []float64 : padding conversion elements
// index int : index of insert
func InsertFloat64ElementsIntoArrayAtIndex(array []float64, elements []float64, index int) []float64 {
	newArray := make([]float64 , 0)
	newArray = append(newArray , array[:index]...)
	newArray = append(newArray , elements...)
	newArray = append(newArray , array[index:]...)
	return newArray
}

// insert an array of type string into the array at [index]
// array []string : padding conversion array
// elements []string : padding conversion elements
// index int : index of insert
func InsertStringElementsIntoArrayAtIndex(array []string, elements []string, index int) []string {
	newArray := make([]string , 0)
	newArray = append(newArray , array[:index]...)
	newArray = append(newArray , elements...)
	newArray = append(newArray , array[index:]...)
	return newArray
}