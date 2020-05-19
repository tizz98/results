//go:generate go run github.com/tizz98/results/cmd -pkg results -t float32 -tup-default 0
//go:generate go run github.com/tizz98/results/cmd -pkg results -t float64 -tup-default 0
//go:generate go run github.com/tizz98/results/cmd -pkg results -t int -tup-default 0
//go:generate go run github.com/tizz98/results/cmd -pkg results -t int8 -tup-default 0
//go:generate go run github.com/tizz98/results/cmd -pkg results -t int16 -tup-default 0
//go:generate go run github.com/tizz98/results/cmd -pkg results -t int32 -tup-default 0
//go:generate go run github.com/tizz98/results/cmd -pkg results -t int64 -tup-default 0
//go:generate go run github.com/tizz98/results/cmd -pkg results -t uint -tup-default 0
//go:generate go run github.com/tizz98/results/cmd -pkg results -t uint8 -tup-default 0
//go:generate go run github.com/tizz98/results/cmd -pkg results -t uint16 -tup-default 0
//go:generate go run github.com/tizz98/results/cmd -pkg results -t uint32 -tup-default 0
//go:generate go run github.com/tizz98/results/cmd -pkg results -t uint64 -tup-default 0
//go:generate go run github.com/tizz98/results/cmd -pkg results -t bool -tup-default false
//go:generate go run github.com/tizz98/results/cmd -pkg results -t string -tup-default emptyString
//go:generate go run github.com/tizz98/results/cmd -pkg results -t []byte -tup-default nil -result-name ByteSliceResult
//go:generate go run github.com/tizz98/results/cmd -pkg results -t time.Time -tup-default time.Time{} -result-name TimeResult
//go:generate go run github.com/tizz98/results/cmd -pkg results -t *time.Time -tup-default nil -result-name TimePtrResult
package results
