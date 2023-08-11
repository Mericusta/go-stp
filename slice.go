package stp

func Compare[T comparable](s1, s2 []T) bool {
	if len(s1) != len(s2) {
		return false
	}
	l := len(s1)
	for index := 0; index != l; index++ {
		if s1[index] != s2[index] {
			return false
		}
	}
	return true
}

type Array[T comparable] struct {
	s         []T
	zeroValue T
}

func NewArray[T comparable](slice []T) *Array[T] {
	return &Array[T]{s: slice}
}

func (a *Array[T]) Slice() []T {
	return a.s
}

// JS Array.concat() 连接两个或更多的数组，并返回结果。
// JS Array.copyWithin() 从数组的指定位置拷贝元素到数组的另一个指定位置中。
// JS Array.entries() 返回数组的可迭代对象。
// JS Array.every() 检测数值元素的每个元素是否都符合条件。
// JS Array.fill() 使用一个固定值来填充数组。

// JS Array.filter() 检测数值元素，并返回符合条件所有元素的数组。
func (a *Array[T]) Filter(f func(v T, i int) bool) *Array[T] {
	_s := make([]T, 0, len(a.s))
	for _i, _v := range a.s {
		if f(_v, _i) {
			_s = append(_s, _v)
		}
	}
	return NewArray(_s)
}

// JS Array.find() 返回符合传入测试（函数）条件的数组元素。
func (a *Array[T]) Find(f func(v T, i int) bool) T {
	for _i, _v := range a.s {
		if f(_v, _i) {
			return _v
		}
	}
	return a.zeroValue
}

// JS Array.findIndex() 返回符合传入测试（函数）条件的数组元素索引。

// JS Array.forEach() 数组每个元素都执行一次回调函数。
func (a *Array[T]) ForEach(f func(v T, i int)) {
	for _i, _v := range a.s {
		f(_v, _i)
	}
}

// JS Array.from() 通过给定的对象中创建一个数组。

// JS Array.includes() 判断一个数组是否包含一个指定的值。
func (a *Array[T]) Includes(v T) bool {
	for _, _v := range a.s {
		if _v == v {
			return true
		}
	}
	return false
}

// JS Array.indexOf() 搜索数组中的元素，并返回它所在的位置。
func (a *Array[T]) IndexOf(v T) int {
	for _i, _v := range a.s {
		if _v == v {
			return _i
		}
	}
	return -1
}

// JS Array.isArray() 判断对象是否为数组。
// JS Array.join() 把数组的所有元素放入一个字符串。
// JS Array.keys() 返回数组的可迭代对象，包含原始数组的键(key)。
// JS Array.lastIndexOf() 搜索数组中的元素，并返回它最后出现的位置。

// JS Array.map() 通过指定函数处理数组的每个元素，并返回处理后的数组。
func (a *Array[T]) Map(f func(v T, i int) T) *Array[T] {
	_s := make([]T, 0, len(a.s))
	for _i, _v := range a.s {
		_s = append(_s, f(_v, _i))
	}
	return NewArray(_s)
}

// JS Array.pop() 删除数组的最后一个元素并返回删除的元素。

// JS Array.push() 向数组的末尾添加一个或更多元素，并返回新的长度。
func (a *Array[T]) Push(vs ...T) int {
	a.s = append(a.s, vs...)
	return len(a.s)
}

// JS Array.reduce() 将数组元素计算为一个值（从左到右）。
// JS Array.reduceRight() 将数组元素计算为一个值（从右到左）。
// JS Array.reverse() 反转数组的元素顺序。

// JS Array.shift() 删除并返回数组的第一个元素。
func (a *Array[T]) Shift() T {
	if len(a.s) == 0 {
		return a.zeroValue
	}
	e := a.s[0]
	a.s = a.s[1:]
	return e
}

// JS Array.some() 检测数组元素中是否有元素符合指定条件。
// JS Array.sort() 对数组的元素进行排序。

// JS.splice() 从数组中添加或删除元素。
func (a *Array[T]) Splice(i, m int) *Array[T] {
	l, _i := len(a.s), i+m
	if m == 0 || i < 0 || i >= l {
		a.s = nil
	} else {
		if _i >= l {
			_i = l - 1
		}
		a.s = append(a.s[0:i], a.s[_i:]...)
	}
	return NewArray(a.s)
}

// JS Array.toString() 把数组转换为字符串，并返回结果。
// JS Array.unshift() 向数组的开头添加一个或更多元素，并返回新的长度。
// JS Array.valueOf() 返回数组对象的原始值。
// JS Array.of() 将一组值转换为数组。
// JS Array.at() 用于接收一个整数值并返回该索引对应的元素，允许正数和负数。负整数从数组中的最后一个元素开始倒数。
// JS Array.flat() 创建一个新数组，这个新数组由原数组中的每个元素都调用一次提供的函数后的返回值组成。
// JS Array.flatMap() 使用映射函数映射每个元素，然后将结果压缩成一个新数组。
