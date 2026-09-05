package token

type TokenType int

const (
	EOF TokenType = iota

	ILLEGAL

	IDENT
	FUNCTION
	LET
	TRUE
	FALSE
	IF
	ELSE
	RETURN

	INT

	COMMA
	SEMICOLON
	LPAREN
	RPAREN
	LBRACE
	RBRACE

	ASSIGN
	PLUS
	MINUS
	BANG
	ASTERISK
	SLASH

	LT
	GT
	EQ
	NOT_EQ
)

func (t TokenType) String() string {
	switch t {
	case EOF:
		return "EOF"
	case ILLEGAL:
		return "ILLEGAL"
	case IDENT:
		return "IDENT"
	case FUNCTION:
		return "FUNCTION"
	case LET:
		return "LET"
	case TRUE:
		return "TRUE"
	case FALSE:
		return "FALSE"
	case IF:
		return "IF"
	case ELSE:
		return "ELSE"
	case RETURN:
		return "RETURN"
	case INT:
		return "INT"
	case COMMA:
		return "COMMA"
	case SEMICOLON:
		return "SEMICOLON"
	case LPAREN:
		return "LPAREN"
	case RPAREN:
		return "RPAREN"
	case LBRACE:
		return "LBRACE"
	case RBRACE:
		return "RBRACE"
	case ASSIGN:
		return "ASSIGN"
	case PLUS:
		return "PLUS"
	case MINUS:
		return "MINUS"
	case BANG:
		return "BANG"
	case ASTERISK:
		return "ASTERISK"
	case SLASH:
		return "SLASH"
	case LT:
		return "LT"
	case GT:
		return "GT"
	case EQ:
		return "EQ"
	case NOT_EQ:
		return "NOT_EQ"
	default:
		return "UNKNOWN"
	}
}

type Token struct {
	Type    TokenType
	Literal string
}

func NewToken(t TokenType, l string) Token {
	return Token{Type: t, Literal: l}
}

