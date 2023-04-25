package stp

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
)

// 简单结构体
type simpleStruct struct {
	b bool
	v int
}

func (s *simpleStruct) Init() { s.b, s.v = true, math.MaxInt64 }

func (s *simpleStruct) Use() { s.b, s.v = false, rand.Intn(s.v) }

func (s *simpleStruct) Free() { s.b, s.v = false, 0 } // null method

// 内存池测试函数
func poolFoo(c int) {
	// 初始化：
	// - 需要显式传递类型参数：结构体
	// 分配内存：
	// - 不需要外部计算对象大小
	// 获取对象：
	// - 不需要外部提供数组的下标
	// - 不需要显式传递类型参数
	// - 不需要定义任何接口
	pool := NewPool[simpleStruct](c)
	fmt.Printf("pool = %v\n", pool.b)

	for i := 0; i < c; i++ {
		o := pool.Get()
		o.Init()
		fmt.Printf("i %v, o %v, ptr %p\n", i, o, &o)
		o.Use()
		fmt.Printf("i %v, o %v, ptr %p\n", i, o, &o)
		fmt.Printf("i %v, pool = %v\n", i, pool.b)
	}
}

func Test_poolFoo(t *testing.T) {
	type args struct {
		c int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"test case 1",
			args{c: 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			poolFoo(tt.args.c)
		})
	}
}
