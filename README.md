# go-stp

## map

### func

- Key
    get key slice from a map
- NewMap
    make a Map struct

### struct

- Map [K comparable, V any]
    - a shrinkable map

### method

- Map [K comparable, V any]
    - func (m *Map[K, V]) Set(k K, v V)
    - func (m *Map[K, V]) Del(k K)
    - func (m Map[K, V]) Get(k K) (V, bool)
    - func (m Map[K, V]) Key() []K
    - func (m Map[K, V]) Range(f func(K, V) bool)
    
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