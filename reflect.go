package stp

import "reflect"

func ReflectStructFieldKeyIndexMap[T any](s T, tagKey string) map[string]int {
	rs := reflect.TypeOf(s)
	fieldCount := rs.NumField()
	if fieldCount == 0 {
		return nil
	}
	kiMap := make(map[string]int)
	for i := 0; i < fieldCount; i++ {
		jsonKey, has := rs.Field(i).Tag.Lookup(tagKey)
		if !has {
			continue
		}
		kiMap[jsonKey] = i
	}
	return kiMap
}
