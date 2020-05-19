package results

type Int32Result struct {
    value *int32
    err error
}

func (r *Int32Result) IsOk() bool {
    return r.err == nil
}

func (r *Int32Result) IsErr() bool {
    return r.err != nil
}

func (r *Int32Result) Unwrap() int32 {
    if r.IsErr() {
        panic("cannot unwrap Int32Result, it is an error")
    }
    return *r.value
}

func (r *Int32Result) UnwrapOr(v int32) int32 {
    if r.IsOk() {
        return r.Unwrap()
    }
    return v
}

func (r *Int32Result) UnwrapOrElse(fn func(err error) int32) int32 {
    if r.IsOk() {
        return r.Unwrap()
    }
    return fn(r.err)
}

func (r *Int32Result) Ok(v int32) {
    r.checkAbilityToSet()
    r.value = &v
}

func (r *Int32Result) Err(err error) {
    r.checkAbilityToSet()
    r.err = err
}

func (r *Int32Result) GetErr() error {
    return r.err
}

func (r *Int32Result) Tup() (int32, error) {
    return r.UnwrapOr(0), r.err
}

func (r *Int32Result) checkAbilityToSet() {
    if r.isSet() {
        panic("Int32Result is already set, cannot set again")
    }
}

func (r *Int32Result) isSet() bool {
    return r.value != nil || r.err != nil
}
