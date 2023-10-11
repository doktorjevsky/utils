package utils

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

func NewMapWrapper[K comparable, V interface{}]() *MapWrapper[K, V] {
	m := MapWrapper[K, V]{}
	m.items = make(map[K]V)
	return &m
}

type MapWrapper[K comparable, V interface{}] struct {
	items map[K]V
}

func (m *MapWrapper[K, V]) Put(key K, value V) {
	m.items[key] = value
}

func (m MapWrapper[K, V]) Get(key K) V {
	return m.items[key]
}

func (m *MapWrapper[K, V]) Remove(key K) {
	delete(m.items, key)
}

func (m MapWrapper[K, V]) Keys() []K {
	keys := make([]K, 0, len(m.items))
	for k := range m.items {
		keys = append(keys, k)
	}
	return keys
}

func (m MapWrapper[K, V]) Values() []V {
	vals := make([]V, 0)
	for _, v := range m.items {
		vals = append(vals, v)
	}
	return vals
}

func (m MapWrapper[K, V]) ContainsKey(key K) bool {
	_, exists := m.items[key]
	return exists
}

func (m *MapWrapper[K, V]) Merge(key K, newVal V, mergeOp func(V, V) V) {
	if m.ContainsKey(key) {
		m.items[key] = mergeOp(newVal, m.items[key])
	} else {
		m.Put(key, newVal)
	}
}

func (m MapWrapper[K, V]) Size() int {
	return len(m.items)
}
