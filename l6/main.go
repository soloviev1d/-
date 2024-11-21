package main

import (
	"github.com/soloviev1d/fsm-course-l6/parser"
)

func main() {
	// Это пример: замените функции лексера и токены на ваши собственные
	// tokens := []lexer.Token{lexer.IDENTIFIER, lexer.KEYWORD, lexer.NONTERM, lexer.LBRACE, lexer.RBRACE}

	parser.Parse("testdata/test")
}
