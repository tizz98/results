package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/iancoleman/strcase"
	"golang.org/x/tools/imports"

	"github.com/tizz98/results"
)

func main() {
	var (
		input   results.Input
		verbose bool
	)

	flag.StringVar(&input.Pkg, "pkg", "", "the name of the output package")
	flag.StringVar(&input.ResultName, "result-name", "", "the result name, like FooResult")
	flag.StringVar(&input.FieldName, "field-name", "value", "the name of the field to be used in the result struct")
	flag.StringVar(&input.T, "t", "", "the type to generate the result for")
	flag.StringVar(&input.TupDefault, "tup-default", "", "the default value to return when calling .Tup()")
	flag.BoolVar(&verbose, "verbose", false, "enable verbose logging")
	flag.BoolVar(&input.GenContext, "gen-ctx", false, "generate context functions")
	flag.StringVar(&input.Name, "name", "", "a human friendly name for the type")
	flag.Parse()

	input.EnsureResultNameNotEmpty()
	input.EnsureNameNotEmpty()

	if verbose {
		fmt.Printf("%#v\n", input)
	}

	if !input.Valid() {
		flag.Usage()
		os.Exit(1)
	}

	fileName := fmt.Sprintf("%s.go", strcase.ToSnake(input.ResultName))
	var buf bytes.Buffer

	results.GenerateResult(&buf, input).Expectf("unable to generate for %s", input.T)
	data := processImports(fileName, buf.Bytes()).Expect("unable to format file using goimports")
	writeFile(fileName, data).Expect("unable to write generated code to file")
}

func processImports(fileName string, data []byte) (result results.ByteSliceResult) {
	result.Set(imports.Process(fileName, data, nil))
	return
}

func writeFile(fileName string, data []byte) (result results.EmptyResult) {
	result.Set(ioutil.WriteFile(fileName, data, 0644))
	return
}
