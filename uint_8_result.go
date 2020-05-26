// Code generated by github.com/tizz98/results; DO NOT EDIT.
package results

import (
	"context"
	"fmt"
)

type Uint8Result struct {
	value *uint8
	err   error
}

// SetNewUint8Result is a shortcut to creating a new Uint8Result and then calling .Set(v, err) on it.
func SetNewUint8Result(v uint8, err error) (result Uint8Result) {
	result.Set(v, err)
	return
}

// SetNewUint8ResultPtr is a shortcut to creating a new Uint8Result and then calling .Set(v, err) on it.
// This function differs from SetNewUint8Result by returning a pointer to Uint8Result.
func SetNewUint8ResultPtr(v uint8, err error) *Uint8Result {
	result := SetNewUint8Result(v, err)
	return &result
}

// NewOptionalUint8Result is a shortcut to creating a new Uint8Result and then calling .SetOptional(v, err) on it.
func NewOptionalUint8Result(v *uint8, err error) (result Uint8Result) {
	result.SetOptional(v, err)
	return
}

// NewOptionalUint8ResultPtr is a shortcut to creating a new Uint8Result and then calling .SetOptional(v, err) on it.
// This function differs from NewOptionalUint8Result by returning a pointer to Uint8Result.
func NewOptionalUint8ResultPtr(v *uint8, err error) *Uint8Result {
	result := NewOptionalUint8Result(v, err)
	return &result
}

// IsOk returns true when the result contains a non-nil result with no error
func (r Uint8Result) IsOk() bool {
	return r.err == nil
}

// IsErr returns true when the result contains a non-nil error
func (r Uint8Result) IsErr() bool {
	return r.err != nil
}

// Unwrap panics if the result contains an error, otherwise it returns the value
func (r Uint8Result) Unwrap() uint8 {
	if r.IsErr() {
		panic("cannot unwrap Uint8Result, it is an error")
	}
	return *r.value
}

// UnwrapTo will call the .Err() method on the other Result if this Uint8Result has an error.
// If other is a pointer to a Uint8Result, then .Ok() will be called if this Uint8Result name does not have an error.
func (r Uint8Result) UnwrapTo(other Result) Result {
	if r.IsErr() {
		other.Err(r.GetErr())
	} else if other, ok := other.(*Uint8Result); ok {
		other.Ok(r.Unwrap())
	}

	return other
}

// Expect panics with the specified message if the result contains an error, otherwise it returns the value
func (r Uint8Result) Expect(message string) uint8 {
	if r.IsErr() {
		panic(fmt.Errorf("%s: %w", message, r.GetErr()))
	}
	return *r.value
}

// Expectf panics with the specified message if the result contains an error, otherwise it returns the value.
// This is different than Expect because if will automatically format the string with the given args.
func (r Uint8Result) Expectf(format string, args ...interface{}) uint8 {
	if r.IsErr() {
		panic(fmt.Errorf("%s: %w", fmt.Sprintf(format, args...), r.GetErr()))
	}
	return *r.value
}

// UnwrapOr returns the value if there is not an error, otherwise the specified value is returned
func (r Uint8Result) UnwrapOr(v uint8) uint8 {
	if r.IsOk() {
		return r.Unwrap()
	}
	return v
}

// UnwrapOrElse returns the value if there is not an error, otherwise the function is called and the result is returned
func (r Uint8Result) UnwrapOrElse(fn func(err error) uint8) uint8 {
	if r.IsOk() {
		return r.Unwrap()
	}
	return fn(r.err)
}

// Ok sets the result to a successful result with the provided value.
// This will panic if the result has already been set to successful or an error.
func (r *Uint8Result) Ok(v uint8) {
	r.clear()
	r.value = &v
}

// Err sets the result to an error result with the provided error.
// This will panic if the result has already been set to successful or an error.
func (r *Uint8Result) Err(err error) {
	r.clear()
	r.err = err
}

// GetError returns the error of the result. It may be nil, so check with Uint8Result.IsErr() first.
func (r Uint8Result) GetErr() error {
	return r.err
}

// Tup returns a tuple of (uint8, error) with 0 being returned for uint8 if there is an error
func (r Uint8Result) Tup() (uint8, error) {
	return r.UnwrapOr(0), r.err
}

// Set is a shortcut to checking the value of an error before setting the result.
// If there is an error, Uint8Result.Err(err) will be called, otherwise Uint8Result.Ok(v) will be called.
func (r *Uint8Result) Set(v uint8, err error) {
	if err != nil {
		r.Err(err)
	} else {
		r.Ok(v)
	}
}

// SetOptional is similar to Set but can be called with a nil value for uint8.
// The error will be checked first to see if it is not nil, then if v is not nil it will be set with .Ok(v).
func (r *Uint8Result) SetOptional(v *uint8, err error) {
	if err != nil {
		r.Err(err)
	} else if v != nil {
		r.Ok(*v)
	}
}

// Ptr returns a pointer to the Uint8Result
func (r Uint8Result) Ptr() *Uint8Result { return &r }

func (r Uint8Result) clear() {
	r.value = nil
	r.err = nil
}

// ResultToUint8Result takes a Result interface and returns a Uint8Result. If "r" contains a Uint8Result,
// it is returned, otherwise a new Uint8Result is returned with an error set.
func ResultToUint8Result(r Result) (result Uint8Result) {
	v, ok := r.(*Uint8Result)
	if !ok {
		result.Err(fmt.Errorf("expected *Uint8Result got %T", v))
		return
	}

	return *v
}

// ContextWithUint8 embeds the given value of uint8 into the context for later retrieval with Uint8FromContext
func ContextWithUint8(ctx context.Context, key interface{}, v uint8) context.Context {
	return context.WithValue(ctx, key, v)
}

// Uint8FromContext attempts to retrieve a uint8 value from the specified context. A Uint8Result is returned
// which can be used to inspect the success or failure of retrieval.
func Uint8FromContext(ctx context.Context, key interface{}) (result Uint8Result) {
	if v, ok := ctx.Value(key).(uint8); !ok {
		result.Err(fmt.Errorf("%#v not found in context", key))
	} else {
		result.Ok(v)
	}
	return
}
