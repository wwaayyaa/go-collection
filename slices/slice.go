package slices

import (
	"encoding/json"
	"github.com/google/go-cmp/cmp"
	"strings"
)

/*
TODO
The following functions are achievable and will be updated soon：
  forget delete prepend concat put shift sortBy sortByDesc slice splice unique
  sum avg max min median

The following features require version 1.19 to allow methods to have type parameters. Because most of them return arbitrary types on demand.
  Example:
	func (co *SliceCollection[T]) Reduce[R any](fn func(R, T) R, R) R{
		* This R is the type we want
	}

  when groupBy keyBy reduce
  split chunk

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

func (co *SliceCollection[T]) Find(fn func(int, T) bool) (ret T, _ bool) {
	for k, v := range co.items {
		if fn(k, v) {
			return v, true
		}
	}
	return ret, false
}

func (co *SliceCollection[T]) Index(fn func(int, T) bool) int {
	for k, v := range co.items {
		if fn(k, v) {
			return k
		}
	}
	return -1
}

//Each cannot change self
func (co *SliceCollection[T]) Each(fn func(int, T) bool) *SliceCollection[T] {
	for k, v := range co.items {
		if !fn(k, v) {
			break
		}
	}
	return co
}

func (co *SliceCollection[T]) Map(fn func(int, T) T) *SliceCollection[T] {
	var ret []T
	for k, v := range co.items {
		ret = append(ret, fn(k, v))
	}
	return NewSliceCollection(ret)
}

func (co *SliceCollection[T]) Transform(fn func(int, T) T) *SliceCollection[T] {
	for k, v := range co.items {
		co.items[k] = fn(k, v)
	}
	return co
}

func (co *SliceCollection[T]) All() []T {
	return co.items
}

func (co *SliceCollection[T]) Contains(fn func(int, T) bool) bool {
	return co.Index(fn) != -1
}

func (co *SliceCollection[T]) Filter(fn func(T) bool) *SliceCollection[T] {
	var ret []T
	for _, v := range co.items {
		if fn(v) {
			ret = append(ret, v)
		}
	}
	return NewSliceCollection(ret)
}
func (co *SliceCollection[T]) Except(fn func(T) bool) *SliceCollection[T] {
	var ret []T
	for _, v := range co.items {
		if !fn(v) {
			ret = append(ret, v)
		}
	}
	return NewSliceCollection(ret)
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
	co.Each(func(i int, v T) bool {
		if _, ok := t.Find(func(_ int, _v T) bool { return cmp.Equal(_v, v) }); !ok {
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

//TODO unique Shuffle Split Splice Reduce
