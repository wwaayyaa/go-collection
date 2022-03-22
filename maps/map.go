package maps

import (
	go_collection "github.com/wwaayyaa/go-collection"
)

type MapCollection[T ~map[K]V, K comparable, V any] struct {
	items T
}

func NewMapCollection[T ~map[K]V, K comparable, V any](v T) *MapCollection[T, K, V] {
	return &MapCollection[T, K, V]{items: v}
}

func (co *MapCollection[T, K, V]) All() T {
	return co.items
}
func (co *MapCollection[T, K, V]) Count() int {
	return len(co.items)
}
func (co *MapCollection[T, K, V]) Empty() bool {
	return co.Count() == 0
}
func (co *MapCollection[T, K, V]) Keys() (keys []K) {
	for k, _ := range co.items {
		keys = append(keys, k)
	}
	return keys
}
func (co *MapCollection[T, K, V]) Values() (values []V) {
	for _, v := range co.items {
		values = append(values, v)
	}
	return values
}

func (co *MapCollection[T, K, V]) Entries() []go_collection.Entry[K, V] {
	ret := make([]go_collection.Entry[K, V], 0, co.Count())

	for k, v := range co.items {
		ret = append(ret, go_collection.Entry[K, V]{k, v})
	}

	return ret
}

func (co *MapCollection[T, K, V]) FromEntries(entries []go_collection.Entry[K, V]) *MapCollection[map[K]V, K, V] {
	ret := map[K]V{}
	for _, e := range entries {
		ret[e.Key] = e.Value
	}
	return NewMapCollection(ret)
}

func (co *MapCollection[T, K, V]) Has(key K) bool {
	if _, ok := co.items[key]; ok {
		return true
	} else {
		return false
	}
}

func (co *MapCollection[T, K, V]) Get(key K) (value V, _ bool) {
	if v, ok := co.items[key]; ok {
		return v, true
	} else {
		return value, false
	}
}
func (co *MapCollection[T, K, V]) Put(key K, value V) *MapCollection[T, K, V] {
	co.items[key] = value
	return co
}

func (co *MapCollection[T, K, V]) Pull(key K) (v V, _ bool) {
	if v, ok := co.items[key]; ok {
		delete(co.items, key)
		return v, true
	}
	return
}
