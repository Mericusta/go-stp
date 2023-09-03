package stp

import (
	"bytes"
	"io"
	"strings"
	"unsafe"

	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

// ConvertStringToStringStruct
// @template1                  string struct
// @param1                     to handle string
// @param2                     string splitter
// @return                     string struct pointer
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

func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func StringToBytes(s string) []byte {
	if len(s) < (1 << 5) {
		return standardStringToBytes(s)
	}
	return unsafeStringToBytes(s)
}

func standardStringToBytes(s string) []byte {
	return []byte(s)
}

func unsafeStringToBytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(
		&struct {
			string
			Cap int
		}{s, len(s)},
	))
}

type Code interface {
	GBK | GB18030
	decoder() *encoding.Decoder
	encoder() *encoding.Encoder
}

type GBK struct{}

func (gbk GBK) decoder() *encoding.Decoder {
	return simplifiedchinese.GBK.NewDecoder()
}

func (gbk GBK) encoder() *encoding.Encoder {
	return simplifiedchinese.GBK.NewEncoder()
}

type GB18030 struct{}

func (gb18030 GB18030) decoder() *encoding.Decoder {
	return simplifiedchinese.GB18030.NewDecoder()
}

func (gb18030 GB18030) encoder() *encoding.Encoder {
	return simplifiedchinese.GB18030.NewEncoder()
}

func Utf8To[T Code](b []byte) ([]byte, error) {
	var code T
	transformer := transform.NewReader(bytes.NewReader(b), code.encoder())
	_b, err := io.ReadAll(transformer)
	if err != nil {
		return nil, err
	}
	return _b, nil
}

func ToUtf8[T Code](b []byte) ([]byte, error) {
	var code T
	transformer := transform.NewReader(bytes.NewReader(b), code.decoder())
	_b, err := io.ReadAll(transformer)
	if err != nil {
		return nil, err
	}
	return _b, nil
}
