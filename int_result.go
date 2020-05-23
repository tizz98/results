// Code generated by github.com/tizz98/results; DO NOT EDIT.
package results

import (
	"context"
	"fmt"
)

type IntResult struct {
	value *int
	err   error
}

// SetNewIntResult is a shortcut to creating a new IntResult and then calling .Set(v, err) on it.
func SetNewIntResult(v int, err error) (result IntResult) {
	result.Set(v, err)
	return
}

// IsOk returns true when the result contains a non-nil result with no error
func (r IntResult) IsOk() bool {
	return r.err == nil
}

// IsErr returns true when the result contains a non-nil error
func (r IntResult) IsErr() bool {
	return r.err != nil
}

// Unwrap panics if the result contains an error, otherwise it returns the value
func (r IntResult) Unwrap() int {
	if r.IsErr() {
		panic("cannot unwrap IntResult, it is an error")
	}
	return *r.value
}

// Expect panics with the specified message if the result contains an error, otherwise it returns the value
func (r IntResult) Expect(message string) int {
	if r.IsErr() {
		panic(fmt.Errorf("%s: %w", message, r.GetErr()))
	}
	return *r.value
}

// Expectf panics with the specified message if the result contains an error, otherwise it returns the value.
// This is different than Expect because if will automatically format the string with the given args.
func (r IntResult) Expectf(format string, args ...interface{}) int {
	if r.IsErr() {
		panic(fmt.Errorf("%s: %w", fmt.Sprintf(format, args...), r.GetErr()))
	}
	return *r.value
}

// UnwrapOr returns the value if there is not an error, otherwise the specified value is returned
func (r IntResult) UnwrapOr(v int) int {
	if r.IsOk() {
		return r.Unwrap()
	}
	return v
}

// UnwrapOrElse returns the value if there is not an error, otherwise the function is called and the result is returned
func (r IntResult) UnwrapOrElse(fn func(err error) int) int {
	if r.IsOk() {
		return r.Unwrap()
	}
	return fn(r.err)
}

// Ok sets the result to a successful result with the provided value.
// This will panic if the result has already been set to successful or an error.
func (r *IntResult) Ok(v int) {
	r.checkAbilityToSet()
	r.value = &v
}

// Err sets the result to an error result with the provided error.
// This will panic if the result has already been set to successful or an error.
func (r *IntResult) Err(err error) {
	r.checkAbilityToSet()
	r.err = err
}

// GetError returns the error of the result. It may be nil, so check with IntResult.IsErr() first.
func (r IntResult) GetErr() error {
	return r.err
}

// Tup returns a tuple of (int, error) with 0 being returned for int if there is an error
func (r IntResult) Tup() (int, error) {
	return r.UnwrapOr(0), r.err
}

// Set is a shortcut to checking the value of an error before setting the result.
// If there is an error, IntResult.Err(err) will be called, otherwise IntResult.Ok(v) will be called.
func (r *IntResult) Set(v int, err error) {
	if err != nil {
		r.Err(err)
		return
	}

	r.Ok(v)
}

func (r IntResult) checkAbilityToSet() {
	if r.isSet() {
		panic("IntResult is already set, cannot set again")
	}
}

func (r IntResult) isSet() bool {
	return r.value != nil || r.err != nil
}

// ContextWithInt embeds the given value of int into the context for later retrieval with IntFromContext
func ContextWithInt(ctx context.Context, key interface{}, v int) context.Context {
	return context.WithValue(ctx, key, v)
}

// IntFromContext attempts to retrieve a int value from the specified context. A IntResult is returned
// which can be used to inspect the success or failure of retrieval.
func IntFromContext(ctx context.Context, key interface{}) (result IntResult) {
	if v, ok := ctx.Value(key).(int); !ok {
		result.Err(fmt.Errorf("%#v not found in context", key))
	} else {
		result.Ok(v)
	}
	return
}
