package main

import (
	"./golang"
	"./lexer"
	"./parser"
	"fmt"
	"os"
)

func main() {
	tokChan := lexer.Lexer(os.Stdin)
	parseTree, errors := parser.Parse(tokChan)

	for _, err := range errors {
		fmt.Println(err)
	}

	if len(errors) == 0 {
		golang.Gen(parseTree, os.Stdout)
	}
}
