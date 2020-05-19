package results

type Float32Result struct {
    value *float32
    err error
}

func (r *Float32Result) IsOk() bool {
    return r.err == nil
}

func (r *Float32Result) IsErr() bool {
    return r.err != nil
}

func (r *Float32Result) Unwrap() float32 {
    if r.IsErr() {
        panic("cannot unwrap Float32Result, it is an error")
    }
    return *r.value
}

func (r *Float32Result) UnwrapOr(v float32) float32 {
    if r.IsOk() {
        return r.Unwrap()
    }
    return v
}

func (r *Float32Result) UnwrapOrElse(fn func(err error) float32) float32 {
    if r.IsOk() {
        return r.Unwrap()
    }
    return fn(r.err)
}

func (r *Float32Result) Ok(v float32) {
    r.checkAbilityToSet()
    r.value = &v
}

func (r *Float32Result) Err(err error) {
    r.checkAbilityToSet()
    r.err = err
}

func (r *Float32Result) GetErr() error {
    return r.err
}

func (r *Float32Result) Tup() (float32, error) {
    return r.UnwrapOr(0), r.err
}

func (r *Float32Result) checkAbilityToSet() {
    if r.isSet() {
        panic("Float32Result is already set, cannot set again")
    }
}

func (r *Float32Result) isSet() bool {
    return r.value != nil || r.err != nil
}
