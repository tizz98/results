package results

import (
	"io"
	"text/template"
)

var (
	t = template.Must(template.ParseFiles("gen.gotpl"))
)

type Input struct {
	Pkg        string
	ResultName string
	FieldName  string
	T          string
	TupDefault string
}

func (i Input) Valid() bool {
	return i.Pkg != "" &&
		i.ResultName != "" &&
		i.FieldName != "" &&
		i.T != "" &&
		i.TupDefault != ""

}

func GenerateResult(w io.Writer, in Input) error {
	return t.Execute(w, in)
}
