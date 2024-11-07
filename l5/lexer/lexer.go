package lexer

import (
	"bufio"
	"io"
	"unicode"
)

type Token int

const (
	EOF = iota
	TYPE
	OPERATOR
	KEYWORD
	IDENTIFIER
	COMMENT
	LITERAL
	NUM
)

var tokens = []string{
	EOF:        "EOF",
	TYPE:       "DEFAULT TYPE",
	OPERATOR:   "OPERATOR/PUNCTUATION",
	KEYWORD:    "KEYWORD",
	IDENTIFIER: "IDENTIFIER",
	COMMENT:    "COMMENT",
	LITERAL:    "LITERAL",
	NUM:        "NUM",
}

func (t Token) String() string {
	return tokens[t]
}

type Position struct {
	Line int
	Col  int
}

type Lexer struct {
	pos *Position
	r   *bufio.Reader
}

func NewLexer(reader io.Reader) *Lexer {
	return &Lexer{
		pos: &Position{Line: 0, Col: 0},
		r:   bufio.NewReader(reader),
	}
}

func (l *Lexer) Lex() (*Position, Token, string) {
	for {
		r, _, err := l.r.ReadRune()
		if err != nil {
			if err == io.EOF {
				return l.pos, EOF, ""
			}
		}

		l.pos.Col++
		// reset if \n
		if r == '\n' {
			l.pos.Col = 0
			l.pos.Line++
			continue
		}

		// ignore comments
		if r == '/' {
			r, _, err := l.r.ReadRune()
			if err != nil {
				if err == io.EOF {
					return l.pos, EOF, ""
				}
			}
			l.pos.Col++
			switch r {
			case '*':
				for {
					r, _, err := l.r.ReadRune()

					if err != nil {
						if err == io.EOF {
							return l.pos, EOF, ""
						}
					}

					l.pos.Col++

					if r != '*' {
						continue
					}
					if r == '\n' {
						l.pos.Col = 0
						l.pos.Line++
					}

					r, _, err = l.r.ReadRune()
					if err != nil {
						if err == io.EOF {
							return l.pos, EOF, ""
						}
					}
					l.pos.Col++
					if r == '/' {
						break
					}
				}
				break
			case '/':
				for {
					r, _, err := l.r.ReadRune()
					if err != nil {
						if err == io.EOF {
							return l.pos, EOF, ""
						}
					}
					l.pos.Col++
					if r == '\n' {
						break

					}
				}

			}
		}
		if unicode.IsSpace(r) {
			continue
		} else if isIdentifier(r) {
			sp := l.pos
			l.Unread()
			lit := l.Identifier()
			if _, ok := KeywordMap[lit]; ok {
				return sp, KEYWORD, lit
			}
			if _, ok := DefaultTypeMap[lit]; ok {
				return sp, TYPE, lit
			}
			return sp, IDENTIFIER, lit
		} else if isOperator(r) {
			sp := l.pos

			l.Unread()
			lit := l.Operator()
			return sp, OPERATOR, lit
		} else if unicode.IsDigit(r) {
			sp := l.pos
			l.Unread()
			lit := l.Number()
			return sp, NUM, lit
		}
		// panic(err)

	}
}

func (l *Lexer) Unread() {
	if err := l.r.UnreadRune(); err != nil {
		panic(err)
	}

	l.pos.Col--
}
func (l *Lexer) Identifier() string {
	var ident string
	for {
		r, _, err := l.r.ReadRune()
		if err != nil {
			if err == io.EOF {
				return ident
			}
		}

		l.pos.Col++
		if unicode.IsLetter(r) || r == '_' {
			ident += string(r)
		} else {
			l.Unread()
			return ident
		}
	}
}

func (l *Lexer) Number() string {
	var num string
	for {
		r, _, err := l.r.ReadRune()
		if err != nil {
			if err == io.EOF {
				return num
			}
		}

		l.pos.Col++
		if unicode.IsDigit(r) || r == '.' {
			num += string(r)
		} else {
			l.Unread()
			return num
		}
	}
}

func (l *Lexer) Operator() string {
	var op string
	for {
		r, _, err := l.r.ReadRune()
		if err != nil {
			if err == io.EOF {
				return op
			}
		}

		l.pos.Col++
		// if !unicode.IsLetter(r) &&
		// 	!unicode.IsSpace(r) &&
		// 	!unicode.IsDigit(r) &&
		// 	r != '\n' &&
		// 	!unicode.IsSpace(r) {
		if r == '\n' {
			l.Unread()
			return op
		}
		if _, ok := OpMap[r]; ok {
			op += string(r)
		} else {
			l.Unread()
			return op
		}
	}

}
