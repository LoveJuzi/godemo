package lexer

import "fmt"
import "monkey/token"

func New(input string) *Lexer {
	l := &Lexer{input: input, position: 0}
	return l
}

type Lexer struct {
	input string

	position int
}

type TokenType byte

const (
	EOF byte = iota
	ILLEGAL
	LETTER
	DIGIT
	EQ
	NOT_EQ
)

var signalToken = map[byte]struct{}{
	'!': {},
	'=': {},
	'+': {},
	',': {},
	';': {},
	'(': {},
	')': {},
	'{': {},
	'}': {},
	'*': {},
	'/': {},
	'-': {},
	'<': {},
	'>': {},
}

type tokenFunc func(l *Lexer) token.Token

var tokenMap = map[byte]tokenFunc{
	'-': func(l *Lexer) token.Token {
		l.nextPostion()
		return token.New(token.MINUS, "-")
	},
	'>': func(l *Lexer) token.Token {
		l.nextPostion()
		return token.New(token.GT, ">")
	},
	'<': func(l *Lexer) token.Token {
		l.nextPostion()
		return token.New(token.LT, "<")
	},
	'/': func(l *Lexer) token.Token {
		l.nextPostion()
		return token.New(token.SLASH, "/")
	},
	'*': func(l *Lexer) token.Token {
		l.nextPostion()
		return token.New(token.ASTERISK, "*")
	},
	'!': func(l *Lexer) token.Token {
		l.nextPostion()
		return token.New(token.BANG, "!")
	},
	'=': func(l *Lexer) token.Token {
		l.nextPostion()
		return token.New(token.ASSIGN, "=")
	},
	';': func(l *Lexer) token.Token {
		l.nextPostion()
		return token.New(token.SEMICOLON, ";")
	},
	'(': func(l *Lexer) token.Token {
		l.nextPostion()
		return token.New(token.LPAREN, "(")
	},
	')': func(l *Lexer) token.Token {
		l.nextPostion()
		return token.New(token.RPAREN, ")")
	},
	',': func(l *Lexer) token.Token {
		l.nextPostion()
		return token.New(token.COMMA, ",")
	},
	'+': func(l *Lexer) token.Token {
		l.nextPostion()
		return token.New(token.PLUS, "+")
	},
	'{': func(l *Lexer) token.Token {
		l.nextPostion()
		return token.New(token.LBRACE, "{")
	},
	'}': func(l *Lexer) token.Token {
		l.nextPostion()
		return token.New(token.RBRACE, "}")
	},
	NOT_EQ: func(l *Lexer) token.Token {
		l.nextPostion()
		l.nextPostion()
		return token.New(token.NOT_EQ, "!=")
	},
	EQ: func(l *Lexer) token.Token {
		l.nextPostion()
		l.nextPostion()
		return token.New(token.EQ, "==")
	},
	LETTER: func(l *Lexer) token.Token {
		return token.NewIdent(l.readIndetifier())
	},
	DIGIT: func(l *Lexer) token.Token {
		return token.New(token.INT, l.readNumber())
	},
	ILLEGAL: func(l *Lexer) token.Token {
		return token.New(token.ILLEGAL, "")
	},
	EOF: func(l *Lexer) token.Token {
		return token.New(token.EOF, "")
	},
}

func (l *Lexer) NextToken() token.Token {
	l.skipWhitespace()

	typ := l.getTokenType()

	fn, ok := tokenMap[typ]
	if !ok {
		panic(fmt.Sprintf("unknown token type: %v", typ))
	}

	return fn(l)
}

func (l *Lexer) getTokenType() byte {
	var ch = l.curChar()

	if 0 == ch {
		return EOF
	}

	if '=' == ch && '=' == l.peekChar() {
		return EQ
	}

	if '!' == ch && '=' == l.peekChar() {
		return NOT_EQ
	}

	if _, ok := signalToken[ch]; ok {
		return ch
	}

	if isLetter(ch) {
		return LETTER
	}

	if isDigit(ch) {
		return DIGIT
	}

	return ILLEGAL
}

func isLetter(ch byte) bool {
	return 'a' <= ch && 'z' >= ch || 'A' <= ch && 'Z' >= ch || '_' == ch
}

func isDigit(ch byte) bool {
	return '0' <= ch && '9' >= ch
}

func isWhitespace(ch byte) bool {
	return ' ' == ch || '\t' == ch || '\n' == ch || '\r' == ch
}

func (l *Lexer) nextPostion() {
	if l.position >= len(l.input) {
		return
	}
	l.position += 1
}

func (l *Lexer) curPostion() int {
	return l.position
}

func (l *Lexer) curChar() byte {
	return l.getChar(l.position)
}

func (l *Lexer) peekChar() byte {
	return l.getChar(l.position + 1)
}

func (l *Lexer) getChar(pos int) byte {
	if pos >= len(l.input) {
		return 0
	}
	return l.input[pos]
}

func (l *Lexer) skipWhitespace() {
	for isWhitespace(l.curChar()) {
		l.nextPostion()
	}
}

func (l *Lexer) readIndetifier() string {
	start := l.curPostion()

	for isLetter(l.curChar()) {
		l.nextPostion()
	}

	return string(l.input[start:l.curPostion()])
}

func (l *Lexer) readNumber() string {
	start := l.curPostion()

	for isDigit(l.curChar()) {
		l.nextPostion()
	}

	return string(l.input[start:l.curPostion()])
}
