package avmap

import "golang.org/x/exp/constraints"

type Ints interface { // expanded constraints.Integer
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

type Values interface {
	constraints.Integer | constraints.Float | ~string
}

func SetIfMissing[K comparable, V any](m map[K]V, key K, value V) {
	if _, ok := m[key]; !ok {
		m[key] = value
	}
}

// Inc the value at the key, create the key if not exists. Integer values only
func Inc[K comparable, V Ints](m map[K]V, key K) {
	if _, ok := m[key]; !ok {
		var zero V
		m[key] = zero
	}
	m[key]++ // cannot be used with constraints.Integer values :|
}

// Add the value to the value at the key, create the key if not exists
func Add[K comparable, V Values](m map[K]V, key K, value V) {
	if _, ok := m[key]; !ok {
		var zero V
		m[key] = zero
	}
	m[key] += value
}

// Append to the []V key, create the key if not exists
func Append[K comparable, V Values](m map[K][]V, key K, value V) {
	if _, ok := m[key]; !ok {
		m[key] = []V{}
	}
	m[key] = append(m[key], value)
}
