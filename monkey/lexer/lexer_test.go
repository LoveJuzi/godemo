package lexer

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"monkey/token"
)

func TestNextToken(t *testing.T) {
	input := `=+(){},;`

	l := New(input)

	var tok token.Token

	tok = l.NextToken()
	assert.Equal(t, token.New(token.ASSIGN, "="), tok)
	tok = l.NextToken()
	assert.Equal(t, token.New(token.PLUS, "+"), tok)
	tok = l.NextToken()
	assert.Equal(t, token.New(token.LPAREN, "("), tok)
	tok = l.NextToken()
	assert.Equal(t, token.New(token.RPAREN, ")"), tok)
	tok = l.NextToken()
	assert.Equal(t, token.New(token.LBRACE, "{"), tok)
	tok = l.NextToken()
	assert.Equal(t, token.New(token.RBRACE, "}"), tok)
	tok = l.NextToken()
	assert.Equal(t, token.New(token.COMMA, ","), tok)
	tok = l.NextToken()
	assert.Equal(t, token.New(token.SEMICOLON, ";"), tok)
	tok = l.NextToken()
	assert.Equal(t, token.New(token.EOF, ""), tok)
}

func TestNextTokenV2(t *testing.T) {
	input := `let five = 5;
let ten = 10;

let add = fn(x, y) {
	x + y;
};

let result = add(five, ten);
!-/*5;
5 < 10 > 5;

if else return true false

==
!=
`

	l := New(input)

	var tok token.Token

	tok = l.NextToken()
	assert.Equal(t, token.New(token.LET, "let"), tok)
	tok = l.NextToken()
	assert.Equal(t, token.New(token.IDENT, "five"), tok)
	tok = l.NextToken()
	assert.Equal(t, token.New(token.ASSIGN, "="), tok)
	tok = l.NextToken()
	assert.Equal(t, token.New(token.INT, "5"), tok)
	tok = l.NextToken()
	assert.Equal(t, token.New(token.SEMICOLON, ";"), tok)
	tok = l.NextToken()
	assert.Equal(t, token.New(token.LET, "let"), tok)
	tok = l.NextToken()
	assert.Equal(t, token.New(token.IDENT, "ten"), tok)
	tok = l.NextToken()
	assert.Equal(t, token.New(token.ASSIGN, "="), tok)
	tok = l.NextToken()
	assert.Equal(t, token.New(token.INT, "10"), tok)
	tok = l.NextToken()
	assert.Equal(t, token.New(token.SEMICOLON, ";"), tok)
	tok = l.NextToken()
	assert.Equal(t, token.New(token.LET, "let"), tok)
	tok = l.NextToken()
	assert.Equal(t, token.New(token.IDENT, "add"), tok)
	tok = l.NextToken()
	assert.Equal(t, token.New(token.ASSIGN, "="), tok)
	tok = l.NextToken()
	assert.Equal(t, token.New(token.FUNCTION, "fn"), tok)
	tok = l.NextToken()
	assert.Equal(t, token.New(token.LPAREN, "("), tok)
	tok = l.NextToken()
	assert.Equal(t, token.New(token.IDENT, "x"), tok)
	tok = l.NextToken()
	assert.Equal(t, token.New(token.COMMA, ","), tok)
	tok = l.NextToken()
	assert.Equal(t, token.New(token.IDENT, "y"), tok)
	tok = l.NextToken()
	assert.Equal(t, token.New(token.RPAREN, ")"), tok)
	tok = l.NextToken()
	assert.Equal(t, token.New(token.LBRACE, "{"), tok)
	tok = l.NextToken()
	assert.Equal(t, token.New(token.IDENT, "x"), tok)
	tok = l.NextToken()
	assert.Equal(t, token.New(token.PLUS, "+"), tok)
	tok = l.NextToken()
	assert.Equal(t, token.New(token.IDENT, "y"), tok)
	tok = l.NextToken()
	assert.Equal(t, token.New(token.SEMICOLON, ";"), tok)
	tok = l.NextToken()
	assert.Equal(t, token.New(token.RBRACE, "}"), tok)
	tok = l.NextToken()
	assert.Equal(t, token.New(token.SEMICOLON, ";"), tok)
	tok = l.NextToken()
	assert.Equal(t, token.New(token.LET, "let"), tok)
	tok = l.NextToken()
	assert.Equal(t, token.New(token.IDENT, "result"), tok)
	tok = l.NextToken()
	assert.Equal(t, token.New(token.ASSIGN, "="), tok)
	tok = l.NextToken()
	assert.Equal(t, token.New(token.IDENT, "add"), tok)
	tok = l.NextToken()
	assert.Equal(t, token.New(token.LPAREN, "("), tok)
	tok = l.NextToken()
	assert.Equal(t, token.New(token.IDENT, "five"), tok)
	tok = l.NextToken()
	assert.Equal(t, token.New(token.COMMA, ","), tok)
	tok = l.NextToken()
	assert.Equal(t, token.New(token.IDENT, "ten"), tok)
	tok = l.NextToken()
	assert.Equal(t, token.New(token.RPAREN, ")"), tok)
	tok = l.NextToken()
	assert.Equal(t, token.New(token.SEMICOLON, ";"), tok)
	tok = l.NextToken()
	assert.Equal(t, token.New(token.BANG, "!"), tok)
	tok = l.NextToken()
	assert.Equal(t, token.New(token.MINUS, "-"), tok)
	tok = l.NextToken()
	assert.Equal(t, token.New(token.SLASH, "/"), tok)
	tok = l.NextToken()
	assert.Equal(t, token.New(token.ASTERISK, "*"), tok)
	tok = l.NextToken()
	assert.Equal(t, token.New(token.INT, "5"), tok)
	tok = l.NextToken()
	assert.Equal(t, token.New(token.SEMICOLON, ";"), tok)
	tok = l.NextToken()
	assert.Equal(t, token.New(token.INT, "5"), tok)
	tok = l.NextToken()
	assert.Equal(t, token.New(token.LT, "<"), tok)
	tok = l.NextToken()
	assert.Equal(t, token.New(token.INT, "10"), tok)
	tok = l.NextToken()
	assert.Equal(t, token.New(token.GT, ">"), tok)
	tok = l.NextToken()
	assert.Equal(t, token.New(token.INT, "5"), tok)
	tok = l.NextToken()
	assert.Equal(t, token.New(token.SEMICOLON, ";"), tok)
	tok = l.NextToken()
	assert.Equal(t, token.New(token.IF, "if"), tok)
	tok = l.NextToken()
	assert.Equal(t, token.New(token.ELSE, "else"), tok)
	tok = l.NextToken()
	assert.Equal(t, token.New(token.RETURN, "return"), tok)
	tok = l.NextToken()
	assert.Equal(t, token.New(token.TRUE, "true"), tok)
	tok = l.NextToken()
	assert.Equal(t, token.New(token.FALSE, "false"), tok)
	tok = l.NextToken()
	assert.Equal(t, token.New(token.EQ, "=="), tok)
	tok = l.NextToken()
	assert.Equal(t, token.New(token.NOT_EQ, "!="), tok)
	tok = l.NextToken()
	assert.Equal(t, token.New(token.EOF, ""), tok)
}
