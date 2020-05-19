package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"

	"github.com/iancoleman/strcase"
	"github.com/joncalhoun/pipe"

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

	p, err := pipe.New(
		exec.Command("gofmt"),
		exec.Command("goimports"),
	)
	if err != nil {
		log.Fatalf("unable to create pipe: %s", err.Error())
	}

	r, w := io.Pipe()
	p.Stdin = r
	p.Stdout = f

	if err := p.Start(); err != nil {
		log.Fatalf("unable to start pipe: %s", err.Error())
	}

	if err := results.GenerateResult(w, input); err != nil {
		panic(err)
	}

	if err := w.Close(); err != nil {
		log.Fatalf("unable to close pipe writer: %s", err.Error())
	}

	if err := p.Wait(); err != nil {
		log.Fatalf("error waitng for pipe to finish: %s", err.Error())
	}

}
