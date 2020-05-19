package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/iancoleman/strcase"

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
	flag.Parse()

	if input.ResultName == "" {
		input.ResultName = fmt.Sprintf("%sResult", strcase.ToCamel(input.T))
	}

	if verbose {
		fmt.Printf("%#v\n", input)
	}

	if !input.Valid() {
		flag.Usage()
		os.Exit(1)
	}

	fileName := fmt.Sprintf("%s.go", strcase.ToSnake(input.ResultName))

	_, err := os.Stat(fileName)
	if err != nil {
		if !os.IsNotExist(err) {
			panic(err)
		}
	} else {
		if err := os.Remove(fileName); err != nil {
			panic(err)
		}
	}

	f, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}

	if err := results.GenerateResult(f, input); err != nil {
		panic(err)
	}
}
