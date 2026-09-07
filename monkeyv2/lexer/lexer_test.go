package lexer

import (
	"monkeyv2/token"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEofToken(t *testing.T) {
	l := NewLexer(" \t\n   ")
	curToken := l.NextToken()

	assert.Equal(t, token.EOF, curToken.Type)
	assert.Equal(t, "", curToken.Literal)
}

func TestAssignToken(t *testing.T) {
	l := NewLexer("=")
	curToken := l.NextToken()

	assert.Equal(t, token.ASSIGN, curToken.Type)
	assert.Equal(t, "=", curToken.Literal)
}

func TestEqToken(t *testing.T) {
	l := NewLexer("==")
	curToken := l.NextToken()

	assert.Equal(t, token.EQ, curToken.Type)
	assert.Equal(t, "==", curToken.Literal)
}

func TestNotEqToken(t *testing.T) {
	l := NewLexer("!=")
	curToken := l.NextToken()

	assert.Equal(t, token.NOT_EQ, curToken.Type)
	assert.Equal(t, "!=", curToken.Literal)
}

func TestIntToken(t *testing.T) {
	l := NewLexer("5")
	curToken := l.NextToken()

	assert.Equal(t, token.INT, curToken.Type)
	assert.Equal(t, "5", curToken.Literal)
}

func TestIllegalToken(t *testing.T) {
	l := NewLexer("@")
	curToken := l.NextToken()

	assert.Equal(t, token.ILLEGAL, curToken.Type)
	assert.Equal(t, "@", curToken.Literal)
}

func TestIfToken(t *testing.T) {
	l := NewLexer("if")
	curToken := l.NextToken()

	assert.Equal(t, token.IF, curToken.Type)
	assert.Equal(t, "if", curToken.Literal)
}

func TestIdentToken(t *testing.T) {
	l := NewLexer("abc")
	curToken := l.NextToken()

	assert.Equal(t, token.IDENT, curToken.Type)
	assert.Equal(t, "abc", curToken.Literal)
}

func TestNextToken(t *testing.T) {
	l := NewLexer(" abc =<>(){}==!;,!=if else fn return")

	cnt := 0
	for {
		curToken := l.NextToken()
		if curToken.Type == token.EOF {
			break
		}
		cnt += 1
	}

	assert.Equal(t, 17, cnt)
}
