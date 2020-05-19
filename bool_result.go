package results

type BoolResult struct {
    value *bool
    err error
}

func (r *BoolResult) IsOk() bool {
    return r.err == nil
}

func (r *BoolResult) IsErr() bool {
    return r.err != nil
}

func (r *BoolResult) Unwrap() bool {
    if r.IsErr() {
        panic("cannot unwrap BoolResult, it is an error")
    }
    return *r.value
}

func (r *BoolResult) UnwrapOr(v bool) bool {
    if r.IsOk() {
        return r.Unwrap()
    }
    return v
}

func (r *BoolResult) UnwrapOrElse(fn func(err error) bool) bool {
    if r.IsOk() {
        return r.Unwrap()
    }
    return fn(r.err)
}

func (r *BoolResult) Ok(v bool) {
    r.checkAbilityToSet()
    r.value = &v
}

func (r *BoolResult) Err(err error) {
    r.checkAbilityToSet()
    r.err = err
}

func (r *BoolResult) GetErr() error {
    return r.err
}

func (r *BoolResult) Tup() (bool, error) {
    return r.UnwrapOr(false), r.err
}

func (r *BoolResult) checkAbilityToSet() {
    if r.isSet() {
        panic("BoolResult is already set, cannot set again")
    }
}

func (r *BoolResult) isSet() bool {
    return r.value != nil || r.err != nil
}