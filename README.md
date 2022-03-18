# go-collection

`go collection` is a tool implemented using generic, it can help you process slice/map data quickly and easily convert
between them.

### Installation

```go 
go get -u github.com/wwaayyaa/go-collection
```

Then import it

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

Coming soon

