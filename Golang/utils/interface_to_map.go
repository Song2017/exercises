package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func InterfaceToMap(v interface{}) map[string]interface{} {
	valueOf := reflect.ValueOf(v)

	if valueOf.Kind() == reflect.Interface {
		valueOf = valueOf.Elem()
	}

	if valueOf.Kind() == reflect.Ptr {
		valueOf = valueOf.Elem()
	}

	switch valueOf.Kind() {
	case reflect.Map:
		resultMap := make(map[string]interface{})
		for _, key := range valueOf.MapKeys() {
			val := valueOf.MapIndex(key).Interface()
			valueOf := reflect.ValueOf(val)
			if valueOf.Kind() == reflect.String || valueOf.Kind() == reflect.Int {
				resultMap[key.String()] = val
			} else if valueOf.Kind() == reflect.Slice {
				resultSlice := make([]interface{}, valueOf.Len())
				for i := 0; i < valueOf.Len(); i++ {
					v2 := valueOf.Index(i).Interface()
					vo2 := reflect.ValueOf(v2)
					if vo2.Kind() == reflect.String || vo2.Kind() == reflect.Int {
						resultSlice[i] = v2
					} else {
						resultSlice[i] = InterfaceToMap(v2)
					}
				}
				resultMap[key.String()] = resultSlice
			} else {
				resultMap[key.String()] = InterfaceToMap(val)
			}
		}
		return resultMap
	case reflect.Struct:
		resultMap := make(map[string]interface{})
		for i := 0; i < valueOf.NumField(); i++ {
			fieldValue := valueOf.Field(i).Interface()
			if reflect.TypeOf(fieldValue).Kind() == reflect.Struct {
				resultMap[valueOf.Type().Field(i).Name] = InterfaceToMap(fieldValue)
			} else {
				resultMap[valueOf.Type().Field(i).Name] = fieldValue
			}
		}
		return resultMap
	default:
		return map[string]interface{}{"value": v}
	}
}

func main() {
	var data interface{}
	jsonStr := `{"name":"Tom","age":18,"address":{"city":"Beijing","detail":{"room":"as6"},"detail2":{"room":"as6"}},"skills":["Golang","Python","Java"]}`
	json.Unmarshal([]byte(jsonStr), &data)

	resultMap := InterfaceToMap(data)
	fmt.Println(resultMap)
	address, ok := resultMap["address"].(map[string]interface{})
    delete(address, "detail2")

	if ok {
		// 对 address 进行索引操作
		detail, ok := address["detail"].(map[string]interface{})
		if ok {
			fmt.Println(111, detail["room"], detail)
		}
		// value, ok := room.(string)
		// fmt.Println(value, ok)
	}
    fmt.Println(999, resultMap)
}
