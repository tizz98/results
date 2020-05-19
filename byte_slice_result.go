// Code generated by github.com/tizz98/results; DO NOT EDIT.
package results

type ByteSliceResult struct {
    value *[]byte
    err error
}

// IsOk returns true when the result contains a non-nil result with no error
func (r ByteSliceResult) IsOk() bool {
    return r.err == nil
}

// IsErr returns true when the result contains a non-nil error
func (r ByteSliceResult) IsErr() bool {
    return r.err != nil
}

// Unwrap panics if the result contains an error, otherwise it returns the value
func (r ByteSliceResult) Unwrap() []byte {
    if r.IsErr() {
        panic("cannot unwrap ByteSliceResult, it is an error")
    }
    return *r.value
}

// UnwrapOr returns the value if there is not an error, otherwise the specified value is returned
func (r ByteSliceResult) UnwrapOr(v []byte) []byte {
    if r.IsOk() {
        return r.Unwrap()
    }
    return v
}

// UnwrapOrElse returns the value if there is not an error, otherwise the function is called and the result is returned
func (r ByteSliceResult) UnwrapOrElse(fn func(err error) []byte) []byte {
    if r.IsOk() {
        return r.Unwrap()
    }
    return fn(r.err)
}

// Ok sets the result to a successful result with the provided value.
// This will panic if the result has already been set to successful or an error.
func (r *ByteSliceResult) Ok(v []byte) {
    r.checkAbilityToSet()
    r.value = &v
}

// Err sets the result to an error result with the provided error.
// This will panic if the result has already been set to successful or an error.
func (r *ByteSliceResult) Err(err error) {
    r.checkAbilityToSet()
    r.err = err
}

// GetError returns the error of the result. It may be nil, so check with ByteSliceResult.IsErr() first.
func (r ByteSliceResult) GetErr() error {
    return r.err
}

// Tup returns a tuple of ([]byte, error) with nil being returned for []byte if there is an error
func (r ByteSliceResult) Tup() ([]byte, error) {
    return r.UnwrapOr(nil), r.err
}

func (r ByteSliceResult) checkAbilityToSet() {
    if r.isSet() {
        panic("ByteSliceResult is already set, cannot set again")
    }
}

func (r ByteSliceResult) isSet() bool {
    return r.value != nil || r.err != nil
}
