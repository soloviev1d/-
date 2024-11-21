package lexer

import "unicode"

/*
	Operator / punctuation starts with:

+    &     +=    &=     &&    ==    !=    (    )
-    |     -=    |=     ||    <     <=    [    ]
*    ^     *=    ^=     <-    >     >=    {    }
/    <<    /=    <<=    ++    =     :=    ,    ;
%    >>    %=    >>=    --    !     ...   .    :
&^          &^=          ~
*/
var OpMap = map[rune]struct{}{
	'+': struct{}{},
	'&': struct{}{},
	'=': struct{}{},
	'!': struct{}{},
	'-': struct{}{},
	'|': struct{}{},
	'<': struct{}{},
	'>': struct{}{},
	'[': struct{}{},
	']': struct{}{},
	'/': struct{}{},
	'.': struct{}{},
	'~': struct{}{},
	':': struct{}{},
}

/* keywords
break        default      func         interface    select
case         defer        go           map          struct
chan         else         goto         package      switch
const        fallthrough  if           range        type
continue     for          import       return       var
*/

var KeywordMap = map[string]struct{}{
	"select":      struct{}{},
	"struct":      struct{}{},
	"switch":      struct{}{},
	"type":        struct{}{},
	"var":         struct{}{},
	"return":      struct{}{},
	"import":      struct{}{},
	"for":         struct{}{},
	"continue":    struct{}{},
	"range":       struct{}{},
	"if":          struct{}{},
	"fallthrough": struct{}{},
	"const":       struct{}{},
	"package":     struct{}{},
	"goto":        struct{}{},
	"else":        struct{}{},
	"chan":        struct{}{},
	"map":         struct{}{},
	"go":          struct{}{},
	"defer":       struct{}{},
	"case":        struct{}{},
	"interface":   struct{}{},
	"func":        struct{}{},
	"default":     struct{}{},
	"break":       struct{}{},
}

var DefaultTypeMap = map[string]struct{}{
	"uint8":      struct{}{},
	"uint16":     struct{}{},
	"uint32":     struct{}{},
	"uint64":     struct{}{},
	"int8":       struct{}{},
	"int16":      struct{}{},
	"int32":      struct{}{},
	"int64":      struct{}{},
	"float32":    struct{}{},
	"float64":    struct{}{},
	"complex64":  struct{}{},
	"complex128": struct{}{},
	"byte":       struct{}{},
	"rune":       struct{}{},
	"string":     struct{}{},
	"bool":       struct{}{},
}

// returns true if r is of operator likeness
func isOperator(r rune) bool {
	if _, ok := OpMap[r]; ok {
		return true
	} else {
		return false
	}
}

// returns true if r is of identifier likeness
func isIdentifier(r rune) bool {
	if unicode.IsLetter(r) || r == '_' {
		return true
	} else {
		return false
	}
}

func isLiteral(r rune) bool {
	if r == '\'' ||
		r == '"' ||
		r == '`' ||
		r == '\n' {
		return true
	} else {
		return false
	}

}
