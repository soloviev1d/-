package lexer

var opMap = map[rune]struct{}{
	'+': struct{}{},
	'-': struct{}{},
	'/': struct{}{},
	'=': struct{}{},
	'%': struct{}{},
}

func isOperator(r rune) bool {
	x := 0
	// identifier? -> x - identifier -> > - operator 
	//concat opers while symbol is of operator likeness

	// isletter -> isletter -> isletter -> delim -> known keyword for
	//
	// for i, x := range a{}
	x>>=2
	return true
}
