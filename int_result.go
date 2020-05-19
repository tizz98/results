package results

type IntResult struct {
    value *int
    err error
}

func (r *IntResult) IsOk() bool {
    return r.err == nil
}

func (r *IntResult) IsErr() bool {
    return r.err != nil
}

func (r *IntResult) Unwrap() int {
    if r.IsErr() {
        panic("cannot unwrap IntResult, it is an error")
    }
    return *r.value
}

func (r *IntResult) UnwrapOr(v int) int {
    if r.IsOk() {
        return r.Unwrap()
    }
    return v
}

func (r *IntResult) UnwrapOrElse(fn func(err error) int) int {
    if r.IsOk() {
        return r.Unwrap()
    }
    return fn(r.err)
}

func (r *IntResult) Ok(v int) {
    r.checkAbilityToSet()
    r.value = &v
}

func (r *IntResult) Err(err error) {
    r.checkAbilityToSet()
    r.err = err
}

func (r *IntResult) GetErr() error {
    return r.err
}

func (r *IntResult) Tup() (int, error) {
    return r.UnwrapOr(0), r.err
}

func (r *IntResult) checkAbilityToSet() {
    if r.isSet() {
        panic("IntResult is already set, cannot set again")
    }
}

func (r *IntResult) isSet() bool {
    return r.value != nil || r.err != nil
}