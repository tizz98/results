//go:generate go run github.com/tizz98/results/cmd -pkg results -t float32 -tup-default 0 -gen-ctx
//go:generate go run github.com/tizz98/results/cmd -pkg results -t float64 -tup-default 0 -gen-ctx
//go:generate go run github.com/tizz98/results/cmd -pkg results -t int -tup-default 0 -gen-ctx
//go:generate go run github.com/tizz98/results/cmd -pkg results -t int8 -tup-default 0 -gen-ctx
//go:generate go run github.com/tizz98/results/cmd -pkg results -t int16 -tup-default 0 -gen-ctx
//go:generate go run github.com/tizz98/results/cmd -pkg results -t int32 -tup-default 0 -gen-ctx
//go:generate go run github.com/tizz98/results/cmd -pkg results -t int64 -tup-default 0 -gen-ctx
//go:generate go run github.com/tizz98/results/cmd -pkg results -t uint -tup-default 0 -gen-ctx
//go:generate go run github.com/tizz98/results/cmd -pkg results -t uint8 -tup-default 0 -gen-ctx
//go:generate go run github.com/tizz98/results/cmd -pkg results -t uint16 -tup-default 0 -gen-ctx
//go:generate go run github.com/tizz98/results/cmd -pkg results -t uint32 -tup-default 0 -gen-ctx
//go:generate go run github.com/tizz98/results/cmd -pkg results -t uint64 -tup-default 0 -gen-ctx
//go:generate go run github.com/tizz98/results/cmd -pkg results -t bool -tup-default false -gen-ctx
//go:generate go run github.com/tizz98/results/cmd -pkg results -t string -tup-default emptyString -gen-ctx
//go:generate go run github.com/tizz98/results/cmd -pkg results -t []byte -tup-default nil -result-name ByteSliceResult -name ByteSlice -gen-ctx
//go:generate go run github.com/tizz98/results/cmd -pkg results -t time.Time -tup-default time.Time{} -result-name TimeResult -name Time -gen-ctx
//go:generate go run github.com/tizz98/results/cmd -pkg results -t *time.Time -tup-default nil -result-name TimePtrResult -name TimePtr -gen-ctx
package results

// Result is the common interface all "results" share. Because the return value of .Unwrap() and similar calls is dynamic,
// These are the only functions that can be used in a generic way, if needed.
type Result interface {
	IsOk() bool
	IsErr() bool
	GetErr() error
	Err(err error)
	UnwrapTo(other Result) Result
}
