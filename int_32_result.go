// Code generated by github.com/tizz98/results; DO NOT EDIT.
package results

import (
	"context"
	"fmt"
)

type Int32Result struct {
	value *int32
	err   error
}

// SetNewInt32Result is a shortcut to creating a new Int32Result and then calling .Set(v, err) on it.
func SetNewInt32Result(v int32, err error) (result Int32Result) {
	result.Set(v, err)
	return
}

// IsOk returns true when the result contains a non-nil result with no error
func (r Int32Result) IsOk() bool {
	return r.err == nil
}

// IsErr returns true when the result contains a non-nil error
func (r Int32Result) IsErr() bool {
	return r.err != nil
}

// Unwrap panics if the result contains an error, otherwise it returns the value
func (r Int32Result) Unwrap() int32 {
	if r.IsErr() {
		panic("cannot unwrap Int32Result, it is an error")
	}
	return *r.value
}

// Expect panics with the specified message if the result contains an error, otherwise it returns the value
func (r Int32Result) Expect(message string) int32 {
	if r.IsErr() {
		panic(fmt.Errorf("%s: %w", message, r.GetErr()))
	}
	return *r.value
}

// Expectf panics with the specified message if the result contains an error, otherwise it returns the value.
// This is different than Expect because if will automatically format the string with the given args.
func (r Int32Result) Expectf(format string, args ...interface{}) int32 {
	if r.IsErr() {
		panic(fmt.Errorf("%s: %w", fmt.Sprintf(format, args...), r.GetErr()))
	}
	return *r.value
}

// UnwrapOr returns the value if there is not an error, otherwise the specified value is returned
func (r Int32Result) UnwrapOr(v int32) int32 {
	if r.IsOk() {
		return r.Unwrap()
	}
	return v
}

// UnwrapOrElse returns the value if there is not an error, otherwise the function is called and the result is returned
func (r Int32Result) UnwrapOrElse(fn func(err error) int32) int32 {
	if r.IsOk() {
		return r.Unwrap()
	}
	return fn(r.err)
}

// Ok sets the result to a successful result with the provided value.
// This will panic if the result has already been set to successful or an error.
func (r *Int32Result) Ok(v int32) {
	r.checkAbilityToSet()
	r.value = &v
}

// Err sets the result to an error result with the provided error.
// This will panic if the result has already been set to successful or an error.
func (r *Int32Result) Err(err error) {
	r.checkAbilityToSet()
	r.err = err
}

// GetError returns the error of the result. It may be nil, so check with Int32Result.IsErr() first.
func (r Int32Result) GetErr() error {
	return r.err
}

// Tup returns a tuple of (int32, error) with 0 being returned for int32 if there is an error
func (r Int32Result) Tup() (int32, error) {
	return r.UnwrapOr(0), r.err
}

// Set is a shortcut to checking the value of an error before setting the result.
// If there is an error, Int32Result.Err(err) will be called, otherwise Int32Result.Ok(v) will be called.
func (r *Int32Result) Set(v int32, err error) {
	if err != nil {
		r.Err(err)
		return
	}

	r.Ok(v)
}

func (r Int32Result) checkAbilityToSet() {
	if r.isSet() {
		panic("Int32Result is already set, cannot set again")
	}
}

func (r Int32Result) isSet() bool {
	return r.value != nil || r.err != nil
}

// ContextWithInt32 embeds the given value of int32 into the context for later retrieval with Int32FromContext
func ContextWithInt32(ctx context.Context, key interface{}, v int32) context.Context {
	return context.WithValue(ctx, key, v)
}

// Int32FromContext attempts to retrieve a int32 value from the specified context. A Int32Result is returned
// which can be used to inspect the success or failure of retrieval.
func Int32FromContext(ctx context.Context, key interface{}) (result Int32Result) {
	if v, ok := ctx.Value(key).(int32); !ok {
		result.Err(fmt.Errorf("%#v not found in context", key))
	} else {
		result.Ok(v)
	}
	return
}
