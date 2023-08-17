package stp

import (
	"reflect"
	"testing"
)

func Test_ReflectStructFieldKeyIndexMap(t *testing.T) {
	type args struct {
		rt     reflect.Type
		tagKey string
	}
	tests := []struct {
		name  string
		args  args
		want0 map[string]int
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got0 := ReflectStructFieldKeyIndexMap(tt.args.rt, tt.args.tagKey)
			if !reflect.DeepEqual(got0, tt.want0) {
				t.Errorf("ReflectStructFieldKeyIndexMap() got0 = %v, want0 %v", got0, tt.want0)
			}
		})
	}
}

func Test_ReflectStructFieldKeyOffsetMap(t *testing.T) {
	type args struct {
		rt     reflect.Type
		tagKey string
	}
	tests := []struct {
		name  string
		args  args
		want0 map[string]uintptr
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got0 := ReflectStructFieldKeyOffsetMap(tt.args.rt, tt.args.tagKey)
			if !reflect.DeepEqual(got0, tt.want0) {
				t.Errorf("ReflectStructFieldKeyOffsetMap() got0 = %v, want0 %v", got0, tt.want0)
			}
		})
	}
}

func Test_ReflectStructValue_b5c367bcf55d571f2cbc8e1e8b62636d(t *testing.T) {
	type T struct {
		I int     `json:"int"`
		F float64 `json:"float64"`
		S string  `json:"string"`
	}
	type args struct {
		fs     []string
		v      []any
		tagKey string
	}
	tests := []struct {
		name  string
		args  args
		want0 *T
	}{
		{
			"test case 1",
			args{
				fs:     []string{"int", "float64", "string"},
				v:      []any{1024, 0.618, "this is gold ratio"},
				tagKey: "json",
			},
			&T{1024, 0.618, "this is gold ratio"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got0 := ReflectStructValue[T](tt.args.fs, tt.args.v, tt.args.tagKey)
			if !reflect.DeepEqual(got0, tt.want0) {
				t.Errorf("ReflectStructValue() got0 = %v, want0 %v", got0, tt.want0)
			}
		})
	}
}

func Test_ReflectStructValueSlice_b5c367bcf55d571f2cbc8e1e8b62636d(t *testing.T) {
	type T struct {
		I int     `json:"int"`
		F float64 `json:"float64"`
		S string  `json:"string"`
	}
	type args struct {
		fs     []string
		vs     [][]any
		tagKey string
	}
	tests := []struct {
		name  string
		args  args
		want0 []*T
	}{
		{
			"test case 1",
			args{
				fs: []string{"int", "float64", "string"},
				vs: [][]any{
					{1024, 0.618, "this is gold ratio"},
					{2048, 3.1415, "this is pi"},
					{4096, 9.8, "this is gravity"},
				},
				tagKey: "json",
			},
			[]*T{
				{1024, 0.618, "this is gold ratio"},
				{2048, 3.1415, "this is pi"},
				{4096, 9.8, "this is gravity"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got0 := ReflectStructValueSlice[T](tt.args.fs, tt.args.vs, tt.args.tagKey)
			if !reflect.DeepEqual(got0, tt.want0) {
				t.Errorf("ReflectStructValueSlice() got0 = %v, want0 %v", got0, tt.want0)
			}
		})
	}
}
