package lexer2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"monkey/token"
)

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
