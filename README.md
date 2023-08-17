# go-stp

## stp map

> updated at 2023.08.17

- map.go
- concurrency_map.go

### func

- Key\[K comparable, V any\](tm map[K]V) []K
    get key slice from a map
- NewMap\[K comparable, V any\]() *Map[K, V]
    make a Map struct
- NewCMap\[K comparable, V any\]() *CMap[K, V]
    make a CMap struct

### struct

- Map [K comparable, V any]
    - a shrinkable map
- CMap [K comparable, V any]
    - a map has a sync.RWMutex

### method

- Map [K comparable, V any]
    - func (m *Map[K, V]) Set(k K, v V)
    - func (m *Map[K, V]) Del(k K)
    - func (m Map[K, V]) Get(k K) (V, bool)
    - func (m Map[K, V]) Key() []K
    - func (m Map[K, V]) Range(f func(K, V) bool)
    - func (m Map[K, V]) Len() int

- CMap [K comparable, V any]
    - func (cm *CMap[K, V]) Get(k K) (V, bool)
    - func (cm *CMap[K, V]) Save(k K, v V) (int, bool)
    - func (cm *CMap[K, V]) Remove(k K) (V, int)
    
## stp slice

> updated at 2023.08.17

- slice.go

### func

- Compare
    compare any two slice
- NewArray
    make a JS-style Array struct

### method

- Array [T comparable]
    - func (a *Array[T]) Slice() []T
    - func (a *Array[T]) Filter(f func(v T, i int) bool) *Array[T]
    - func (a *Array[T]) Find(f func(v T, i int) bool) T
    - func (a *Array[T]) ForEach(f func(v T, i int))
    - func (a *Array[T]) Includes(v T) bool
    - func (a *Array[T]) IndexOf(v T) int
    - func (a *Array[T]) Map(f func(v T, i int) T) *Array[T]
    - func (a *Array[T]) Push(vs ...T) int
    - func (a *Array[T]) Shift() T
    - func (a *Array[T]) Splice(i, m int) *Array[T]

## stp string

> updated at 2023.08.17

- string.go

### func

- ConvertStringToStringStruct

split a string by splitter, then convert it to a struct, which the first several members's type must be string, eg,
```go
var s string = "I am a boy,You are a girl,We are human"
type _s struct {
    s1 string
    s2 string
    s3 string
    // ...
}

// &_s{s1: "I am a boy", s2: "You are a girl", s3: "We are human"}
fmt.Println(ConvertStringToStringStruct[_s](s, ","))
```

## stp channel

> updated at 2023.08.17

- channel.go

### func

- NewSharedChannel
    make a shared channel
- NewSharedBufferChannel
    make a shared channel with buffer

### struct

- SharedChannel [T any]
    - shared channel

### method

- SharedChannel [T any]
    - func (sc *SharedChannel[T]) Share() *SharedChannel[T]
    - func (sc *SharedChannel[T]) Get() chan T
    - func (sc *SharedChannel[T]) UseCount() int64
    - func (sc *SharedChannel[T]) Close()

## stp pool

> updated at 2023.08.17

- pool.go

### func

- NewPool[T any](c int) *Pool[T]
    - make an object pool with counter to allocate memory
    - in particular, any type that uses `Pool` needs to implement its null method, that is, the byte array element pointing to memory is 0, for example: `simpleStruct.Free`

### struct

- Pool [T any]
    - pool struct, holding memory byte

### method

- Pool [T any]
    - func (p *Pool[T]) Get() *T

## stp queue

> copy from $GOROOT/src/sync/poolqueue.go and make it support multi producers

> updated at 2023.08.17

- queue.go
- poolqueue.go

### func

- func NewPoolDequeue(n int) PoolDequeue
- func NewPoolChain() PoolDequeue

### interface

- PoolDequeue
    - PushHead(val any) bool
    - PopHead() (any, bool)
    - PopTail() (any, bool)

## stp reflect

> updated at 2023.08.17

- reflect.go

### func

- ReflectStructFieldKeyIndexMap(rt reflect.Type, tagKey string) map[string]int
- ReflectStructFieldKeyOffsetMap(rt reflect.Type, tagKey string) map[string]uintptr
- ReflectStructValue[T any](fs []string, v []any, tagKey string) *T

reflect struct field value slice to struct

```go
type T struct {
    I int     `json:"int"`
    F float64 `json:"float64"`
    S string  `json:"string"`
}
var fs []string = []string{"int", "float64", "string"}
var v []any = []any{1024, 0.618, "this is gold ratio"}


// &T{I: 1024, F: 0.618, S: "this is gold ratio"}
fmt.Println(ReflectStructValue[T](fs, v, "json"))
```
- ReflectStructValueSlice[T any](fs []string, vs [][]any, tagKey string) []*T
