package utils

/*
 A generic Map interface
*/
type Map[K comparable, V interface{}] interface {
	Put(key K, value V)
	Get(key K) V
	Remove(key K)
	Keys() []K
	Values() []V
	ContainsKey(key K) bool
	Merge(key K, newValue V, mergeOp func(V, V) V)
	Size() int
}

/*
 An implementation of the Map interface
 Works as a wrapper around the native golang map for a more Object-Oriented style of coding
*/
type MapWrapper[K comparable, V interface{}] struct {
	items map[K]V
}

/*
 O(1)
 Instantiates a new MapWrapper that is empty
*/
func NewMapWrapper[K comparable, V interface{}]() *MapWrapper[K, V] {
	m := MapWrapper[K, V]{}
	m.items = make(map[K]V)
	return &m
}

/*
 O(1)
 Assumes: MapWrapper m has been instantiated
 Ensures: the key is mapped to the value
*/
func (m *MapWrapper[K, V]) Put(key K, value V) {
	m.items[key] = value
}

/*
 O(1)
 Assumes: MapWrapper m has been instantiated
 Returns the value that is associated with the supplied key
 If there is no mapping, the function will return the nil value of the value type
*/
func (m MapWrapper[K, V]) Get(key K) V {
	return m.items[key]
}

/*
 O(1)
 Assumes: MapWrapper m has been instantiated
 Removes the mapping with the supplied key
 No-op if such a mapping doesn't exist
*/
func (m *MapWrapper[K, V]) Remove(key K) {
	delete(m.items, key)
}

/*
 O(n)
 Assumes: MapWrapper m has been instantiated
 Returns the MapWrappers keys as a slice
*/
func (m MapWrapper[K, V]) Keys() []K {
	keys := make([]K, 0, len(m.items))
	for k := range m.items {
		keys = append(keys, k)
	}
	return keys
}

/*
 O(n)
 Assumes: MapWrapper m has been instantiated
 Returns the MapWrappers values as a slice
*/
func (m MapWrapper[K, V]) Values() []V {
	vals := make([]V, 0)
	for _, v := range m.items {
		vals = append(vals, v)
	}
	return vals
}

/*
 O(1)
 Assumes: MapWrapper m has been instantiated
 Returns true if there exist a mapping with the supplied key
*/
func (m MapWrapper[K, V]) ContainsKey(key K) bool {
	_, exists := m.items[key]
	return exists
}

/*
 O(1)
 Assumes: MapWrapper m has been instantiated
 If there exists a mapping with the supplied key, the new value will be mergeOp(newVal, oldVal)
 If not, it works as a regular Put
*/
func (m *MapWrapper[K, V]) Merge(key K, newVal V, mergeOp func(V, V) V) {
	if m.ContainsKey(key) {
		m.items[key] = mergeOp(newVal, m.items[key])
	} else {
		m.Put(key, newVal)
	}
}

/*
 Assumes: MapWrapper m has been instantiated
 Returns the number of mappings
*/
func (m MapWrapper[K, V]) Size() int {
	return len(m.items)
}
