// Code generated by github.com/tizz98/results; DO NOT EDIT.
package results

import (
	"context"
	"fmt"
)

type Int64Result struct {
	value *int64
	err   error
}

// SetNewInt64Result is a shortcut to creating a new Int64Result and then calling .Set(v, err) on it.
func SetNewInt64Result(v int64, err error) (result Int64Result) {
	result.Set(v, err)
	return
}

// IsOk returns true when the result contains a non-nil result with no error
func (r Int64Result) IsOk() bool {
	return r.err == nil
}

// IsErr returns true when the result contains a non-nil error
func (r Int64Result) IsErr() bool {
	return r.err != nil
}

// Unwrap panics if the result contains an error, otherwise it returns the value
func (r Int64Result) Unwrap() int64 {
	if r.IsErr() {
		panic("cannot unwrap Int64Result, it is an error")
	}
	return *r.value
}

// Expect panics with the specified message if the result contains an error, otherwise it returns the value
func (r Int64Result) Expect(message string) int64 {
	if r.IsErr() {
		panic(fmt.Errorf("%s: %w", message, r.GetErr()))
	}
	return *r.value
}

// Expectf panics with the specified message if the result contains an error, otherwise it returns the value.
// This is different than Expect because if will automatically format the string with the given args.
func (r Int64Result) Expectf(format string, args ...interface{}) int64 {
	if r.IsErr() {
		panic(fmt.Errorf("%s: %w", fmt.Sprintf(format, args...), r.GetErr()))
	}
	return *r.value
}

// UnwrapOr returns the value if there is not an error, otherwise the specified value is returned
func (r Int64Result) UnwrapOr(v int64) int64 {
	if r.IsOk() {
		return r.Unwrap()
	}
	return v
}

// UnwrapOrElse returns the value if there is not an error, otherwise the function is called and the result is returned
func (r Int64Result) UnwrapOrElse(fn func(err error) int64) int64 {
	if r.IsOk() {
		return r.Unwrap()
	}
	return fn(r.err)
}

// Ok sets the result to a successful result with the provided value.
// This will panic if the result has already been set to successful or an error.
func (r *Int64Result) Ok(v int64) {
	r.checkAbilityToSet()
	r.value = &v
}

// Err sets the result to an error result with the provided error.
// This will panic if the result has already been set to successful or an error.
func (r *Int64Result) Err(err error) {
	r.checkAbilityToSet()
	r.err = err
}

// GetError returns the error of the result. It may be nil, so check with Int64Result.IsErr() first.
func (r Int64Result) GetErr() error {
	return r.err
}

// Tup returns a tuple of (int64, error) with 0 being returned for int64 if there is an error
func (r Int64Result) Tup() (int64, error) {
	return r.UnwrapOr(0), r.err
}

// Set is a shortcut to checking the value of an error before setting the result.
// If there is an error, Int64Result.Err(err) will be called, otherwise Int64Result.Ok(v) will be called.
func (r *Int64Result) Set(v int64, err error) {
	if err != nil {
		r.Err(err)
		return
	}

	r.Ok(v)
}

func (r Int64Result) checkAbilityToSet() {
	if r.isSet() {
		panic("Int64Result is already set, cannot set again")
	}
}

func (r Int64Result) isSet() bool {
	return r.value != nil || r.err != nil
}

// ContextWithInt64 embeds the given value of int64 into the context for later retrieval with Int64FromContext
func ContextWithInt64(ctx context.Context, key interface{}, v int64) context.Context {
	return context.WithValue(ctx, key, v)
}

// Int64FromContext attempts to retrieve a int64 value from the specified context. A Int64Result is returned
// which can be used to inspect the success or failure of retrieval.
func Int64FromContext(ctx context.Context, key interface{}) (result Int64Result) {
	if v, ok := ctx.Value(key).(int64); !ok {
		result.Err(fmt.Errorf("%#v not found in context", key))
	} else {
		result.Ok(v)
	}
	return
}
