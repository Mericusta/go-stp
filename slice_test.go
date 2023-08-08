package stp

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_Array_Filter_368fd4a02116f4867a426a9376d517e1(t *testing.T) {
	type args struct {
		f func(v int, i int) bool
	}
	tests := []struct {
		name  string
		a     *Array[int]
		args  args
		want0 *Array[int]
	}{
		{
			"test case 1",
			&Array[int]{s: []int{1, 2, 3, 4, 5}},
			args{
				f: func(v, i int) bool {
					return v%2 == 0
				},
			},
			NewArray([]int{2, 4}),
		},
		{
			"test case 2",
			&Array[int]{s: []int{1, 2, 3, 4, 5}},
			args{
				f: func(v, i int) bool {
					return v%2 != 0
				},
			},
			NewArray([]int{1, 3, 5}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got0 := tt.a.Filter(tt.args.f)
			if !reflect.DeepEqual(got0, tt.want0) {
				t.Errorf("Filter() got0 = %v, want0 %v", got0, tt.want0)
			}
		})
	}
}

func Test_Array_Find_368fd4a02116f4867a426a9376d517e1(t *testing.T) {
	type args struct {
		f func(v int, i int) bool
	}
	tests := []struct {
		name  string
		a     *Array[int]
		args  args
		want0 int
	}{
		{
			"test case 1",
			&Array[int]{s: []int{1, 2, 3, 4, 5}},
			args{
				f: func(v, i int) bool {
					return v%2 == 0
				},
			},
			2,
		},
		{
			"test case 2",
			&Array[int]{s: []int{1, 2, 3, 4, 5}},
			args{
				f: func(v, i int) bool {
					return v%2 != 0
				},
			},
			1,
		},
		{
			"test case 3",
			&Array[int]{s: []int{1, 2, 3, 4, 5}},
			args{
				f: func(v, i int) bool {
					return v > 5
				},
			},
			(&Array[int]{}).zeroValue,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got0 := tt.a.Find(tt.args.f)
			if !reflect.DeepEqual(got0, tt.want0) {
				t.Errorf("Find() got0 = %v, want0 %v", got0, tt.want0)
			}
		})
	}
}

func Test_Array_ForEach_368fd4a02116f4867a426a9376d517e1(t *testing.T) {
	type args struct {
		f func(v int, i int)
	}
	tests := []struct {
		name string
		a    *Array[int]
		args args
	}{
		{
			"test case 1",
			&Array[int]{s: []int{1, 2, 3, 4, 5}},
			args{
				f: func(v, i int) {
					if v%2 == 0 {
						fmt.Printf("test case 1, v, i = %v, %v\n", v, i)
					}
				},
			},
		},
		{
			"test case 2",
			&Array[int]{s: []int{1, 2, 3, 4, 5}},
			args{
				f: func(v, i int) {
					if v%2 != 0 {
						fmt.Printf("test case 2, v, i = %v, %v\n", v, i)
					}
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.a.ForEach(tt.args.f)
		})
	}
}

func Test_Array_Includes_368fd4a02116f4867a426a9376d517e1(t *testing.T) {
	type args struct {
		v int
	}
	tests := []struct {
		name  string
		a     *Array[int]
		args  args
		want0 bool
	}{
		{
			"test case 1",
			&Array[int]{s: []int{1, 2, 3, 4, 5}},
			args{v: 3},
			true,
		},
		{
			"test case 1",
			&Array[int]{s: []int{1, 2, 3, 4, 5}},
			args{v: 6},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got0 := tt.a.Includes(tt.args.v)
			if !reflect.DeepEqual(got0, tt.want0) {
				t.Errorf("Includes() got0 = %v, want0 %v", got0, tt.want0)
			}
		})
	}
}

func Test_Array_Map_368fd4a02116f4867a426a9376d517e1(t *testing.T) {
	type args struct {
		f func(v int, i int) int
	}
	tests := []struct {
		name  string
		a     *Array[int]
		args  args
		want0 *Array[int]
	}{
		{
			"test case 1",
			&Array[int]{s: []int{1, 2, 3, 4, 5}},
			args{
				f: func(v, i int) int {
					return v * v
				},
			},
			NewArray([]int{1, 4, 9, 16, 25}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got0 := tt.a.Map(tt.args.f)
			if !reflect.DeepEqual(got0, tt.want0) {
				t.Errorf("Map() got0 = %v, want0 %v", got0, tt.want0)
			}
		})
	}
}
