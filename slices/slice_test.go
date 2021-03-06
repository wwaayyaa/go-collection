package slices

import (
	"strconv"
	"strings"
	"testing"
)
import "github.com/stretchr/testify/assert"

func TestNewSlice(t *testing.T) {
	expected := []int{1, 2, 3}
	actual := NewSliceCollection(expected).All()
	assert.ElementsMatch(t, expected, actual)

	expected1 := []string{"a", "c", "zzzz"}
	actual1 := NewSliceCollection(expected1).All()
	assert.ElementsMatch(t, expected1, actual1)

	var n []int
	assert.ElementsMatch(t, 0, NewSliceCollection(n).Len())
}

func TestSliceCollection_Length(t *testing.T) {
	expected := 3
	actual := NewSliceCollection([]int{1, 2, 3}).Len()
	assert.Equal(t, expected, actual)
}

func TestSliceCollection_Get(t *testing.T) {
	expected := 33
	actual, _ := NewSliceCollection([]int{1, 2, 33}).Get(2)
	assert.Equal(t, expected, actual)

	expected1 := "hello"
	actual1, ok := NewSliceCollection([]string{"a", "c", "zzzz"}).Get(4)
	if !ok {
		actual1 = "hello"
	}
	assert.Equal(t, expected1, actual1)
}

func TestSliceCollection_First(t *testing.T) {
	expected := 1
	actual, _ := NewSliceCollection([]int{1, 2, 33}).First()
	assert.Equal(t, expected, actual)

	expected1 := "a"
	actual1, _ := NewSliceCollection([]string{"a", "c", "zzzz"}).First()
	assert.Equal(t, expected1, actual1)
}

func TestSliceCollection_Last(t *testing.T) {
	expected := 3
	actual, _ := NewSliceCollection([]int{1, 2, 3}).Last()
	assert.Equal(t, expected, actual)
}

func TestSliceCollection_Put(t *testing.T) {
	expected := []int{1, 2, 3}
	actual := NewSliceCollection([]int{1, 1, 3}).Put(1, 2).All()
	assert.Equal(t, expected, actual)
}

func TestSliceCollection_Find(t *testing.T) {
	expected := 33
	actual, _ := NewSliceCollection([]int{1, 2, 33, 2}).Find(func(k, i int) bool { return i == 2 })
	assert.Equal(t, expected, actual)

	type People struct {
		Name string
	}
	people := []People{{Name: "jack"}, {Name: "bob"}}
	expected1 := people[1]
	actual1, _ := NewSliceCollection(people).Find(func(p People, _ int) bool { return p.Name == "bob" })
	assert.Equal(t, expected1, actual1)
}

func TestSliceCollection_Index(t *testing.T) {
	type People struct {
		Name string
	}
	people := []People{{Name: "jack"}, {Name: "bob"}}
	expected1 := 0
	actual1 := NewSliceCollection(people).Index(func(p People, _ int) bool { return p.Name == "jack" })
	assert.Equal(t, expected1, actual1)
}

func TestSliceCollection_Values(t *testing.T) {
	exptected := []int{1, 2, 3}
	actual := NewSliceCollection([]int{1, 2, 3}).Values()
	assert.ElementsMatch(t, exptected, actual)
}

func TestSliceCollection_Each(t *testing.T) {
	expected := 6
	actual := 0
	NewSliceCollection([]int{1, 2, 3}).Each(func(v int, i int) bool { actual += v; return true })
	assert.Equal(t, expected, actual)

	expected = 3
	actual = 0
	NewSliceCollection([]int{1, 2, 3}).Each(func(v int, i int) bool { actual += i; return true })
	assert.Equal(t, expected, actual)

	expected = 3
	actual = 0
	NewSliceCollection([]int{1, 2, 3}).Each(func(v int, i int) bool {
		if v > 2 {
			return false
		} else {
			actual += v
		}
		return true
	})
	assert.Equal(t, expected, actual)
}

func TestSliceCollection_Map(t *testing.T) {
	expected := []int{2, 3, 4}
	actual := NewSliceCollection([]int{1, 2, 3}).Map(func(v int, _ int) int { return v + 1 }).All()
	assert.Equal(t, expected, actual)

	expected1 := []string{"hello world", "hello girl"}
	actual1 := NewSliceCollection([]string{"world", "girl"}).Map(func(v string, _ int) string { return "hello " + v }).All()
	assert.Equal(t, expected1, actual1)
}

func TestSliceCollection_Transform(t *testing.T) {
	expected := []int{4, 9, 16}
	actual := NewSliceCollection([]int{1, 2, 3})
	actual.
		Transform(func(v int, _ int) int { return v + 1 }).
		Transform(func(v int, _ int) int { return v * v })
	assert.Equal(t, expected, actual.All())
}

func TestSliceCollection_Contains(t *testing.T) {
	expected := true
	actual := NewSliceCollection([]string{"hello", "fine"}).Contains(func(v string, _ int) bool { return v == "fine" })
	assert.Equal(t, expected, actual)

	expected = false
	actual = NewSliceCollection([]string{"hello", "fine"}).Contains(func(v string, _ int) bool { return v == "fire" })
	assert.Equal(t, expected, actual)
}

func TestSliceCollection_Empty(t *testing.T) {
	expected := true
	actual := NewSliceCollection([]string{}).Empty()
	assert.Equal(t, expected, actual)
}

func TestSliceCollection_Filter(t *testing.T) {
	expected := []int{1, 2, 3}
	actual := NewSliceCollection([]int{1, 2, 3, 4, 5, 6}).
		Filter(func(v, _ int) bool { return v <= 3 }).
		All()
	assert.ElementsMatch(t, expected, actual)

	expected1 := []string{"test"}
	actual1 := NewSliceCollection([]string{"asdasd", "123123", "test", "test123123"}).Filter(func(x string, _ int) bool { return x == "test" }).All()
	assert.ElementsMatch(t, expected1, actual1)

	type People struct {
		Name string
		Age  int
	}
	people := []People{{Name: "jack", Age: 12}, {Name: "bob", Age: 32}, {Name: "jack", Age: 23}}
	expected2 := people[2]
	actual2, _ := NewSliceCollection(people).
		Filter(func(p People, _ int) bool { return p.Name == "jack" && p.Age == 23 }).
		First()
	assert.Equal(t, expected2, actual2)
}

func TestSliceCollection_Except(t *testing.T) {
	expected := []int{4, 5, 6}
	actual := NewSliceCollection([]int{1, 2, 3, 4, 5, 6}).
		Reject(func(x int) bool { return x <= 3 }).
		All()
	assert.ElementsMatch(t, expected, actual)

	type People struct {
		Name string
		Age  int
	}
	people := []People{{Name: "jack", Age: 12}, {Name: "bob", Age: 32}, {Name: "jack", Age: 23}}
	expected2 := people[:2]
	actual2 := NewSliceCollection(people).
		Reject(func(x People) bool { return x.Name == "jack" && x.Age == 23 }).
		All()
	assert.ElementsMatch(t, expected2, actual2)
}

func TestSliceCollection_Join(t *testing.T) {
	expected := "4-5-6"
	actual := NewSliceCollection([]int{1, 2, 3, 4, 5, 6}).
		Reject(func(x int) bool { return x <= 3 }).
		Join(func(v int) string { return strconv.Itoa(v) }, "-")
	assert.Equal(t, expected, actual)
}

func TestSliceCollection_Clone(t *testing.T) {
	expected := []int{4, 5, 6}
	actual := NewSliceCollection(expected).
		Clone().
		All()
	assert.ElementsMatch(t, expected, actual)
}

func TestSliceCollection_Tap(t *testing.T) {
	expected := 3
	actual := 0
	NewSliceCollection([]int{4, 5, 6}).
		Tap(func(nums *SliceCollection[int]) {
			actual = nums.Len()
		}).
		All()
	assert.Equal(t, expected, actual)
}

func TestSliceCollection_ToJson(t *testing.T) {
	expected := "[2,3]"
	bytes, _ := NewSliceCollection([]int{2, 3}).ToJson()
	actual := string(bytes)
	assert.Equal(t, expected, actual)
}

func TestSliceCollection_Diff(t *testing.T) {
	expected := []int{1, 5}
	actual := NewSliceCollection([]int{1, 2, 3, 4, 5}).Diff([]int{2, 3, 4}).All()
	assert.Equal(t, expected, actual)
}

func TestSliceCollection_Merge(t *testing.T) {
	expected := []int{1, 2, 1, 3}
	actual := NewSliceCollection([]int{1}).Merge([]int{2, 1}).Merge([]int{3}).All()
	assert.Equal(t, expected, actual)
}
func TestSliceCollection_Pop(t *testing.T) {
	expected := []int{1, 2}
	actual := NewSliceCollection([]int{1, 2, 3})
	l, ok := actual.Pop()
	assert.Equal(t, expected, actual.All())
	assert.Equal(t, 3, l)
	assert.Equal(t, true, ok)

	expected1 := 0
	actual = NewSliceCollection([]int{})
	l, ok = actual.Pop()
	assert.Equal(t, expected1, l)
	assert.Equal(t, false, ok)
}

func TestSliceCollection_Push(t *testing.T) {
	expected := []int{1, 2}
	actual := NewSliceCollection([]int{1}).Push(2).All()
	assert.Equal(t, expected, actual)
}
func TestSliceCollection_Shift(t *testing.T) {
	co := NewSliceCollection([]int{1, 2, 3})
	actual, ok := co.Shift()
	assert.Equal(t, true, ok)
	assert.Equal(t, []int{2, 3}, co.All())
	assert.Equal(t, 1, actual)

	actual, ok = NewSliceCollection([]int{}).Shift()
	assert.Equal(t, false, ok)
	assert.Equal(t, 0, actual)
}

func TestSliceCollection_Delete(t *testing.T) {
	expected := []int{5, 7, 8}
	actual := NewSliceCollection([]int{5, 6, 7, 8}).Delete(1).All()
	assert.Equal(t, expected, actual)
}

func TestSliceCollection_Reverse(t *testing.T) {
	expected := []int{4, 3, 2, 1}
	actual := NewSliceCollection([]int{1, 2, 3, 4}).Reverse().All()
	assert.Equal(t, expected, actual)

	expected = []int{3, 2, 1}
	actual = NewSliceCollection([]int{1, 2, 3}).Reverse().All()
	assert.Equal(t, expected, actual)
}

func TestSliceCollection_Slice(t *testing.T) {
	expected := []int{2, 3}
	actual := NewSliceCollection([]int{1, 2, 3, 4}).Slice(1, 2).All()
	assert.Equal(t, expected, actual)

	expected = []int{2, 3, 4}
	actual = NewSliceCollection([]int{1, 2, 3, 4}).Slice(1).All()
	actual1 := NewSliceCollection([]int{1, 2, 3, 4}).Slice(1, -1).All()
	assert.Equal(t, expected, actual)
	assert.Equal(t, expected, actual1)
}

func TestSliceCollection_Prepend(t *testing.T) {
	expected := []int{0, 1, 2, 3}
	actual := NewSliceCollection([]int{1, 2, 3}).Prepend(0).All()
	assert.Equal(t, expected, actual)
}

func TestReduce(t *testing.T) {
	expected := "let go 123"
	actual := Reduce[int, string](NewSliceCollection([]int{1, 2, 3}).All(), func(v int, last string, _ int) string {
		return last + strconv.Itoa(v)
	}, "let go ")
	assert.Equal(t, expected, actual)
}

func TestSliceCollection_Chunk(t *testing.T) {
	expected := [][]int{{0, 1}, {2, 3}, {4}}
	actual := NewSliceCollection([]int{0, 1, 2, 3, 4}).Chunk(2)
	assert.Equal(t, expected, actual)
}

func TestSliceCollection_Uniq(t *testing.T) {
	expected := []int{1, 2, 3, 4}
	actual := NewSliceCollection([]int{1, 1, 2, 2, 3, 4, 1, 2, 3}).Uniq().All()
	assert.Equal(t, expected, actual)
}

func TestGroupBy(t *testing.T) {
	expected := map[string][]int{"yes": {0, 2}, "no": {1, 3}}
	actual := GroupBy([]int{0, 1, 2, 3}, func(t int, i int) string {
		if t%2 == 0 {
			return "yes"
		}
		return "no"
	})
	assert.Equal(t, expected, actual)
}

func TestSliceCollection_Shuffle(t *testing.T) {
	assert.NotEqual(t, []int{1, 2, 3, 4, 5, 6, 7}, NewSliceCollection([]int{1, 2, 3, 4, 5, 6, 7}).Shuffle().All())
}

func TestSliceCollection_Concat(t *testing.T) {
	expected := []int{1, 2, 3, 4}
	actual := NewSliceCollection([]int{1, 2}).Concat([]int{3, 4}).All()
	assert.Equal(t, expected, actual)
}

func TestSliceCollection_OnlyAndExcept(t *testing.T) {
	expected := []int{1, 4}
	data := NewSliceCollection([]int{1, 2, 3, 4})
	assert.Equal(t, expected, data.Only([]int{0, 3}).All())

	expected = []int{2, 3}
	assert.Equal(t, expected, data.Except([]int{0, 3}).All())
}

func TestKeyBy(t *testing.T) {
	expected := map[string]string{"a": "abc", "b": "book"}
	actual := KeyBy([]string{"abc", "book"}, func(s string) string {
		return s[0:1]
	})
	assert.Equal(t, expected, actual)
}

func TestFlatten(t *testing.T) {
	expected := []string{"a", "b", "c", "d"}
	actual := Flatten([][]string{{"a", "b"}, {"c", "d"}})
	assert.Equal(t, expected, actual)
}

func TestFlatMap(t *testing.T) {
	expected := []int{11, 12, 2, 34, 4}
	actual := FlatMap([]string{"11.12", "2", "34.4"}, func(v string, _ int) (ret []int) {
		for _, s := range strings.Split(v, ".") {
			n, _ := strconv.Atoi(s)
			ret = append(ret, n)
		}
		return ret
	})
	assert.Equal(t, expected, actual)
}
