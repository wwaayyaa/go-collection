package maps

import (
	"github.com/stretchr/testify/assert"
	go_collection "github.com/wwaayyaa/go-collection"
	"testing"
)

func TestMapCollection_All(t *testing.T) {
	expected := map[string]int{"a": 1}
	actual := NewMapCollection(expected).All()
	assert.Equal(t, expected, actual)
}

func TestMapCollection_Count(t *testing.T) {
	expected := 2
	actual := NewMapCollection(map[string]int{"a": 1, "c": 3}).Count()
	assert.Equal(t, expected, actual)
}

func TestMapCollection_Empty(t *testing.T) {
	expected := true
	actual := NewMapCollection(map[string]int{}).Empty()
	assert.Equal(t, expected, actual)

	expected = false
	actual = NewMapCollection(map[string]int{"a": 1}).Empty()
	assert.Equal(t, expected, actual)
}

func TestMapCollection_Keys(t *testing.T) {
	expected := []string{"a", "z"}
	actual := NewMapCollection(map[string]int{"a": 1, "z": 100}).Keys()
	assert.Equal(t, expected, actual)
}

func TestMapCollection_Values(t *testing.T) {
	expected := []int{1, 100}
	actual := NewMapCollection(map[string]int{"a": 1, "z": 100}).Values()
	assert.Equal(t, expected, actual)
}

func TestMapCollection_Entries(t *testing.T) {
	expected := []go_collection.Entry[string, int]{{"a", 1}, {"b", 2}}
	actual := NewMapCollection(map[string]int{"a": 1, "b": 2}).Entries()
	assert.ElementsMatch(t, expected, actual)
}

func TestMapCollection_FromEntries(t *testing.T) {
	expected := map[string]int{"a": 1, "b": 2}
	actual := NewMapCollection(map[string]int{}).FromEntries([]go_collection.Entry[string, int]{{"a", 1}, {"b", 2}}).All()
	assert.Equal(t, expected, actual)
}

func TestMapCollection_Has(t *testing.T) {
	expected := true
	actual := NewMapCollection(map[string]int{"a": 1, "z": 100}).Has("z")
	assert.Equal(t, expected, actual)

	expected = false
	actual = NewMapCollection(map[string]int{"a": 1, "z": 100}).Has("x")
	assert.Equal(t, expected, actual)
}

func TestMapCollection_Get(t *testing.T) {
	expected := 100
	actual, ok := NewMapCollection(map[string]int{"a": 1, "z": 100}).Get("z")
	assert.Equal(t, true, ok)
	assert.Equal(t, expected, actual)

	expected = 0
	actual, ok = NewMapCollection(map[string]int{"a": 1, "z": 100}).Get("x")
	assert.Equal(t, false, ok)
	assert.Equal(t, expected, actual)
}

func TestMapCollection_Put(t *testing.T) {
	expected := map[string]int{"a": 1, "b": 2}
	actual := NewMapCollection(map[string]int{"a": 1}).Put("b", 2).All()
	assert.Equal(t, expected, actual)
}
func TestMapCollection_Pull(t *testing.T) {
	expected := map[string]int{"b": 2}
	actual := NewMapCollection(map[string]int{"a": 1, "b": 2})
	v, ok := actual.Pull("a")
	assert.Equal(t, true, ok)
	assert.Equal(t, v, 1)
	assert.Equal(t, expected, actual.All())

	expected = map[string]int{"a": 1, "b": 2}
	actual = NewMapCollection(map[string]int{"a": 1, "b": 2})
	v, ok = actual.Pull("z")
	assert.Equal(t, false, ok)
	assert.Equal(t, v, 0)
	assert.Equal(t, expected, actual.All())
}
