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
