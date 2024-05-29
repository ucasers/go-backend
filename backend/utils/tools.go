package utils

import (
	"reflect"
)

// StructToMap 将结构体转换为 map[string]interface{}
func structToMap(obj interface{}, exclude ...string) map[string]interface{} {
	objValue := reflect.ValueOf(obj)
	if objValue.Kind() == reflect.Ptr {
		objValue = objValue.Elem()
	}
	objType := objValue.Type()

	data := make(map[string]interface{})

	for i := 0; i < objType.NumField(); i++ {
		field := objType.Field(i)
		fieldName := field.Name

		// 获取字段对应的 JSON tag
		jsonTag := field.Tag.Get("json")
		if jsonTag == "" {
			// 如果没有指定 JSON tag，则默认使用字段名
			jsonTag = fieldName
		}

		// 检查是否在排除列表中
		shouldExclude := false
		for _, ex := range exclude {
			if ex == fieldName {
				shouldExclude = true
				break
			}
		}

		if !shouldExclude {
			fieldValue := objValue.Field(i).Interface()
			data[jsonTag] = fieldValue
		}
	}

	return data
}
