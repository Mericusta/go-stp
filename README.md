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

- Compare
    compare any two slice