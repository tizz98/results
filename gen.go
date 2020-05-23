package results

import (
	"fmt"
	"io"
	"strings"
	"text/template"
	"unicode"

	"github.com/iancoleman/strcase"
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

func (i *Input) EnsureResultNameNotEmpty() {
	if i.ResultName == "" {
		i.ResultName = fmt.Sprintf("%sResult", strcase.ToCamel(i.T))
	}
}

func (i *Input) EnsureNameNotEmpty() {
	if i.Name == "" {
		i.Name = strcase.ToCamel(i.T)
	}
}

func (i *Input) replacePlaceholders() {
	i.TupDefault = strings.ReplaceAll(i.TupDefault, "emptyString", `""`)
	i.Name = fmt.Sprintf("%s%s", string(unicode.ToUpper(rune(i.Name[0]))), i.Name[1:len(i.Name)])
}

func GenerateResult(w io.Writer, in Input) (result EmptyResult) {
	in.replacePlaceholders()
	result.Set(tmpl.Execute(w, in))
	return
}
