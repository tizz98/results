# Go Results

Similar `Result` structs like in Rust or other languages.

See [`example_test.go`](example_test.go) for usage.

## Struct example

```go
//go:generate go run github.com/tizz98/results/cmd -pkg foo -t *Bar -tup-default nil -result-name BarPtrResult
//go:generate go run github.com/tizz98/results/cmd -pkg foo -t Bar -tup-default Bar{} -result-name BarResult
package foo

type Bar struct {
    Baz int
    Field string
}
```
