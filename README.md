# Go Results

Similar `Result` structs like in Rust or other languages.

At a very high level, if Go had generics it would look like this:

```go
type Result<T> struct {
    value *T
    err error
}
```

In [Go 2](https://blog.tempus-ex.com/generics-in-go-how-they-work-and-how-to-play-with-them/), we might someday be able to write:

```go
type Result(type T) struct {
    value *T
    err error
}
```

See [`example_test.go`](example_test.go) for usage.

## Example

```go
package main

import (
    "context"
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

    // Context usage
    ctx := context.Background()
    ctx = results.ContextWithInt(ctx, "foo", 4242)

    result3 := results.IntFromContext(ctx, "foo")
    result4 := results.IntFromContext(ctx, "bar")

    fmt.Printf("%v: %d, %v: %d\n", result3.IsOk(), result3.Unwrap(), result4.IsOk(), result4.UnwrapOr(4567))
    // Output: true: 4242, false: 4567
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
