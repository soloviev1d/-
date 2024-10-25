package lexer

import (
	"bufio"
	"fmt"
	"io"
	"unicode"
)

type Token int
const (
	EOF = iota
	OPERATOR
	KEYWORD
	IDENTEFIER
	SEPARATOR
	LITERAL
	COMMENT
	SPACE
	STRING
	RUNE
	INUM
	RNUM
)

var tokens = []string{
	EOF: "EOF",
	OPERATOR: "OPERATOR",
	KEYWORD: "KEYWORD",
	IDENTEFIER: "IDENTEFIER",
	SEPARATOR: "SEPARATOR",
	LITERAL: "LITERAL",
	COMMENT: "COMMENT",
	SPACE: "SPACE",
	STRING: "STRING",
	RUNE: "RUNE",
	INUM: "INTEGER NUM",
	RNUM: "REAL NUM",
}

func (t Token) String() string{
	return tokens[t]
}

type Position struct{
	line int
	col int
}

type Lexer struct{
	pos *Position
	r *bufio.Reader
}

func NewLexer(reader io.Reader) *Lexer {
	return &Lexer{
		pos: &Position{line: 0, col: 0},
		r: bufio.NewReader(reader),
	}
}

func (l *Lexer) Lex() {
	
}


func (l *Lexer) Unread() {
	if err := l.r.UnreadRune(); err != nil {
		panic(err)
	}

	l.pos.col--
}
func (l *Lexer) Identifier() string {
	var ident string
	for {
		r, _, err := l.r.ReadRune()
		if err != nil{
			if err == io.EOF{
				return ident
			}
		}

		l.pos.col++
		if unicode.IsLetter(r) || unicode.IsDigit(r){
			ident += string(r)
		}else {
			l.Unread()
			return ident
		}
	}
}

func (l* Lexer) Number() string{
	var num string
	for {
		r, _, err := l.r.ReadRune()
		if err != nil{
			if err == io.EOF{
				return num
			}
		}

		l.pos.col++
		if unicode.IsDigit(r){
			num += string(r)
		}else {
			l.Unread()
			return num
		}
	}
}
