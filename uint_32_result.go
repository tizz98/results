// Code generated by github.com/tizz98/results; DO NOT EDIT.
package results

type Uint32Result struct {
	value *uint32
	err   error
}

// IsOk returns true when the result contains a non-nil result with no error
func (r Uint32Result) IsOk() bool {
	return r.err == nil
}

// IsErr returns true when the result contains a non-nil error
func (r Uint32Result) IsErr() bool {
	return r.err != nil
}

// Unwrap panics if the result contains an error, otherwise it returns the value
func (r Uint32Result) Unwrap() uint32 {
	if r.IsErr() {
		panic("cannot unwrap Uint32Result, it is an error")
	}
	return *r.value
}

// UnwrapOr returns the value if there is not an error, otherwise the specified value is returned
func (r Uint32Result) UnwrapOr(v uint32) uint32 {
	if r.IsOk() {
		return r.Unwrap()
	}
	return v
}

// UnwrapOrElse returns the value if there is not an error, otherwise the function is called and the result is returned
func (r Uint32Result) UnwrapOrElse(fn func(err error) uint32) uint32 {
	if r.IsOk() {
		return r.Unwrap()
	}
	return fn(r.err)
}

// Ok sets the result to a successful result with the provided value.
// This will panic if the result has already been set to successful or an error.
func (r *Uint32Result) Ok(v uint32) {
	r.checkAbilityToSet()
	r.value = &v
}

// Err sets the result to an error result with the provided error.
// This will panic if the result has already been set to successful or an error.
func (r *Uint32Result) Err(err error) {
	r.checkAbilityToSet()
	r.err = err
}

// GetError returns the error of the result. It may be nil, so check with Uint32Result.IsErr() first.
func (r Uint32Result) GetErr() error {
	return r.err
}

// Tup returns a tuple of (uint32, error) with 0 being returned for uint32 if there is an error
func (r Uint32Result) Tup() (uint32, error) {
	return r.UnwrapOr(0), r.err
}

// Set is a shortcut to checking the value of an error before setting the result.
// If there is an error, Uint32Result.Err(err) will be called, otherwise Uint32Result.Ok(v) will be called.
func (r *Uint32Result) Set(v uint32, err error) {
	if err != nil {
		r.Err(err)
		return
	}

	r.Ok(v)
}

func (r Uint32Result) checkAbilityToSet() {
	if r.isSet() {
		panic("Uint32Result is already set, cannot set again")
	}
}

func (r Uint32Result) isSet() bool {
	return r.value != nil || r.err != nil
}
