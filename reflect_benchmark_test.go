package stp

import (
	"testing"
	"time"
)

func Benchmark_ReflectStructValue_b5c367bcf55d571f2cbc8e1e8b62636d(b *testing.B) {
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
		limit time.Duration
		args  args
	}{
		{
			"test case 1",
			time.Millisecond,
			args{
				fs:     []string{"int", "float64", "string"},
				v:      []any{1024, 0.618, "this is gold ratio"},
				tagKey: "json",
			},
		},
	}
	for _, tt := range tests {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = ReflectStructValue[T](tt.args.fs, tt.args.v, tt.args.tagKey)
		}
		b.StopTimer()
		if b.Elapsed() > tt.limit*time.Duration(b.N) {
			b.Fatalf("overtime limit %v, got %.2f\n", tt.limit, float64(b.Elapsed())/float64(b.N))
		}
	}
}

func Benchmark_ReflectStructValueSlice_b5c367bcf55d571f2cbc8e1e8b62636d(b *testing.B) {
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
		limit time.Duration
		args  args
	}{
		{
			"test case 1",
			time.Millisecond,
			args{
				fs: []string{"int", "float64", "string"},
				vs: [][]any{
					{1024, 0.618, "this is gold ratio"},
					{2048, 3.1415, "this is pi"},
					{4096, 9.8, "this is gravity"},
				},
				tagKey: "json",
			},
		},
	}
	for _, tt := range tests {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = ReflectStructValueSlice[T](tt.args.fs, tt.args.vs, tt.args.tagKey)
		}
		b.StopTimer()
		if b.Elapsed() > tt.limit*time.Duration(b.N) {
			b.Fatalf("overtime limit %v, got %.2f\n", tt.limit, float64(b.Elapsed())/float64(b.N))
		}
	}
}
