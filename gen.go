package results

import (
	"fmt"
	"io"
	"strings"
	"text/template"
	"unicode"
)

var (
	tmpl = template.Must(template.New("").Parse(templateText))
)

type Input struct {
	Pkg        string
	Name       string
	ResultName string
	FieldName  string
	T          string
	TupDefault string
	GenContext bool
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
	i.Name = fmt.Sprintf("%s%s", string(unicode.ToUpper(rune(i.Name[0]))), i.Name[1:len(i.Name)])
}

func GenerateResult(w io.Writer, in Input) error {
	in.ReplacePlaceholders()
	return tmpl.Execute(w, in)
}
