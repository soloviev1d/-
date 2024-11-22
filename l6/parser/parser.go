package parser

import (
	"fmt"
	"os"

	"github.com/soloviev1d/fsm-course-l6/lexer"
)

type Parser struct {
	stack []lexer.Token
}

var reduceRuleset map[string]lexer.Token = map[string]lexer.Token{
	// S -> any IDENTIFIER
	stringOfTokens([]lexer.Token{lexer.IDENTIFIER}): lexer.NONTERM,
	// S -> for S
	stringOfTokens([]lexer.Token{lexer.KEYWORD, lexer.NONTERM}): lexer.NONTERM,
	// S -> SS
	stringOfTokens([]lexer.Token{lexer.NONTERM, lexer.NONTERM}): lexer.NONTERM,
	// S -> {}
	stringOfTokens([]lexer.Token{lexer.LBRACE, lexer.RBRACE}): lexer.NONTERM,

	// S -> boolean
	stringOfTokens([]lexer.Token{lexer.BOOLVAL}): lexer.NONTERM,
	// S -> S LOGOP S
	stringOfTokens([]lexer.Token{lexer.NONTERM, lexer.LOGOP, lexer.NONTERM}): lexer.NONTERM,
	// S -> S LOGOP x
	stringOfTokens([]lexer.Token{lexer.NONTERM, lexer.LOGOP, lexer.NUM}): lexer.NONTERM,
	// S -> x LOGOP S
	stringOfTokens([]lexer.Token{lexer.NUM, lexer.LOGOP, lexer.NONTERM}): lexer.NONTERM,

	// S -> x LOGOP x
	stringOfTokens([]lexer.Token{lexer.NUM, lexer.LOGOP, lexer.NUM}): lexer.NONTERM,
}

// func (p *Parser) reduce() bool {
// }
func stringOfTokens(ts []lexer.Token) string {
	s := ""
	for _, t := range ts {
		s += t.String()
	}
	return s
}

func (p *Parser) parseTokens(tokens []lexer.Token) {
	mem := []lexer.Token{}
	fmt.Println(reduceRuleset)
	for len(tokens) > 0 {
		if len(mem) > 0 {
			mem = append(mem[:1], mem[0:]...)
			mem[0] = tokens[len(tokens)-1]
		} else {
			mem = append(mem, tokens[len(tokens)-1])
		}

		fmt.Println("stack: ", mem)
		tokens = tokens[:len(tokens)-1]
		for i := 0; i < len(mem)-1; i++ {
			for j := 0; j < len(mem)-i; j++ {
				// fmt.Println(stringOfTokens(mem[i : j+1]))
				if sub, ok := reduceRuleset[stringOfTokens(mem[i:j+1])]; ok {

					// subbed := mem[:i+1]
					// subbed = append(subbed, sub)
					// subbed = append(subbed, mem[j:]...)
					// mem = subbed
					mem = append(mem[:i], mem[i+j+1:]...)
					if len(mem) == 0 {
						mem = append(mem, sub)
					} else {
						mem = append(mem[:1], mem[0:]...)
						mem[0] = sub
					}
					j = 0

				}
			}
		}
	}
	fmt.Println("stack: ", mem)
}

func Parse(filename string) string {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	l := lexer.NewLexer(file)
	var tokens []lexer.Token
	for {
		_, tok, _ := l.Lex()
		if tok == lexer.EOF {
			break
		}
		if tok == lexer.OPERATOR {
			continue
		}
		tokens = append(tokens, tok)
	}

	parser := Parser{}
	parser.parseTokens(tokens)
	return "foo"
}
