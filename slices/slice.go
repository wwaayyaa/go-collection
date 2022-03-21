package slices

import (
	"encoding/json"
	"github.com/google/go-cmp/cmp"
	"math/rand"
	"strings"
)

/*
TODO
The following functions are achievable and will be updated soon：
  forget delete concat sortBy sortByDesc splice pluck
  sum avg max min median

The following features require version 1.19 to allow methods to have type parameters. Because most of them return arbitrary types on demand.
  Example:
	func (co *SliceCollection[T]) Reduce[R any](fn func(R, T) R, R) R{
		* This R is the type we want
	}

  when  keyBy
*/

type SliceCollection[T any] struct {
	items []T
}

func NewSliceCollection[T any](v []T) *SliceCollection[T] {
	if v == nil {
		return &SliceCollection[T]{}
	}

	clone := make([]T, len(v))
	copy(clone, v)
	return &SliceCollection[T]{items: clone}
}

func (co *SliceCollection[T]) Len() int {
	return len(co.items)
}

func (co *SliceCollection[T]) Get(i int) (ret T, _ bool) {
	if i >= co.Len() {
		return ret, false
	}
	return co.items[i], true
}

func (co *SliceCollection[T]) First() (ret T, _ bool) {
	return co.Get(0)
}

func (co *SliceCollection[T]) Last() (ret T, _ bool) {
	return co.Get(co.Len() - 1)
}

func (co *SliceCollection[T]) Put(i int, v T) *SliceCollection[T] {
	co.items[i] = v
	return co
}

func (co *SliceCollection[T]) Find(fn func(T, int) bool) (ret T, _ bool) {
	for i, v := range co.items {
		if fn(v, i) {
			return v, true
		}
	}
	return ret, false
}

func (co *SliceCollection[T]) Index(fn func(T, int) bool) int {
	for i, v := range co.items {
		if fn(v, i) {
			return i
		}
	}
	return -1
}

//Each cannot change self
func (co *SliceCollection[T]) Each(fn func(T, int) bool) *SliceCollection[T] {
	for i, v := range co.items {
		if !fn(v, i) {
			break
		}
	}
	return co
}

func (co *SliceCollection[T]) Map(fn func(T, int) T) *SliceCollection[T] {
	ret := make([]T, co.Len())
	for i, v := range co.items {
		ret = append(ret, fn(v, i))
	}
	return NewSliceCollection(ret)
}

func (co *SliceCollection[T]) Transform(fn func(T, int) T) *SliceCollection[T] {
	for i, v := range co.items {
		co.items[i] = fn(v, i)
	}
	return co
}

func (co *SliceCollection[T]) All() []T {
	return co.items
}

func (co *SliceCollection[T]) Contains(fn func(T, int) bool) bool {
	return co.Index(fn) != -1
}

func (co *SliceCollection[T]) Filter(fn func(T, int) bool) *SliceCollection[T] {
	var ret []T
	for i, v := range co.items {
		if fn(v, i) {
			ret = append(ret, v)
		}
	}
	return NewSliceCollection(ret)
}
func (co *SliceCollection[T]) Reject(fn func(T) bool) *SliceCollection[T] {
	var ret []T
	for _, v := range co.items {
		if !fn(v) {
			ret = append(ret, v)
		}
	}
	return NewSliceCollection(ret)
}

func (co *SliceCollection[T]) Concat(items []T) *SliceCollection[T] {
	co.items = append(co.items, items...)
	return co
}

func (co *SliceCollection[T]) Join(fn func(T) string, sep string) string {
	var str []string
	for _, v := range co.items {
		str = append(str, fn(v))
	}
	return strings.Join(str, sep)
}

func (co *SliceCollection[T]) Clone() *SliceCollection[T] {
	clone := make([]T, co.Len())
	copy(clone, co.items)
	return NewSliceCollection(clone)
}

func (co *SliceCollection[T]) Tap(fn func(*SliceCollection[T])) *SliceCollection[T] {
	fn(co.Clone())
	return co
}

func (co *SliceCollection[T]) ToJson() ([]byte, error) {
	return json.Marshal(co.items)
}

func (co *SliceCollection[T]) Empty() bool {
	return co.Len() == 0
}

func (co *SliceCollection[T]) Diff(target []T) *SliceCollection[T] {
	var different []T
	t := NewSliceCollection(target)
	co.Each(func(v T, i int) bool {
		if _, ok := t.Find(func(_v T, _ int) bool { return cmp.Equal(_v, v) }); !ok {
			different = append(different, v)
		}
		return true
	})
	return NewSliceCollection(different)
}

func (co *SliceCollection[T]) Merge(targets ...[]T) *SliceCollection[T] {
	for _, target := range targets {
		co.items = append(co.items, target...)
	}
	return co
}

func (co *SliceCollection[T]) Pop() (T, bool) {
	l := co.Len()
	if l == 0 {
		var t T
		return t, false
	}
	value := co.items[l-1]
	co.items = co.items[:l-1]
	return value, true
}

func (co *SliceCollection[T]) Push(v T) *SliceCollection[T] {
	co.items = append(co.items, v)
	return co
}

func (co *SliceCollection[T]) Reverse() *SliceCollection[T] {
	for i, j := co.Len()-1, 0; i > j; i, j = i-1, j+1 {
		co.items[i], co.items[j] = co.items[j], co.items[i]
	}
	return co
}

func (co *SliceCollection[T]) Slice(offset int, length ...int) *SliceCollection[T] {
	var ret []T
	if len(length) == 0 || (len(length) > 0 && length[0] == -1) {
		ret = co.items[offset:]
		return NewSliceCollection(ret)
	}
	return NewSliceCollection(co.items[offset:(offset + length[0])])
}

func (co *SliceCollection[T]) Prepend(v T) *SliceCollection[T] {
	co.items = append([]T{v}, co.items...)
	return co
}

func (co *SliceCollection[T]) Shift() (T, bool) {
	v, ok := co.Get(0)
	if !ok {
		return v, false
	}
	co.items = co.items[1:]
	return v, true
}

func (co *SliceCollection[T]) Delete(i int) *SliceCollection[T] {
	co.items = append(co.items[:i], co.items[i+1:]...)
	return co
}

func (co *SliceCollection[T]) Chunk(n int) (ret [][]T) {
	i := 1
	var chunk []T
	for _, v := range co.items {
		chunk = append(chunk, v)
		if i += 1; i > n {
			i, ret = 1, append(ret, chunk)
			chunk = []T{}
		}
	}
	if len(chunk) > 0 {
		ret = append(ret, chunk)
	}

	return ret
}

func (co *SliceCollection[T]) Uniq() *SliceCollection[T] {
	ret := NewSliceCollection([]T{})

	for i, v := range co.items {
		if !ret.Contains(func(r T, _ int) bool {
			return cmp.Equal(r, v)
		}) {
			ret.Push(v)
			continue
		}

		isEqual := false
		for j, w := range co.items {
			if i != j && cmp.Equal(v, w) {
				isEqual = true
			}
		}
		if !isEqual {
			ret.Push(v)
		}
	}
	return ret
}

func (co *SliceCollection[T]) Shuffle() *SliceCollection[T] {
	rand.Shuffle(co.Len(), func(i, j int) {
		co.items[i], co.items[j] = co.items[j], co.items[i]
	})

	return co
}

func (co *SliceCollection[T]) Keys() []int {
	var ret []int
	for k, _ := range co.items {
		ret = append(ret, k)
	}
	return ret
}

func (co *SliceCollection[T]) Values() []T {
	var ret []T
	for _, v := range co.items {
		ret = append(ret, v)
	}
	return ret
}

func (co *SliceCollection[T]) Only(keys []int) *SliceCollection[T] {
	var ret []T
	for _, key := range keys {
		if v, ok := co.Get(key); ok {
			ret = append(ret, v)
		}
	}
	return NewSliceCollection(ret)
}

func (co *SliceCollection[T]) Except(keys []int) *SliceCollection[T] {
	var ret []T
	for _, key := range NewSliceCollection(co.Keys()).Diff(keys).All() {
		if v, ok := co.Get(key); ok {
			ret = append(ret, v)
		}
	}
	return NewSliceCollection(ret)
}

// 1.18 not allow type parameters in methods
// In order to increase flexibility, return any type, so it can only be a function independently.
// https://github.com/golang/go/issues/49085

func Reduce[T, R any](d []T, h func(T, R, int) R, init R) R {
	for i, v := range d {
		init = h(v, init, i)
	}
	return init
}

func FlatMap[T, R any](items []T, it func(T, int) []R) []R {
	var result []R

	for i, v := range items {
		result = append(result, it(v, i)...)
	}

	return result
}

func GroupBy[T any, U comparable](items []T, it func(T, int) U) map[U][]T {
	result := map[U][]T{}

	for i, item := range items {
		key := it(item, i)

		if _, ok := result[key]; !ok {
			result[key] = []T{}
		}

		result[key] = append(result[key], item)
	}

	return result
}

func KeyBy[K comparable, V any](items []V, fn func(V) K) map[K]V {
	result := make(map[K]V, len(items))

	for _, v := range items {
		k := fn(v)
		result[k] = v
	}

	return result
}

func Flatten[T any](items [][]T) []T {
	var result []T

	for _, item := range items {
		result = append(result, item...)
	}

	return result
}
