package utils

type Map[K comparable, V any] interface {
	New() Map[K, V]
	Put(key K, value V)
	Get(key K) V
	Remove(key K)
	Keys() []K
	Values() []V
	Merge(key K, newValue V, binOp func(V, V) V)
	ContainsKey(key K) bool
}
