package parser

import (
	"monkeyv2/ast"
	"monkeyv2/lexer"
	"monkeyv2/token"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIdentifier(t *testing.T) {
	l := lexer.NewLexer("abc")
	p := NewParser(l)
	program := p.ParserProgram()

	assert.Equal(t, 1, len(program.Statements))
	es, ok := program.Statements[0].(*ast.ExpressionStatement)
	assert.True(t, ok)
	assert.Equal(t, token.IDENT, es.Token.Type)
	ident, ok := es.Expression.(*ast.Identifier)
	assert.True(t, ok)
	assert.Equal(t, token.IDENT, ident.Token.Type)
	assert.Equal(t, "abc", ident.Value)

	assert.Equal(t, "abc", program.String())
}

func TestIntegerLieral(t *testing.T) {
	l := lexer.NewLexer("50")
	p := NewParser(l)
	program := p.ParserProgram()

	assert.Equal(t, 1, len(program.Statements))
	es, ok := program.Statements[0].(*ast.ExpressionStatement)
	assert.True(t, ok)
	assert.Equal(t, token.INT, es.Token.Type)
	ident, ok := es.Expression.(*ast.IntegerLiteral)
	assert.True(t, ok)
	assert.Equal(t, token.INT, ident.Token.Type)
	assert.Equal(t, int64(50), ident.Value)

	assert.Equal(t, "50", program.String())
}

func TestTrue(t *testing.T) {
	l := lexer.NewLexer("true")
	p := NewParser(l)
	program := p.ParserProgram()

	assert.Equal(t, 1, len(program.Statements))
	es, ok := program.Statements[0].(*ast.ExpressionStatement)
	assert.True(t, ok)
	assert.Equal(t, token.TRUE, es.Token.Type)
	ident, ok := es.Expression.(*ast.Boolean)
	assert.True(t, ok)
	assert.Equal(t, token.TRUE, ident.Token.Type)
	assert.Equal(t, true, ident.Value)

	assert.Equal(t, "true", program.String())
}

func TestFalse(t *testing.T) {
	l := lexer.NewLexer("false")
	p := NewParser(l)
	program := p.ParserProgram()

	assert.Equal(t, 1, len(program.Statements))
	es, ok := program.Statements[0].(*ast.ExpressionStatement)
	assert.True(t, ok)
	assert.Equal(t, token.FALSE, es.Token.Type)
	ident, ok := es.Expression.(*ast.Boolean)
	assert.True(t, ok)
	assert.Equal(t, token.FALSE, ident.Token.Type)
	assert.Equal(t, false, ident.Value)

	assert.Equal(t, "false", program.String())
}

func TestPrefixExpression(t *testing.T) {
	l := lexer.NewLexer("-3")
	p := NewParser(l)
	program := p.ParserProgram()

	assert.Equal(t, 1, len(program.Statements))
	es, ok := program.Statements[0].(*ast.ExpressionStatement)
	assert.True(t, ok)
	assert.Equal(t, token.MINUS, es.Token.Type)
	ident, ok := es.Expression.(*ast.PrefixExpression)
	assert.True(t, ok)
	assert.Equal(t, token.MINUS, ident.Token.Type)
	assert.Equal(t, "-", ident.Operator)

	assert.Equal(t, "(-3)", program.String())
}

func TestInfixExpression(t *testing.T) {
	l := lexer.NewLexer("5+3")
	p := NewParser(l)
	program := p.ParserProgram()

	assert.Equal(t, 1, len(program.Statements))
	es, ok := program.Statements[0].(*ast.ExpressionStatement)
	assert.True(t, ok)
	assert.Equal(t, token.INT, es.Token.Type)
	ident, ok := es.Expression.(*ast.InfixExpression)
	assert.True(t, ok)
	assert.Equal(t, token.PLUS, ident.Token.Type)
	assert.Equal(t, "+", ident.Operator)

	assert.Equal(t, "(5 + 3)", program.String())
}

func TestIfExpression(t *testing.T) {
	l := lexer.NewLexer(`
if (true) { 
	false; 
} else { 
	true 
}`)
	p := NewParser(l)
	program := p.ParserProgram()
	assert.Equal(t, 1, len(program.Statements))
	es, ok := program.Statements[0].(*ast.ExpressionStatement)
	assert.True(t, ok)
	ident, ok := es.Expression.(*ast.IfExpression)
	assert.True(t, ok)
	assert.Equal(t, token.IF, ident.Token.Type)

	assert.Equal(t, "if true false else true", program.String())
}


func TestFunctionLiteral(t *testing.T) {
	l := lexer.NewLexer(`
fn (a, b) {
	a + b
	a - b
}`)

	p := NewParser(l)
	program := p.ParserProgram()
	assert.Equal(t, 1, len(program.Statements))
	es, ok := program.Statements[0].(*ast.ExpressionStatement)
	assert.True(t, ok)
	ident, ok := es.Expression.(*ast.FunctionLiteral)
	assert.True(t, ok)
	assert.Equal(t, token.FUNCTION, ident.Token.Type)

	assert.Equal(t, "fn(a, b) (a + b)(a - b)", program.String())
}

func TestFunctionLiteralCase2(t *testing.T) {
	l := lexer.NewLexer(`
fn () {
`)

	p := NewParser(l)
	program := p.ParserProgram()
	assert.Equal(t, 0, len(program.Statements))

	assert.Equal(t, "", program.String())

	assert.Equal(t, 1, len(p.Errors()))
	assert.EqualError(t, p.Errors()[0], "expected next token to be RBRACE, got EOF instead")
}

func TestCallExpression(t *testing.T) {
	l := lexer.NewLexer(`
fn (a, b) {
	a + b
	a - b
}(1, 2 + 3)`)

	p := NewParser(l)
	program := p.ParserProgram()
	assert.Equal(t, 2, len(program.Statements))
	assert.Equal(t, "fn(a, b) (a + b)(a - b)(2 + 3)", program.String())
}

func TestSemicolon(t *testing.T) {
	l := lexer.NewLexer(`;;`)

	p := NewParser(l)
	program := p.ParserProgram()
	assert.Equal(t, 0, len(program.Statements))
}
