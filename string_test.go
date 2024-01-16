package stp

import (
	"reflect"
	"testing"
)

type _s struct {
	s1 string
	s2 string
	s3 string
	s4 []string
	s5 string
	// ...
}

func TestConvertStringToStringStruct(t *testing.T) {
	type args struct {
		str      string
		splitter string
	}
	tests := []struct {
		name string
		args args
		want *_s
	}{
		// TODO: Add test cases.
		{
			"test case 1",
			args{
				str:      "I am a boy,You are a girl,We are human",
				splitter: ",",
			},
			&_s{
				s1: "I am a boy",
				s2: "You are a girl",
				s3: "We are human",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertStringToStringStruct[_s](tt.args.str, tt.args.splitter); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertStringToStringStruct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBytesToString(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BytesToString(tt.args.b); got != tt.want {
				t.Errorf("BytesToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringToBytes(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		// TODO: Add test cases.
		{
			"test case 1",
			args{s: "Hello go-stp!"},
			[]byte{72, 101, 108, 108, 111, 32, 103, 111, 45, 115, 116, 112, 33},
		},
		{
			"test case 2",
			args{s: `
				Hello go-stp!Hello go-stp!Hello go-stp!Hello go-stp!
				Hello go-stp!Hello go-stp!Hello go-stp!Hello go-stp!
				Hello go-stp!Hello go-stp!Hello go-stp!Hello go-stp!
				Hello go-stp!Hello go-stp!Hello go-stp!Hello go-stp!
				Hello go-stp!Hello go-stp!Hello go-stp!Hello go-stp!
				Hello go-stp!Hello go-stp!Hello go-stp!Hello go-stp!
				Hello go-stp!Hello go-stp!Hello go-stp!Hello go-stp!
				Hello go-stp!Hello go-stp!Hello go-stp!Hello go-stp!
			`},
			[]byte{
				10, 9, 9, 9,
				9, 72, 101, 108, 108, 111, 32, 103, 111, 45, 115, 116, 112, 33, 72, 101, 108, 108, 111, 32, 103, 111, 45, 115, 116, 112, 33, 72, 101, 108, 108, 111, 32, 103, 111, 45, 115, 116, 112, 33, 72, 101, 108, 108, 111, 32, 103, 111, 45, 115, 116, 112, 33,
				10, 9, 9, 9,
				9, 72, 101, 108, 108, 111, 32, 103, 111, 45, 115, 116, 112, 33, 72, 101, 108, 108, 111, 32, 103, 111, 45, 115, 116, 112, 33, 72, 101, 108, 108, 111, 32, 103, 111, 45, 115, 116, 112, 33, 72, 101, 108, 108, 111, 32, 103, 111, 45, 115, 116, 112, 33,
				10, 9, 9, 9,
				9, 72, 101, 108, 108, 111, 32, 103, 111, 45, 115, 116, 112, 33, 72, 101, 108, 108, 111, 32, 103, 111, 45, 115, 116, 112, 33, 72, 101, 108, 108, 111, 32, 103, 111, 45, 115, 116, 112, 33, 72, 101, 108, 108, 111, 32, 103, 111, 45, 115, 116, 112, 33,
				10, 9, 9, 9,
				9, 72, 101, 108, 108, 111, 32, 103, 111, 45, 115, 116, 112, 33, 72, 101, 108, 108, 111, 32, 103, 111, 45, 115, 116, 112, 33, 72, 101, 108, 108, 111, 32, 103, 111, 45, 115, 116, 112, 33, 72, 101, 108, 108, 111, 32, 103, 111, 45, 115, 116, 112, 33,
				10, 9, 9, 9,
				9, 72, 101, 108, 108, 111, 32, 103, 111, 45, 115, 116, 112, 33, 72, 101, 108, 108, 111, 32, 103, 111, 45, 115, 116, 112, 33, 72, 101, 108, 108, 111, 32, 103, 111, 45, 115, 116, 112, 33, 72, 101, 108, 108, 111, 32, 103, 111, 45, 115, 116, 112, 33,
				10, 9, 9, 9,
				9, 72, 101, 108, 108, 111, 32, 103, 111, 45, 115, 116, 112, 33, 72, 101, 108, 108, 111, 32, 103, 111, 45, 115, 116, 112, 33, 72, 101, 108, 108, 111, 32, 103, 111, 45, 115, 116, 112, 33, 72, 101, 108, 108, 111, 32, 103, 111, 45, 115, 116, 112, 33,
				10, 9, 9, 9,
				9, 72, 101, 108, 108, 111, 32, 103, 111, 45, 115, 116, 112, 33, 72, 101, 108, 108, 111, 32, 103, 111, 45, 115, 116, 112, 33, 72, 101, 108, 108, 111, 32, 103, 111, 45, 115, 116, 112, 33, 72, 101, 108, 108, 111, 32, 103, 111, 45, 115, 116, 112, 33,
				10, 9, 9, 9,
				9, 72, 101, 108, 108, 111, 32, 103, 111, 45, 115, 116, 112, 33, 72, 101, 108, 108, 111, 32, 103, 111, 45, 115, 116, 112, 33, 72, 101, 108, 108, 111, 32, 103, 111, 45, 115, 116, 112, 33, 72, 101, 108, 108, 111, 32, 103, 111, 45, 115, 116, 112, 33,
				10, 9, 9, 9,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringToBytes(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StringToBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}
