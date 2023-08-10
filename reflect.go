package stp

import (
	"reflect"
	"unsafe"
)

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

// AssignStructMember
// @param1           field slice
// @param2           value slice
// @param3           tagKey if need
// return            *T slice
func AssignStructMember[T any](fs []string, vs [][]any, tagKey string) []*T {
	var t T
	rt := reflect.TypeOf(t)
	if rt.Kind() != reflect.Struct {
		return nil
	}

	useTagKey := len(tagKey) > 0

	fieldCount := rt.NumField()
	koMap := make(map[string]uintptr)
	for i := 0; i < fieldCount; i++ {
		if useTagKey {
			key, has := rt.Field(i).Tag.Lookup(tagKey)
			if !has {
				continue
			}
			koMap[key] = rt.Field(i).Offset
		} else {
			koMap[rt.Field(i).Name] = rt.Field(i).Offset
		}
	}

	ss := make([]*T, 0, len(vs))
	for _, d := range vs {
		var s T
		sPtr := &s
		for i, f := range fs {
			offset, has := koMap[f]
			if !has {
				panic(f)
			}
			l, r := (unsafe.Pointer(uintptr(unsafe.Pointer(sPtr)) + offset)), d[i]
			switchTypeKind(rt.Field(i).Type.Kind())(l, r)
		}
		ss = append(ss, sPtr)
	}

	return ss
}

func switchTypeKind(k reflect.Kind) func(unsafe.Pointer, any) {
	switch k {
	case reflect.Bool:
		return assign[bool]
	case reflect.Int:
		return assign[int]
	case reflect.Int8:
		return assign[int8]
	case reflect.Int16:
		return assign[int16]
	case reflect.Int32:
		return assign[int32]
	case reflect.Int64:
		return assign[int64]
	case reflect.Uint:
		return assign[uint]
	case reflect.Uint8:
		return assign[uint8]
	case reflect.Uint16:
		return assign[uint16]
	case reflect.Uint32:
		return assign[uint32]
	case reflect.Uint64:
		return assign[uint64]
	case reflect.Float32:
		return assign[float32]
	case reflect.Float64:
		return assign[float64]
	case reflect.Complex64:
		return assign[complex64]
	case reflect.Complex128:
		return assign[complex128]
	case reflect.String:
		return assign[string]
	}
	return nil
}

func assign[T any](l unsafe.Pointer, r any) {
	*(*T)(l) = r.(T)
}
