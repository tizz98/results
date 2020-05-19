package results

type Int64Result struct {
    value *int64
    err error
}

func (r *Int64Result) IsOk() bool {
    return r.err == nil
}

func (r *Int64Result) IsErr() bool {
    return r.err != nil
}

func (r *Int64Result) Unwrap() int64 {
    if r.IsErr() {
        panic("cannot unwrap Int64Result, it is an error")
    }
    return *r.value
}

func (r *Int64Result) UnwrapOr(v int64) int64 {
    if r.IsOk() {
        return r.Unwrap()
    }
    return v
}

func (r *Int64Result) UnwrapOrElse(fn func(err error) int64) int64 {
    if r.IsOk() {
        return r.Unwrap()
    }
    return fn(r.err)
}

func (r *Int64Result) Ok(v int64) {
    r.checkAbilityToSet()
    r.value = &v
}

func (r *Int64Result) Err(err error) {
    r.checkAbilityToSet()
    r.err = err
}

func (r *Int64Result) GetErr() error {
    return r.err
}

func (r *Int64Result) Tup() (int64, error) {
    return r.UnwrapOr(0), r.err
}

func (r *Int64Result) checkAbilityToSet() {
    if r.isSet() {
        panic("Int64Result is already set, cannot set again")
    }
}

func (r *Int64Result) isSet() bool {
    return r.value != nil || r.err != nil
}
