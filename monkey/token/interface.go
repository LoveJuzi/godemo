package token

func New(tokenType TokenType, literal string) Token {
	return Token{Type: tokenType, Literal: literal}
}

func NewIdent(literal string) Token {
	return Token{Type: lookupIdent(literal), Literal: literal}
}

var ILLEGALToken = Token{Type: ILLEGAL, Literal: ""}

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

func lookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
