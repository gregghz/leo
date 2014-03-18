package main

import (
	"./lexer"
	"fmt"
	"os"
)

func main() {
	tokChan := lexer.Lexer(os.Stdin)

	for tok := range tokChan {
		fmt.Println(tok, tok.Value()
	}
}
