// Code generated by github.com/tizz98/results; DO NOT EDIT.
package results

import (
	"context"
	"fmt"
)

type BoolResult struct {
	value *bool
	err   error
}

// SetNewBoolResult is a shortcut to creating a new BoolResult and then calling .Set(v, err) on it.
func SetNewBoolResult(v bool, err error) (result BoolResult) {
	result.Set(v, err)
	return
}

// SetNewBoolResultPtr is a shortcut to creating a new BoolResult and then calling .Set(v, err) on it.
// This function differs from SetNewBoolResult by returning a pointer to BoolResult.
func SetNewBoolResultPtr(v bool, err error) *BoolResult {
	result := SetNewBoolResult(v, err)
	return &result
}

// NewOptionalBoolResult is a shortcut to creating a new BoolResult and then calling .SetOptional(v, err) on it.
func NewOptionalBoolResult(v *bool, err error) (result BoolResult) {
	result.SetOptional(v, err)
	return
}

// NewOptionalBoolResultPtr is a shortcut to creating a new BoolResult and then calling .SetOptional(v, err) on it.
// This function differs from NewOptionalBoolResult by returning a pointer to BoolResult.
func NewOptionalBoolResultPtr(v *bool, err error) *BoolResult {
	result := NewOptionalBoolResult(v, err)
	return &result
}

// IsOk returns true when the result contains a non-nil result with no error
func (r BoolResult) IsOk() bool {
	return r.err == nil
}

// IsErr returns true when the result contains a non-nil error
func (r BoolResult) IsErr() bool {
	return r.err != nil
}

// Unwrap panics if the result contains an error, otherwise it returns the value
func (r BoolResult) Unwrap() bool {
	if r.IsErr() {
		panic("cannot unwrap BoolResult, it is an error")
	}
	return *r.value
}

// UnwrapTo will call the .Err() method on the other Result if this BoolResult has an error.
// If other is a pointer to a BoolResult, then .Ok() will be called if this BoolResult name does not have an error.
func (r BoolResult) UnwrapTo(other Result) Result {
	if r.IsErr() {
		other.Err(r.GetErr())
	} else if other, ok := other.(*BoolResult); ok {
		other.Ok(r.Unwrap())
	}

	return other
}

// Expect panics with the specified message if the result contains an error, otherwise it returns the value
func (r BoolResult) Expect(message string) bool {
	if r.IsErr() {
		panic(fmt.Errorf("%s: %w", message, r.GetErr()))
	}
	return *r.value
}

// Expectf panics with the specified message if the result contains an error, otherwise it returns the value.
// This is different than Expect because if will automatically format the string with the given args.
func (r BoolResult) Expectf(format string, args ...interface{}) bool {
	if r.IsErr() {
		panic(fmt.Errorf("%s: %w", fmt.Sprintf(format, args...), r.GetErr()))
	}
	return *r.value
}

// UnwrapOr returns the value if there is not an error, otherwise the specified value is returned
func (r BoolResult) UnwrapOr(v bool) bool {
	if r.IsOk() {
		return r.Unwrap()
	}
	return v
}

// UnwrapOrElse returns the value if there is not an error, otherwise the function is called and the result is returned
func (r BoolResult) UnwrapOrElse(fn func(err error) bool) bool {
	if r.IsOk() {
		return r.Unwrap()
	}
	return fn(r.err)
}

// Ok sets the result to a successful result with the provided value.
// This will panic if the result has already been set to successful or an error.
func (r *BoolResult) Ok(v bool) {
	r.clear()
	r.value = &v
}

// Err sets the result to an error result with the provided error.
// This will panic if the result has already been set to successful or an error.
func (r *BoolResult) Err(err error) {
	r.clear()
	r.err = err
}

// GetError returns the error of the result. It may be nil, so check with BoolResult.IsErr() first.
func (r BoolResult) GetErr() error {
	return r.err
}

// Tup returns a tuple of (bool, error) with false being returned for bool if there is an error
func (r BoolResult) Tup() (bool, error) {
	return r.UnwrapOr(false), r.err
}

// Set is a shortcut to checking the value of an error before setting the result.
// If there is an error, BoolResult.Err(err) will be called, otherwise BoolResult.Ok(v) will be called.
func (r *BoolResult) Set(v bool, err error) {
	if err != nil {
		r.Err(err)
	} else {
		r.Ok(v)
	}
}

// SetOptional is similar to Set but can be called with a nil value for bool.
// The error will be checked first to see if it is not nil, then if v is not nil it will be set with .Ok(v).
func (r *BoolResult) SetOptional(v *bool, err error) {
	if err != nil {
		r.Err(err)
	} else if v != nil {
		r.Ok(*v)
	}
}

// Ptr returns a pointer to the BoolResult
func (r BoolResult) Ptr() *BoolResult { return &r }

func (r BoolResult) clear() {
	r.value = nil
	r.err = nil
}

// ResultToBoolResult takes a Result interface and returns a BoolResult. If "r" contains a BoolResult,
// it is returned, otherwise a new BoolResult is returned with an error set.
func ResultToBoolResult(r Result) (result BoolResult) {
	v, ok := r.(*BoolResult)
	if !ok {
		result.Err(fmt.Errorf("expected *BoolResult got %T", v))
		return
	}

	return *v
}

// ContextWithBool embeds the given value of bool into the context for later retrieval with BoolFromContext
func ContextWithBool(ctx context.Context, key interface{}, v bool) context.Context {
	return context.WithValue(ctx, key, v)
}

// BoolFromContext attempts to retrieve a bool value from the specified context. A BoolResult is returned
// which can be used to inspect the success or failure of retrieval.
func BoolFromContext(ctx context.Context, key interface{}) (result BoolResult) {
	if v, ok := ctx.Value(key).(bool); !ok {
		result.Err(fmt.Errorf("%#v not found in context", key))
	} else {
		result.Ok(v)
	}
	return
}
