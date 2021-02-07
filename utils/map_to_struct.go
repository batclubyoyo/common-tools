package utils

import (
	"fmt"
	"reflect"
	"strconv"
)

/*
SetField key, value为查询结果，将查询结果字段跟obj字段一一对应
*/
func SetField(obj interface{}, key string, value interface{}) error {
	structValue := reflect.ValueOf(obj).Elem()
	structValueType := structValue.Type()

	// 匹配数据库字段跟结构字段
	var name string
	for i := 0; i < structValueType.NumField(); i++ {
		if structValueType.Field(i).Tag.Get("MYSQL") == key {
			name = structValueType.Field(i).Name
		}
	}

	// 获取结构字段值
	structFieldValue := structValue.FieldByName(name)
	if !structFieldValue.IsValid() {
		return fmt.Errorf("No such field: %s in obj", name)
	}
	if !structFieldValue.CanSet() {
		return fmt.Errorf("Cannot set %s field value", name)
	}

	// 类型判断
	structFieldType := structFieldValue.Type()

	val := reflect.ValueOf(value)
	assertValue, _ := value.(string)

	// 这里其实还有时间格式之类的类型没处理
	switch structFieldType.String() {
	case "uint":
		v, _ := strconv.Atoi(assertValue)
		val = reflect.ValueOf(uint(v))
	case "int":
		v, _ := strconv.Atoi(assertValue)
		val = reflect.ValueOf(v)
	case "int64":
		v, _ := strconv.ParseInt(assertValue, 10, 64)
		val = reflect.ValueOf(v)
	case "float32":
		v, _ := strconv.ParseFloat(assertValue, 32)
		val = reflect.ValueOf(v)
	case "float64":
		v, _ := strconv.ParseFloat(assertValue, 64)
		val = reflect.ValueOf(v)
	case "bool":
		v, _ := strconv.ParseBool(assertValue)
		val = reflect.ValueOf(v)
	default:
		v := assertValue
		val = reflect.ValueOf(v)
	}

	structFieldValue.Set(val)
	return nil
}

// obj: target struct
// name: key name of map type param
// value: value of map type param
// func SetField(obj interface{}, name string, value interface{}) error {
// 	structValue := reflect.ValueOf(obj).Elem()
// 	structFieldValue := structValue.FieldByName(name)

// 	if !structFieldValue.IsValid() {
// 		return fmt.Errorf("No such field: %s in obj", name)
// 	}

// 	if !structFieldValue.CanSet() {
// 		return fmt.Errorf("Cannot set %s field value", name)
// 	}

// 	structFieldType := structFieldValue.Type()
// 	val := reflect.ValueOf(value)
// 	if structFieldType != val.Type() {
// 		return errors.New("Provided value type didn't match obj field type")
// 	}

// 	structFieldValue.Set(val)
// 	return nil
// }

// type MyStruct struct {
// 	Name string
// 	Age  int64
// }
//
// func (s *MyStruct) FillStruct(m map[string]interface{}) error {
// 	for k, v := range m {
// 		err := SetField(s, k, v)
// 		if err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }
