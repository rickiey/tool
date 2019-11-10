package tool

import "reflect"

// 结构体转 map 根据 json tag, 没有就把名字 ToSnakeCase
func StructToMap(obj interface{}) map[string]interface{} {
	obj1 := reflect.TypeOf(obj)
	obj2 := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < obj1.NumField(); i++ {
		k := obj1.Field(i).Tag.Get("json")
		if k == "" {
			k = ToSnakeCase(obj1.Field(i).Name)
		}
		data[k] = obj2.Field(i).Interface()
	}
	return data
}
