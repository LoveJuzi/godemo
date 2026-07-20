package lexer

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"monkey/token"
)

func TestGetTokenCase2(t *testing.T) {
	var input = `
a + add(b * c) + d
`

	var golden = []string{
		"IDENT: a",
		"+: +",
		"IDENT: add",
		"(: (",
		"IDENT: b",
		"*: *",
		"IDENT: c",
		"): )",
		"+: +",
		"IDENT: d",
	}

	l := New(input)

	var tok token.Token

	for idx := 0; ; idx++ {
		tok = l.GetToken()
		if tok.IsEofToken() {
			break
		}

		assert.Equal(t, golden[idx], tok.String(), fmt.Sprintf("index: %d", idx))
	}
}

func TestGetToken(t *testing.T) {
	var input = `+-=(){},;
123
abc
fn
true
false
if
else
return
==
!=
><
`

	var golden = []string{
		"+: +",
		"-: -",
		"=: =",
		"(: (",
		"): )",
		"{: {",
		"}: }",
		",: ,",
		";: ;",
		"INT: 123",
		"IDENT: abc",
		"FUNCTION: fn",
		"TRUE: true",
		"FALSE: false",
		"IF: if",
		"ELSE: else",
		"RETURN: return",
		"==: ==",
		"!=: !=",
		">: >",
		"<: <",
	}

	l := New(input)

	var tok token.Token

	for idx := 0; ; idx++ {
		tok = l.GetToken()
		if tok.IsEofToken() {
			break
		}

		assert.Equal(t, golden[idx], tok.String(), fmt.Sprintf("index: %d", idx))
	}
}
