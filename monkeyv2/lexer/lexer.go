package lexer

import (
	"monkeyv2/token"
)

type Lexer struct {
	input  string
	curPos int
	endPos int
}

type tokenParser interface {
	run(l *Lexer) token.Token
}

type eofParser struct{}

func (eofParser) run(l *Lexer) token.Token {
	return token.NewToken(token.EOF, "")
}

type signalCharParser struct {
	tokenType token.TokenType
}

func (s signalCharParser) run(l *Lexer) token.Token {
	ch := l.getChar()
	l.nextPos()
	return token.NewToken(s.tokenType, string(ch))
}

type eqParser struct{}

func (eqParser) run(l *Lexer) token.Token {
	l.nextPos()
	l.nextPos()
	return token.NewToken(token.EQ, "==")
}

type notEqParser struct{}

func (notEqParser) run(l *Lexer) token.Token {
	l.nextPos()
	l.nextPos()
	return token.NewToken(token.NOT_EQ, "!=")
}

var identifierKeywords = map[string]token.TokenType{
	"fn":     token.FUNCTION,
	"let":    token.LET,
	"true":   token.TRUE,
	"false":  token.FALSE,
	"if":     token.IF,
	"else":   token.ELSE,
	"return": token.RETURN,
}

type identifierParser struct{}

func (identifierParser) run(l *Lexer) token.Token {
	word := l.readIdentifier()

	if tokenType, ok := identifierKeywords[word]; ok {
		return token.NewToken(tokenType, word)
	}

	return token.NewToken(token.IDENT, word)
}

type numberParser struct{}

func (numberParser) run(l *Lexer) token.Token {
	word := l.readNumber()
	return token.NewToken(token.INT, word)
}

func NewLexer(input string) *Lexer {
	return &Lexer{
		input:  input,
		endPos: len(input),
	}
}

func (l *Lexer) NextToken() token.Token {
	// 跳过空白符号
	l.skipWhitespace()

	// 解析的类型
	parserObj := l.genParserObj()

	// 针对不同的类型进行不同的分发
	return parserObj.run(l)
}

func (l *Lexer) genParserObj() tokenParser {
	ch := l.peekChar()

	if 0 == ch {
		return eofParser{}
	}

	if ',' == ch {
		return signalCharParser{tokenType: token.COMMA}
	}

	if ';' == ch {
		return signalCharParser{tokenType: token.SEMICOLON}
	}

	if '(' == ch {
		return signalCharParser{tokenType: token.LPAREN}
	}

	if ')' == ch {
		return signalCharParser{tokenType: token.RPAREN}
	}

	if '{' == ch {
		return signalCharParser{tokenType: token.LBRACE}
	}

	if '}' == ch {
		return signalCharParser{tokenType: token.RBRACE}
	}

	if '=' == ch {
		if '=' == l.peekCharN(1) {
			return eqParser{}
		}
		return signalCharParser{tokenType: token.ASSIGN}
	}

	if '+' == ch {
		return signalCharParser{tokenType: token.PLUS}
	}

	if '-' == ch {
		return signalCharParser{tokenType: token.MINUS}
	}

	if '!' == ch {
		if '=' == l.peekCharN(1) {
			return notEqParser{}
		}
		return signalCharParser{tokenType: token.BANG}
	}

	if '*' == ch {
		return signalCharParser{tokenType: token.ASTERISK}
	}

	if '/' == ch {
		return signalCharParser{tokenType: token.SLASH}
	}

	if '<' == ch {
		return signalCharParser{tokenType: token.LT}
	}

	if '>' == ch {
		return signalCharParser{tokenType: token.GT}
	}

	if l.isLetter(ch) {
		return identifierParser{}
	}

	if l.isDigit(ch) {
		return numberParser{}
	}

	return signalCharParser{tokenType: token.ILLEGAL}
}

func (l *Lexer) peekChar() byte {
	return l.peekCharN(0)
}

func (l *Lexer) peekCharN(n int) byte {
	pos := l.curPos + n
	if pos >= l.endPos {
		return 0
	}
	return l.input[pos]
}

func (l *Lexer) getChar() byte {
	if l.curPos >= l.endPos {
		return 0
	}
	return l.input[l.curPos]
}

func (l *Lexer) nextPos() {
	l.curPos += 1
}

func (l *Lexer) isLetter(ch byte) bool {
	if ch >= 'a' && ch <= 'z' {
		return true
	}
	if ch >= 'A' && ch <= 'Z' {
		return true
	}
	return false
}

func (l *Lexer) readIdentifier() string {
	pos := l.curPos
	for l.isLetter(l.getChar()) {
		l.nextPos()
	}
	return l.input[pos:l.curPos]
}

func (l *Lexer) isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

func (l *Lexer) readNumber() string {
	pos := l.curPos
	for l.isDigit(l.getChar()) {
		l.nextPos()
	}
	return l.input[pos:l.curPos]
}

func (l *Lexer) isWhitespace(ch byte) bool {
	return ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r'
}

func (l *Lexer) skipWhitespace() {
	for l.isWhitespace(l.getChar()) {
		l.nextPos()
	}
}
