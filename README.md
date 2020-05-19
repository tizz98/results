# Go Results

Similar `Result` structs like in Rust or other languages.

At a very high level, if Go had generics it would look like this:

```go
type Result<T> struct {
    value *T
    err error
}
```

See [`example_test.go`](example_test.go) for usage.

## Example

```go
package main

import (
    "fmt"
    "strconv"

    "github.com/tizz98/results"
)

func betterParser(v string) (result results.IntResult) {
    val, err := strconv.Atoi(v)
    if err != nil {
        result.Err(err)
        return
    }

    result.Ok(val)
    return
}

func betterParser2(v string) (result results.IntResult) {
    result.Set(strconv.Atoi(v))
    return
}


func main() {
    result := betterParser("123").Unwrap()
    fmt.Printf("Got: %d\n", result)

    result2 := betterParser2("456").Unwrap()
    fmt.Printf("Got: %d\n", result2)

    // This will panic if you uncomment
    // _ = betterParser2("foo").Unwrap()
}
```

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
