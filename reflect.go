package stp

import (
	"reflect"
	"unsafe"
)

func ReflectStructFieldKeyIndexMap(rt reflect.Type, tagKey string) map[string]int {
	if rt.Kind() != reflect.Struct {
		return nil
	}

	useTagKey := len(tagKey) > 0
	fieldCount := rt.NumField()
	kiMap := make(map[string]int)
	for i := 0; i < fieldCount; i++ {
		if useTagKey {
			key, has := rt.Field(i).Tag.Lookup(tagKey)
			if !has {
				continue
			}
			kiMap[key] = i
		} else {
			kiMap[rt.Field(i).Name] = i
		}
	}

	return kiMap
}

func ReflectStructFieldKeyOffsetMap(rt reflect.Type, tagKey string) map[string]uintptr {
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

	return koMap
}

// ReflectStructValue
// @param1            struct field slice
// @param2            struct value
// @param3            tagKey if need
// return             *T
func ReflectStructValue[T any](fs []string, v []any, tagKey string) *T {
	var t T
	rt := reflect.TypeOf(t)
	koMap := ReflectStructFieldKeyOffsetMap(rt, tagKey)
	if len(koMap) == 0 {
		return nil
	}

	tPtr := &t
	for i, f := range fs {
		offset, has := koMap[f]
		if !has {
			panic(f)
		}
		l, r := (unsafe.Pointer(uintptr(unsafe.Pointer(tPtr)) + offset)), v[i]
		switchTypeKind(rt.Field(i).Type.Kind())(l, r)
	}

	return &t
}

// ReflectStructValueSlice
// @param1                 struct field slice
// @param2                 struct value slice
// @param3                 tagKey if need
// return                  []*T
func ReflectStructValueSlice[T any](fs []string, vs [][]any, tagKey string) []*T {
	var t T
	rt := reflect.TypeOf(t)
	koMap := ReflectStructFieldKeyOffsetMap(rt, tagKey)
	if len(koMap) == 0 {
		return nil
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
