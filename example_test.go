package results

import (
	"fmt"
	"strconv"
)

func ExampleIntResult_Ok() {
	var result IntResult
	result.Ok(123)
	fmt.Printf("%d\n", result.Unwrap())
	// Output: 123
}

func ExampleIntResult_Err() {
	var result IntResult
	result.Err(fmt.Errorf("FAIL"))
	fmt.Printf("%s\n", result.GetErr())
	// Output: FAIL
}

func ExampleIntResult_IsOk_WithErr() {
	var errResult IntResult
	errResult.Err(fmt.Errorf("FAIL"))
	fmt.Printf("%v\n", errResult.IsOk())
	// Output: false
}

func ExampleIntResult_IsOk_WithInt() {
	var result IntResult
	result.Ok(123)
	fmt.Printf("%v\n", result.IsOk())
	// Output: true
}

func Example() {
	var foo = func(input string) (result IntResult) {
		val, err := strconv.Atoi(input)
		if err != nil {
			result.Err(err)
			return
		}

		result.Ok(val)
		return
	}

	result := foo("123").Unwrap()
	fmt.Printf("%d\n", result)
	// Output: 123
}
