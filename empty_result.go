package results

import "fmt"

// EmptyResult only contains an error, no other value.
// It is useful for functions that only return errors.
type EmptyResult struct {
	err error
}

func SetNewEmptyResult(err error) (result EmptyResult) {
	result.Set(err)
	return
}

// IsOk returns true when the result contains a non-nil result with no error
func (r EmptyResult) IsOk() bool {
	return r.err == nil
}

// IsErr returns true when the result contains a non-nil error
func (r EmptyResult) IsErr() bool {
	return r.err != nil
}

// Unwrap panics if the result contains an error
func (r EmptyResult) Unwrap() {
	if r.IsErr() {
		panic("cannot unwrap EmptyResult, it is an error")
	}
}

// Expect panics with the specified message if the result contains an error.
func (r EmptyResult) Expect(message string) {
	if r.IsErr() {
		panic(fmt.Errorf("%s: %w", message, r.GetErr()))
	}
}

// Expectf panics with the specified message if the result contains an error.
// This is different than Expect because if will automatically format the string with the given args.
func (r EmptyResult) Expectf(format string, args ...interface{}) {
	if r.IsErr() {
		panic(fmt.Errorf("%s: %w", fmt.Sprintf(format, args...), r.GetErr()))
	}
}

// Err sets the result to an error result with the provided error.
// This will panic if the result has already been set to successful or an error.
func (r *EmptyResult) Err(err error) {
	r.checkAbilityToSet()
	r.err = err
}

// GetError returns the error of the result. It may be nil, so check with EmptyResult.IsErr() first.
func (r EmptyResult) GetErr() error {
	return r.err
}

// Set is a shortcut to checking the value of an error before setting the result.
// If there is an error, EmptyResult.Err(err) will be called.
func (r *EmptyResult) Set(err error) {
	if err != nil {
		r.Err(err)
	}
}

func (r EmptyResult) checkAbilityToSet() {
	if r.isSet() {
		panic("EmptyResult is already set, cannot set again")
	}
}

func (r EmptyResult) isSet() bool {
	return r.err != nil
}
