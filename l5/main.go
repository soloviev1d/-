package main

import (
	"fmt"
	"os"

	"github.com/soloviev1d/simple-lexer/lexer"
)

func main() {
	file, err := os.Open("test")
	if err != nil {
		panic(err)
	}

	l := lexer.NewLexer(file)
	for {
		pos, tok, lit := l.Lex()
		if tok == lexer.EOF {
			break
		}

		fmt.Printf("%d:%d\t%s\t%s\n", pos.Line, pos.Col, tok, lit)
	}
}
