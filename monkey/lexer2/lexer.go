package lexer2

import (
	"fmt"
	"monkey/token"
)

type lexerType int

const (
	lexerILLEGAL lexerType = iota
	lexerEOF
	lexerPLUS
	lexerMINUS
	lexerASSIGN
	lexerBANG
	lexerCOMMA
	lexerSEMICOLON
	lexerLPAREN
	lexerRPAREN
	lexerLBRACE
	lexerRBRACE
	lexerEQ
	lexerNOT_EQ
	lexerDIGIT
	lexerLETTER
)

func New(inputStr string) *Lexer {
	return &Lexer{input: &inputStream{input: inputStr, pos: 0}, stack: []token.Token{}}
}

type Lexer struct {
	input *inputStream

	stack []token.Token
}

func (l *Lexer) GetToken() token.Token {
	// 从缓存中获得token
	if tok, ok := l.getTokenFromStack(); ok {
		return tok
	}

	// 跳过空白符号
	l.skipWhiteSpace()

	// 获取分词的类型
	var lexerType_ = l.getLexerType()

	// 获取处理函数
	if fn, ok := getTokenTable[lexerType_]; ok {
		return fn(l)
	}
	panic(fmt.Sprintf("unknown lexer type: %v", lexerType_))
}

func (l *Lexer) UngetToken(tok token.Token) {
	l.stack = append(l.stack, tok)
}

func (l *Lexer) getTokenFromStack() (token.Token, bool) {
	n := len(l.stack)
	if 0 == n {
		return token.ILLEGALToken, false
	}

	defer func() { l.stack = l.stack[:n-1] }()
	return l.stack[n-1], true
}

func (l *Lexer) getLexerType() lexerType {
	var ch byte
	ch = l.input.getChar()
	if 0 == ch {
		l.input.ungetChar()
		return lexerEOF
	}

	if '+' == ch {
		l.input.ungetChar()
		return lexerPLUS
	}

	if '-' == ch {
		l.input.ungetChar()
		return lexerMINUS
	}

	if '=' == ch {
		ch = l.input.getChar()
		if '=' == ch {
			l.input.ungetChar()
			l.input.ungetChar()
			return lexerEQ
		}
		l.input.ungetChar()
		l.input.ungetChar()
		return lexerASSIGN
	}

	if '!' == ch {
		ch = l.input.getChar()
		if '=' == ch {
			l.input.ungetChar()
			l.input.ungetChar()
			return lexerNOT_EQ
		}
		l.input.ungetChar()
		l.input.ungetChar()
		return lexerBANG
	}

	if '(' == ch {
		l.input.ungetChar()
		return lexerLPAREN
	}

	if ')' == ch {
		l.input.ungetChar()
		return lexerRPAREN
	}

	if '{' == ch {
		l.input.ungetChar()
		return lexerLBRACE
	}

	if '}' == ch {
		l.input.ungetChar()
		return lexerRBRACE
	}

	if ',' == ch {
		l.input.ungetChar()
		return lexerCOMMA
	}

	if ';' == ch {
		l.input.ungetChar()
		return lexerSEMICOLON
	}

	if l.isDigit(ch) {
		l.input.ungetChar()
		return lexerDIGIT
	}

	if l.isLetter(ch) {
		l.input.ungetChar()
		return lexerLETTER
	}

	l.input.ungetChar()
	return lexerILLEGAL
}

func (l *Lexer) isWhiteChar(ch byte) bool {
	return ' ' == ch || '\t' == ch || '\n' == ch || '\r' == ch
}

func (l *Lexer) skipWhiteSpace() {
	var ch = l.input.getChar()
	for l.isWhiteChar(ch) {
		ch = l.input.getChar()
	}
	l.input.ungetChar()
}

func (l *Lexer) isDigit(ch byte) bool {
	return '0' <= ch && '9' >= ch
}

func (l *Lexer) readNumber() string {
	var number []byte

	var ch = l.input.getChar()
	for l.isDigit(ch) {
		number = append(number, ch)
		ch = l.input.getChar()
	}
	l.input.ungetChar()

	return string(number)
}

func (l *Lexer) isLetter(ch byte) bool {
	return ('a' <= ch && 'z' >= ch) || ('A' <= ch && 'Z' >= ch) || '_' == ch
}

func (l *Lexer) readWord() string {
	var word []byte

	var ch = l.input.getChar()
	for l.isLetter(ch) {
		word = append(word, ch)
		ch = l.input.getChar()
	}
	l.input.ungetChar()

	return string(word)
}

func (l *Lexer) getIllegalToken() token.Token {
	var ch = l.input.getChar()
	return token.New(token.ILLEGAL, string(ch))
}

func (l *Lexer) getEofToken() token.Token {
	l.input.getChar()
	return token.New(token.EOF, "")
}

func (l *Lexer) getPlusToken() token.Token {
	l.input.getChar()
	return token.New(token.PLUS, "+")
}

func (l *Lexer) getMinusToken() token.Token {
	l.input.getChar()
	return token.New(token.MINUS, "-")
}

func (l *Lexer) getAssignToken() token.Token {
	l.input.getChar()
	return token.New(token.ASSIGN, "=")
}

func (l *Lexer) getBangToken() token.Token {
	l.input.getChar()
	return token.New(token.BANG, "!")
}

func (l *Lexer) getCommaToken() token.Token {
	l.input.getChar()
	return token.New(token.COMMA, ",")
}

func (l *Lexer) getSemicolonToken() token.Token {
	l.input.getChar()
	return token.New(token.SEMICOLON, ";")
}

func (l *Lexer) getLparenToken() token.Token {
	l.input.getChar()
	return token.New(token.LPAREN, "(")
}

func (l *Lexer) getRparenToken() token.Token {
	l.input.getChar()
	return token.New(token.RPAREN, ")")
}

func (l *Lexer) getLbraceToken() token.Token {
	l.input.getChar()
	return token.New(token.LBRACE, "{")
}

func (l *Lexer) getRbraceToken() token.Token {
	l.input.getChar()
	return token.New(token.RBRACE, "}")
}

func (l *Lexer) getEqToken() token.Token {
	l.input.getChar()
	l.input.getChar()
	return token.New(token.EQ, "==")
}

func (l *Lexer) getNotEqToken() token.Token {
	l.input.getChar()
	l.input.getChar()
	return token.New(token.NOT_EQ, "!=")
}

func (l *Lexer) getDigitToken() token.Token {
	return token.New(token.INT, l.readNumber())
}

func (l *Lexer) getWordToken() token.Token {
	return token.NewIdent(l.readWord())
}

type getTokenFunc func(l *Lexer) token.Token

var getTokenTable map[lexerType]getTokenFunc

func init() {
	getTokenTable = map[lexerType]getTokenFunc{
		lexerILLEGAL:   (*Lexer).getIllegalToken,
		lexerEOF:       (*Lexer).getEofToken,
		lexerPLUS:      (*Lexer).getPlusToken,
		lexerMINUS:     (*Lexer).getMinusToken,
		lexerASSIGN:    (*Lexer).getAssignToken,
		lexerBANG:      (*Lexer).getBangToken,
		lexerCOMMA:     (*Lexer).getCommaToken,
		lexerSEMICOLON: (*Lexer).getSemicolonToken,
		lexerLPAREN:    (*Lexer).getLparenToken,
		lexerRPAREN:    (*Lexer).getRparenToken,
		lexerLBRACE:    (*Lexer).getLbraceToken,
		lexerRBRACE:    (*Lexer).getRbraceToken,
		lexerEQ:        (*Lexer).getEqToken,
		lexerNOT_EQ:    (*Lexer).getNotEqToken,
		lexerDIGIT:     (*Lexer).getDigitToken,
		lexerLETTER:    (*Lexer).getWordToken,
	}
}
