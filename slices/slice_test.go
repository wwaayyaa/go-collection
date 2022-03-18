package slices

import (
	"strconv"
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
}

func TestSlice_Length(t *testing.T) {
	expected := 3
	actual := NewSliceCollection([]int{1, 2, 3}).Length()
	assert.Equal(t, expected, actual)
}

func TestSlice_Get(t *testing.T) {
	expected := 33
	actual := NewSliceCollection([]int{1, 2, 33}).Get(2)
	assert.Equal(t, expected, actual)

	expected1 := "hello"
	actual1 := NewSliceCollection([]string{"a", "c", "zzzz"}).Get(4, "hello")
	assert.Equal(t, expected1, actual1)
}

func TestSlice_First(t *testing.T) {
	expected := 1
	actual := NewSliceCollection([]int{1, 2, 33}).First()
	assert.Equal(t, expected, actual)

	expected1 := "a"
	actual1 := NewSliceCollection([]string{"a", "c", "zzzz"}).First()
	assert.Equal(t, expected1, actual1)
}

func TestSlice_Find(t *testing.T) {
	expected := 2
	actual := NewSliceCollection([]int{1, 2, 33, 2}).Find(func(i int) bool { return i == 2 })
	assert.Equal(t, expected, actual)

	type People struct {
		Name string
	}
	people := []People{{Name: "jack"}, {Name: "bob"}}
	expected1 := people[1]
	actual1 := NewSliceCollection(people).Find(func(p People) bool { return p.Name == "bob" })
	assert.Equal(t, expected1, actual1)
}

func TestSlice_Index(t *testing.T) {
	type People struct {
		Name string
	}
	people := []People{{Name: "jack"}, {Name: "bob"}}
	expected1 := 0
	actual1 := NewSliceCollection(people).Index(func(p People) bool { return p.Name == "jack" })
	assert.Equal(t, expected1, actual1)
}

func TestSlice_Each(t *testing.T) {
	expected := 6
	actual := 0
	NewSliceCollection([]int{1, 2, 3}).Each(func(k int, v int) bool { actual += v; return true })
	assert.Equal(t, expected, actual)

	expected = 3
	actual = 0
	NewSliceCollection([]int{1, 2, 3}).Each(func(k int, v int) bool { actual += k; return true })
	assert.Equal(t, expected, actual)
}

func TestSlice_Map(t *testing.T) {
	expected := []int{2, 3, 4}
	actual := NewSliceCollection([]int{1, 2, 3}).Map(func(k int, v int) int { return v + 1 }).All()
	assert.ElementsMatch(t, expected, actual)

	expected1 := []string{"hello world", "hello girl"}
	actual1 := NewSliceCollection([]string{"world", "girl"}).Map(func(k int, v string) string { return "hello " + v }).All()
	assert.ElementsMatch(t, expected1, actual1)
}

func TestSlice_Transform(t *testing.T) {
	expected := []int{4, 9, 16}
	actual := NewSliceCollection([]int{1, 2, 3})
	actual.
		Transform(func(k int, v int) int { return v + 1 }).
		Transform(func(k int, v int) int { return v * v })
	assert.ElementsMatch(t, expected, actual.All())
}

func TestSlice_Contains(t *testing.T) {
	expected := true
	actual := NewSliceCollection([]string{"hello", "fine"}).Contains(func(v string) bool { return v == "fine" })
	assert.Equal(t, expected, actual)

	expected = false
	actual = NewSliceCollection([]string{"hello", "fine"}).Contains(func(v string) bool { return v == "fire" })
	assert.Equal(t, expected, actual)
}

func TestSlice_Filter(t *testing.T) {
	expected := []int{1, 2, 3}
	actual := NewSliceCollection([]int{1, 2, 3, 4, 5, 6}).
		Filter(func(x int) bool { return x <= 3 }).
		All()
	assert.ElementsMatch(t, expected, actual)

	expected1 := []string{"test"}
	actual1 := NewSliceCollection([]string{"asdasd", "123123", "test", "test123123"}).Filter(func(x string) bool { return x == "test" }).All()
	assert.ElementsMatch(t, expected1, actual1)

	type People struct {
		Name string
		Age  int
	}
	people := []People{{Name: "jack", Age: 12}, {Name: "bob", Age: 32}, {Name: "jack", Age: 23}}
	expected2 := people[2]
	actual2 := NewSliceCollection(people).
		Filter(func(x People) bool { return x.Name == "jack" && x.Age == 23 }).
		First()
	assert.Equal(t, expected2, actual2)
}

func TestSlice_Except(t *testing.T) {
	expected := []int{4, 5, 6}
	actual := NewSliceCollection([]int{1, 2, 3, 4, 5, 6}).
		Except(func(x int) bool { return x <= 3 }).
		All()
	assert.ElementsMatch(t, expected, actual)

	type People struct {
		Name string
		Age  int
	}
	people := []People{{Name: "jack", Age: 12}, {Name: "bob", Age: 32}, {Name: "jack", Age: 23}}
	expected2 := people[:2]
	actual2 := NewSliceCollection(people).
		Except(func(x People) bool { return x.Name == "jack" && x.Age == 23 }).
		All()
	assert.ElementsMatch(t, expected2, actual2)
}

func TestSlice_Join(t *testing.T) {
	expected := "4-5-6"
	actual := NewSliceCollection([]int{1, 2, 3, 4, 5, 6}).
		Except(func(x int) bool { return x <= 3 }).
		Join(func(v int) string { return strconv.Itoa(v) }, "-")
	assert.Equal(t, expected, actual)
}

func TestSlice_Clone(t *testing.T) {
	expected := []int{4, 5, 6}
	actual := NewSliceCollection(expected).
		Clone().
		All()
	assert.ElementsMatch(t, expected, actual)
}

func TestSlice_Tap(t *testing.T) {
	expected := 30
	actual := 0
	NewSliceCollection([]int{4, 5, 6}).
		Tap(func(nums SliceCollection[int]) {
			actual = nums.Get(1) * nums.Get(2)
		}).
		All()
	assert.Equal(t, expected, actual)
}

func TestSlice_ToJson(t *testing.T) {
	expected := "[2,3]"
	bytes, _ := NewSliceCollection([]int{2, 3}).ToJson()
	actual := string(bytes)
	assert.Equal(t, expected, actual)
}
