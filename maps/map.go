package maps

import (
	go_collection "github.com/wwaayyaa/go-collection"
)

type MapCollection[K comparable, V any] struct {
	items map[K]V
}

func NewMapCollection[K comparable, V any](v map[K]V) *MapCollection[K, V] {
	return &MapCollection[K, V]{items: v}
}

func (co *MapCollection[K, V]) All() map[K]V {
	return co.items
}

func (co *MapCollection[K, V]) Count() int {
	return len(co.items)
}

func (co *MapCollection[K, V]) Empty() bool {
	return co.Count() == 0
}

func (co *MapCollection[K, V]) Keys() (keys []K) {
	for k, _ := range co.items {
		keys = append(keys, k)
	}
	return keys
}

func (co *MapCollection[K, V]) Values() (values []V) {
	for _, v := range co.items {
		values = append(values, v)
	}
	return values
}

func (co *MapCollection[K, V]) Entries() []go_collection.Entry[K, V] {
	ret := make([]go_collection.Entry[K, V], 0, co.Count())

	for k, v := range co.items {
		ret = append(ret, go_collection.Entry[K, V]{k, v})
	}

	return ret
}

func (co *MapCollection[K, V]) FromEntries(entries []go_collection.Entry[K, V]) *MapCollection[K, V] {
	ret := map[K]V{}
	for _, e := range entries {
		ret[e.Key] = e.Value
	}
	return NewMapCollection(ret)
}

func (co *MapCollection[K, V]) Has(key K) bool {
	if _, ok := co.items[key]; ok {
		return true
	} else {
		return false
	}
}

func (co *MapCollection[K, V]) Get(key K) (value V, _ bool) {
	if v, ok := co.items[key]; ok {
		return v, true
	} else {
		return value, false
	}
}

func (co *MapCollection[K, V]) Put(key K, value V) *MapCollection[K, V] {
	co.items[key] = value
	return co
}

func (co *MapCollection[K, V]) Pull(key K) (v V, _ bool) {
	if v, ok := co.items[key]; ok {
		delete(co.items, key)
		return v, true
	}
	return
}

func (co *MapCollection[K, V]) Union(items map[K]V) *MapCollection[K, V] {
	for k, v := range items {
		co.items[k] = v
	}

	return co
}

func (co *MapCollection[K, V]) Intersect(items map[K]V) *MapCollection[K, V] {
	ret := NewMapCollection(map[K]V{})
	for k := range co.items {
		if v, ok := items[k]; ok {
			ret.items[k] = v
		}
	}
	return ret
}

func (co *MapCollection[K, V]) Diff(items map[K]V) *MapCollection[K, V] {
	ret := NewMapCollection(map[K]V{})
	for k, v := range co.items {
		if _, ok := items[k]; !ok {
			ret.items[k] = v
		}
	}
	return ret
}

func (co *MapCollection[K, V]) SymmetricDiff(items map[K]V) *MapCollection[K, V] {
	return co.Diff(items).Union(NewMapCollection(items).Diff(co.items).All())
}
