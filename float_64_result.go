package results

type Float64Result struct {
    value *float64
    err error
}

func (r *Float64Result) IsOk() bool {
    return r.err == nil
}

func (r *Float64Result) IsErr() bool {
    return r.err != nil
}

func (r *Float64Result) Unwrap() float64 {
    if r.IsErr() {
        panic("cannot unwrap Float64Result, it is an error")
    }
    return *r.value
}

func (r *Float64Result) UnwrapOr(v float64) float64 {
    if r.IsOk() {
        return r.Unwrap()
    }
    return v
}

func (r *Float64Result) UnwrapOrElse(fn func(err error) float64) float64 {
    if r.IsOk() {
        return r.Unwrap()
    }
    return fn(r.err)
}

func (r *Float64Result) Ok(v float64) {
    r.checkAbilityToSet()
    r.value = &v
}

func (r *Float64Result) Err(err error) {
    r.checkAbilityToSet()
    r.err = err
}

func (r *Float64Result) GetErr() error {
    return r.err
}

func (r *Float64Result) Tup() (float64, error) {
    return r.UnwrapOr(0), r.err
}

func (r *Float64Result) checkAbilityToSet() {
    if r.isSet() {
        panic("Float64Result is already set, cannot set again")
    }
}

func (r *Float64Result) isSet() bool {
    return r.value != nil || r.err != nil
}
