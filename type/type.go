package stptype

type STPComparable[T any] interface {
	EQ(T) bool
}

type STPOrdered[T any] interface {
	LT(T) bool
	GT(T) bool
}

type STPType[T any] interface {
	V() T
}

type Int int

func (v Int) LT(rv Int) bool { return v < rv }
func (v Int) GT(rv Int) bool { return v > rv }
func (v Int) EQ(rv Int) bool { return v == rv }
func (v Int) V() Int         { return v }

type Int8 int8

func (v Int8) LT(rv Int8) bool { return v < rv }
func (v Int8) GT(rv Int8) bool { return v > rv }
func (v Int8) EQ(rv Int8) bool { return v == rv }
func (v Int8) V() Int8         { return v }

type Int16 int16

func (v Int16) LT(rv Int16) bool { return v < rv }
func (v Int16) GT(rv Int16) bool { return v > rv }
func (v Int16) EQ(rv Int16) bool { return v == rv }
func (v Int16) V() Int16         { return v }

type Int32 int32

func (v Int32) LT(rv Int32) bool { return v < rv }
func (v Int32) GT(rv Int32) bool { return v > rv }
func (v Int32) EQ(rv Int32) bool { return v == rv }
func (v Int32) V() Int32         { return v }

type Int64 int64

func (v Int64) LT(rv Int64) bool { return v < rv }
func (v Int64) GT(rv Int64) bool { return v > rv }
func (v Int64) EQ(rv Int64) bool { return v == rv }
func (v Int64) V() Int64         { return v }

type Uint uint

func (v Uint) LT(rv Uint) bool { return v < rv }
func (v Uint) GT(rv Uint) bool { return v > rv }
func (v Uint) EQ(rv Uint) bool { return v == rv }
func (v Uint) V() Uint         { return v }

type Uint8 uint8

func (v Uint8) LT(rv Uint8) bool { return v < rv }
func (v Uint8) GT(rv Uint8) bool { return v > rv }
func (v Uint8) EQ(rv Uint8) bool { return v == rv }
func (v Uint8) V() Uint8         { return v }

type Uint16 uint16

func (v Uint16) LT(rv Uint16) bool { return v < rv }
func (v Uint16) GT(rv Uint16) bool { return v > rv }
func (v Uint16) EQ(rv Uint16) bool { return v == rv }
func (v Uint16) V() Uint16         { return v }

type Uint32 uint32

func (v Uint32) LT(rv Uint32) bool { return v < rv }
func (v Uint32) GT(rv Uint32) bool { return v > rv }
func (v Uint32) EQ(rv Uint32) bool { return v == rv }
func (v Uint32) V() Uint32         { return v }

type Uintptr uintptr

func (v Uintptr) LT(rv Uintptr) bool { return v < rv }
func (v Uintptr) GT(rv Uintptr) bool { return v > rv }
func (v Uintptr) EQ(rv Uintptr) bool { return v == rv }
func (v Uintptr) V() Uintptr         { return v }

type Uint64 uint64

func (v Uint64) LT(rv Uint64) bool { return v < rv }
func (v Uint64) GT(rv Uint64) bool { return v > rv }
func (v Uint64) EQ(rv Uint64) bool { return v == rv }
func (v Uint64) V() Uint64         { return v }

type Float32 float32

func (v Float32) LT(rv Float32) bool { return v < rv }
func (v Float32) GT(rv Float32) bool { return v > rv }
func (v Float32) EQ(rv Float32) bool { return v == rv }
func (v Float32) V() Float32         { return v }

type Float64 float64

func (v Float64) LT(rv Float64) bool { return v < rv }
func (v Float64) GT(rv Float64) bool { return v > rv }
func (v Float64) EQ(rv Float64) bool { return v == rv }
func (v Float64) V() Float64         { return v }

type String string

func (v String) LT(rv String) bool { return v < rv }
func (v String) GT(rv String) bool { return v > rv }
func (v String) EQ(rv String) bool { return v == rv }
func (v String) V() String         { return v }

// type Complex64 complex64

// func (v Complex64) LT(rv Complex64) bool { return v < rv }
// func (v Complex64) GT(rv Complex64) bool { return v > rv }
// func (v Complex64) EQ(rv Complex64) bool { return v == rv }

// type Complex128 complex128

// func (v Complex128) LT(rv Complex128) bool { return v < rv }
// func (v Complex128) GT(rv Complex128) bool { return v > rv }
// func (v Complex128) EQ(rv Complex128) bool { return v == rv }
