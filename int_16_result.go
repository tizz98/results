// Code generated by github.com/tizz98/results; DO NOT EDIT.
package results

type Int16Result struct {
    value *int16
    err error
}

func (r *Int16Result) IsOk() bool {
    return r.err == nil
}

func (r *Int16Result) IsErr() bool {
    return r.err != nil
}

func (r *Int16Result) Unwrap() int16 {
    if r.IsErr() {
        panic("cannot unwrap Int16Result, it is an error")
    }
    return *r.value
}

func (r *Int16Result) UnwrapOr(v int16) int16 {
    if r.IsOk() {
        return r.Unwrap()
    }
    return v
}

func (r *Int16Result) UnwrapOrElse(fn func(err error) int16) int16 {
    if r.IsOk() {
        return r.Unwrap()
    }
    return fn(r.err)
}

func (r *Int16Result) Ok(v int16) {
    r.checkAbilityToSet()
    r.value = &v
}

func (r *Int16Result) Err(err error) {
    r.checkAbilityToSet()
    r.err = err
}

func (r *Int16Result) GetErr() error {
    return r.err
}

func (r *Int16Result) Tup() (int16, error) {
    return r.UnwrapOr(0), r.err
}

func (r *Int16Result) checkAbilityToSet() {
    if r.isSet() {
        panic("Int16Result is already set, cannot set again")
    }
}

func (r *Int16Result) isSet() bool {
    return r.value != nil || r.err != nil
}
