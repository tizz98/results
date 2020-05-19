// Code generated by github.com/tizz98/results; DO NOT EDIT.
package results

type Int8Result struct {
    value *int8
    err error
}

func (r *Int8Result) IsOk() bool {
    return r.err == nil
}

func (r *Int8Result) IsErr() bool {
    return r.err != nil
}

func (r *Int8Result) Unwrap() int8 {
    if r.IsErr() {
        panic("cannot unwrap Int8Result, it is an error")
    }
    return *r.value
}

func (r *Int8Result) UnwrapOr(v int8) int8 {
    if r.IsOk() {
        return r.Unwrap()
    }
    return v
}

func (r *Int8Result) UnwrapOrElse(fn func(err error) int8) int8 {
    if r.IsOk() {
        return r.Unwrap()
    }
    return fn(r.err)
}

func (r *Int8Result) Ok(v int8) {
    r.checkAbilityToSet()
    r.value = &v
}

func (r *Int8Result) Err(err error) {
    r.checkAbilityToSet()
    r.err = err
}

func (r *Int8Result) GetErr() error {
    return r.err
}

func (r *Int8Result) Tup() (int8, error) {
    return r.UnwrapOr(0), r.err
}

func (r *Int8Result) checkAbilityToSet() {
    if r.isSet() {
        panic("Int8Result is already set, cannot set again")
    }
}

func (r *Int8Result) isSet() bool {
    return r.value != nil || r.err != nil
}
