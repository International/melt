package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"gitlab.com/alehander42/melt/compiler"
)

func main() {
	if len(os.Args) < 2 {
		problem("Please: filename")
	}
	source, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		problem(fmt.Sprintf("File: %s", err))
	}

	ast, err := compiler.Parse(string(source))
	if err != nil {
		problem(fmt.Sprintf("Parser: %s", err))
	}

	ctx := compiler.NewContext()
	ctx.LoadBuiltinTypes()

	err = ast.TypeCheck(&ctx)
	if err != nil {
		problem(fmt.Sprintf("%s", err))
	}

	text, err := compiler.Generate(ast)
	if err != nil {
		problem(fmt.Sprintf("%s", err))
	}

	err = ioutil.WriteFile(fmt.Sprintf("%s.go", os.Args[1]), []byte(text), 0644)
	if err != nil {
		problem(fmt.Sprintf("%s", err))
	}
}

func problem(message string) {
	fmt.Printf("ERROR:\n  %s\n", message)
	os.Exit(1)
}