// Code generated by github.com/tizz98/results; DO NOT EDIT.
package results

type StringResult struct {
    value *string
    err error
}

func (r *StringResult) IsOk() bool {
    return r.err == nil
}

func (r *StringResult) IsErr() bool {
    return r.err != nil
}

func (r *StringResult) Unwrap() string {
    if r.IsErr() {
        panic("cannot unwrap StringResult, it is an error")
    }
    return *r.value
}

func (r *StringResult) UnwrapOr(v string) string {
    if r.IsOk() {
        return r.Unwrap()
    }
    return v
}

func (r *StringResult) UnwrapOrElse(fn func(err error) string) string {
    if r.IsOk() {
        return r.Unwrap()
    }
    return fn(r.err)
}

func (r *StringResult) Ok(v string) {
    r.checkAbilityToSet()
    r.value = &v
}

func (r *StringResult) Err(err error) {
    r.checkAbilityToSet()
    r.err = err
}

func (r *StringResult) GetErr() error {
    return r.err
}

func (r *StringResult) Tup() (string, error) {
    return r.UnwrapOr(""), r.err
}

func (r *StringResult) checkAbilityToSet() {
    if r.isSet() {
        panic("StringResult is already set, cannot set again")
    }
}

func (r *StringResult) isSet() bool {
    return r.value != nil || r.err != nil
}
