package slices

import (
	"encoding/json"
	"strings"
)

/*
TODO
The following functions are achievable and will be updated soon：
  diff forget insert delete pop prepend push pull concat put shift sortBy sortByDesc slice splice unique
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

func (co *SliceCollection[T]) Length() int {
	return len(co.items)
}

func (co *SliceCollection[T]) Get(i int, _default ...T) (ret T) {
	if co.Length() < i && len(_default) > 0 {
		return _default[0]
	}
	return co.items[i]
}

func (co *SliceCollection[T]) First() (ret T) {
	return co.Get(0, ret)
}

func (co *SliceCollection[T]) Find(fn func(T) bool) (ret T) {
	for _, v := range co.items {
		if fn(v) {
			return v
		}
	}
	return ret
}

func (co *SliceCollection[T]) Index(fn func(T) bool) int {
	for k, v := range co.items {
		if fn(v) {
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

func (co *SliceCollection[T]) Contains(fn func(T) bool) bool {
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
	clone := make([]T, co.Length())
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
