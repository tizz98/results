package results

import (
	"io"
	"strings"
	"text/template"
)

var (
	tmpl = template.Must(template.New("").Parse(templateText))
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

func (i *Input) ReplacePlaceholders() {
	i.TupDefault = strings.ReplaceAll(i.TupDefault, "emptyString", `""`)
}

func GenerateResult(w io.Writer, in Input) error {
	in.ReplacePlaceholders()
	return tmpl.Execute(w, in)
}
