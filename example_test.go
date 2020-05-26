package results

import (
	"context"
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

func ExampleIntResult_Set() {
	var result IntResult
	result.Set(strconv.Atoi("456"))
	fmt.Printf("%d\n", result.Unwrap())
	// Output: 456
}

func ExampleIntFromContext() {
	ctx := context.Background()
	ctx = ContextWithInt(ctx, "foo", 4242)

	result1 := IntFromContext(ctx, "foo")
	result2 := IntFromContext(ctx, "bar")

	fmt.Printf("%v: %d, %v: %d\n", result1.IsOk(), result1.Unwrap(), result2.IsOk(), result2.UnwrapOr(4567))
	// Output: true: 4242, false: 4567
}

func newInt(v int) *int    { return &v }
func newBool(b bool) *bool { return &b }

func ExampleEmptyResult_UnwrapTo() {
	var foo = func() (result EmptyResult) {
		SetNewByteSliceResult([]byte(`123`), nil).
			UnwrapTo(NewOptionalIntResultPtr(newInt(12), nil)).
			UnwrapTo(NewOptionalIntResultPtr(newInt(100), nil)).
			UnwrapTo(NewOptionalIntResultPtr(nil, fmt.Errorf("foo"))).
			UnwrapTo(NewOptionalBoolResultPtr(newBool(false), nil)).
			UnwrapTo(&result)

		return
	}

	result := foo()
	fmt.Printf("%v %s\n", result.IsOk(), result.GetErr())
	// Output: false foo
}

func ExampleBoolResult_UnwrapTo() {
	var foo = func() (result BoolResult) {
		SetNewBoolResult(true, nil).
			UnwrapTo(SetNewBoolResultPtr(false, nil)).
			//UnwrapTo(NewOptionalBoolResultPtr(nil, fmt.Errorf("some failure"))).
			UnwrapTo(SetNewBoolResultPtr(false, nil)).
			UnwrapTo(&result)

		return
	}

	result := foo()
	fmt.Printf("%v %v\n", result.IsOk(), result.Unwrap()) // the unwrapped value is the initial "true"
	// Output: true true
}
