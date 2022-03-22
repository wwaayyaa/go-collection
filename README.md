# go-collection

`go collection` is a tool implemented using generic, it can help you process slice/map data quickly and easily convert
between them.

Note: To use this project, you need to upgrade to go1.18 version. 
Some methods cannot be implemented in the 1.18 version, 
and will be supported after version 1.19 supports type methods.

## 🚀 Install

```go 
go get -u github.com/wwaayyaa/go-collection
```

## ✏️ Usage

```go 
import collect "github.com/wwaayyaa/go-collection"
```

### Slice example

```go
// filter  
NewSlice[int]([]int{1, 2, 3, 4, 5, 6}).
Filter(func (x int) bool { return x <= 3 }).
All() // [1,2,3]

//join
NewSlice[int]([]int{1, 2, 3, 4, 5, 6}).
Except(func (x int) bool { return x <= 3 }).
Join(func (v int) string { return strconv.Itoa(v) }, "-")
// 1-2-3

//map
NewSlice[string]([]string{"world", "girl"}).
Map(func (k int, v string) string { return "hello " + v }).
All() // ["hello world", "hello girl"]

```

### Maps example

```go
//["a", "z"]
NewMapCollection(map[string]int{"a": 1, "z": 100}).Keys()

//[1, 100]
NewMapCollection(map[string]int{"a": 1, "z": 100}).Values()
```

## 📖 API

### Slice

- `Len` 
- `First` 
- `Last` 
- `Get`
- `Put`
- `Prepend`
- `Shift`
- `Push`
- `Pop`
- `Find` 
- `Index` 
- `Each` 
- `Map` 
- `Transform` 
- `All` 
- `Contains` 
- `Filter` 
- `Reject` 
- `Concat` 
- `Join` 
- `Clone`
- `Tap`
- `ToJson`
- `Empty`
- `Diff` 
- `Merge` 
- `Reverse` 
- `Slice` 
- `Delete` 
- `Chunk` 
- `Uniq` 
- `Shuffle` 
- `Keys` 
- `Values` 
- `Only` 
- `Except` 
- `Reduce` 
- `FlatMap` 
- `GroupBy` 
- `KeyBy` 
- `Flatten`

### Map

- `All`
- `Count`
- `Empty`
- `Keys`
- `Values`
- `Entries`
- `FromEntries`
- `Has`
- `Get`
- `Put`
- `Pull`
- `Union`
- `Intersect`
- `Diff`
- `SymmetricDiff`

