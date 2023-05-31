# go-stp

## map

### func

- Key[K comparable, V any](tm map[K]V) []K
    get key slice from a map
- NewMap[K comparable, V any]() *Map[K, V]
    make a Map struct
- NewCMap[K comparable, V any]() *CMap[K, V]
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

- CMap [K comparable, V any]
    - func (cm *CMap[K, V]) Get(k K) (V, bool)
    - func (cm *CMap[K, V]) Save(k K, v V) (int, bool)
    - func (cm *CMap[K, V]) Remove(k K) (V, int)
    
## slice

### func

- Compare
    compare any two slice

## convert

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

## channel

### func

- NewSharedChannel

create a shared channel

- NewSharedBufferChannel

create a shared channel with buffer

### struct

- SharedChannel[T any]
    - shared channel

### method

- SharedChannel[T any]
    - func (sc *SharedChannel[T]) Share() *SharedChannel[T]
    - func (sc *SharedChannel[T]) Get() chan T
    - func (sc *SharedChannel[T]) UseCount() int64
    - func (sc *SharedChannel[T]) Close()

## pool

### func

- NewPool[T any](c int) *Pool[T]
    - make an object pool with counter to allocate memory
    - in particular, any type that uses `Pool` needs to implement its null method, that is, the byte array element pointing to memory is 0, for example: `simpleStruct.Free`

### struct

- Pool[T any]
    - pool struct, holding memory byte

### method

- Pool[T any]
    - func (p *Pool[T]) Get() *T

## queue

> copy from $GOROOT/src/sync/poolqueue.go

### func

- func NewPoolDequeue(n int) PoolDequeue
- func NewPoolChain() PoolDequeue

### interface

- PoolDequeue
    - PushHead(val any) bool
    - PopHead() (any, bool)
    - PopTail() (any, bool)

### struct

- poolDequeue
- poolChain
