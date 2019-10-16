package utils

import (
	"fmt"
	"reflect"
	"strconv"
	"time"
)

// export key and value of struct to make a map
// obj interface{} : padding conversion object
func StructToMap(obj interface{})  map[string]interface{} {
	dic := make(map[string]interface{})
	if obj == nil {
		return dic
	}

	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	for i := 0; i < t.NumField(); i++ {
		if v.Field(i).CanInterface() {
			dic[t.Field(i).Name] = v.Field(i).Interface()
		}
	}

	return dic
}

// accept any type of data replace to string
func AutoToString(data interface{}) string {
	if data == nil {
		return ""
	}

	t := reflect.TypeOf(data)

	switch t.Kind() {
		case reflect.Bool:
			return strconv.FormatBool(data.(bool))
		case reflect.Int:
			return strconv.FormatInt(int64(data.(int)) , 10)
		case reflect.Int8:
			return strconv.FormatInt(int64(data.(int8)) , 10)
		case reflect.Int16:
			return strconv.FormatInt(int64(data.(int16)) , 10)
		case reflect.Int32:
			return strconv.FormatInt(int64(data.(int32)) , 10)
		case reflect.Int64:
			return strconv.FormatInt(data.(int64), 10)
		case reflect.Uint:
			return strconv.FormatUint(uint64(data.(uint)), 10)
		case reflect.Uint8:
			return strconv.FormatUint(uint64(data.(uint8)), 10)
		case reflect.Uint16:
			return strconv.FormatUint(uint64(data.(uint16)), 10)
		case reflect.Uint32:
			return strconv.FormatUint(uint64(data.(uint32)), 10)
		case reflect.Uint64:
			return strconv.FormatUint(data.(uint64), 10)
		case reflect.Float32:
			return fmt.Sprintf("%f" , data.(float32))
		case reflect.Float64:
			return fmt.Sprintf("%f" , data.(float64))
		case reflect.String:
			return data.(string)
		//case reflect.Uintptr
		//case reflect.Complex64
		//case reflect.Complex128
		//case reflect.Array:
		//case reflect.Chan:
		//case reflect.Func:
		//case reflect.Interface:
		//case reflect.Map:
		//case reflect.Ptr
		//case reflect.Slice
		//case reflect.Struct
		//case reflect.UnsafePointer
	default:
		return fmt.Sprintf("%+v" , data)
	}
}

// a ticker which interval [interval] seconds do [f]
func Ticker(f func(int64)bool , interval int64) {
	go func() {
		count := int64(0)
		ticker := time.NewTicker(time.Duration(interval)*time.Second)

		for {
			<- ticker.C
			count++
			if stop := f(count); stop {
				ticker.Stop()
				break
			}
		}
	}()
}