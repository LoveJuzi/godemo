package token

import "fmt"

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

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// 运算符
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"

	LT     = "<"
	GT     = ">"
	EQ     = "=="
	NOT_EQ = "!="

	// 分隔符
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// 标识符 + 字面量
	IDENT = "IDENT"
	INT   = "INT"

	// 关键字
	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"

	// EXPRESSION
	EXPRESSION = "EXPRESSION"
)

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

func (t Token) String() string {
	return fmt.Sprintf("%v: %v", t.Type, t.Literal)
}

func (t Token) IsEofToken() bool {
	return EOF == t.Type
}
