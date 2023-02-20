package convert

import (
	"strings"
	"unsafe"
)

func ConvertStringToStringStruct[T any](str, splitter string) *T {
	sSlice := strings.Split(str, splitter)
	offset := unsafe.Sizeof(str) // must be 16
	len := len(sSlice)

	sStruct := new(T)
	ptr := uintptr(unsafe.Pointer(sStruct))
	for index := uintptr(0); uintptr(len) > index; index++ {
		*(*string)(unsafe.Pointer(ptr + offset*index)) = sSlice[index]
	}

	return sStruct
}
