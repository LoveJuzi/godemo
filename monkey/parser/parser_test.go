package parser

import (
	"fmt"
	"monkey/ast"
	"monkey/lexer"
	"monkey/token"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseCallExpression(t *testing.T) {
	var input = `
a + add(b * c) + d
`

	var l = lexer.New(input)
	var p = New(l)
	var program = p.ParseProgram()
	checkParserErrors(t, p)

	checkStatementString(t, program.Statements, []string{
		"((a + add((b * c))) + d)", // 0
	})
}

func TestParseFunctionLiteral(t *testing.T) {
	var input = `
fn(x, y) { x + y; }
fn(x, y) { x + y; x;}
fn() { }
`

	var l = lexer.New(input)
	var p = New(l)
	var program = p.ParseProgram()
	checkParserErrors(t, p)

	checkStatementString(t, program.Statements, []string{
		"fn (x, y) {(x + y)}",   // 0
		"fn (x, y) {(x + y) x}", // 1
		"fn () {}",              // 2
	})
}

func TestParseIfExpression(t *testing.T) {
	var input = `
if ( x > y) { x };
if ( x > y) { x } else { y };
`

	var l = lexer.New(input)
	var p = New(l)
	var program = p.ParseProgram()
	checkParserErrors(t, p)

	checkStatementString(t, program.Statements, []string{
		"if (x > y) {x}",          // 0
		"if (x > y) {x} else {y}", // 1
	})
}

func TestParseGroupedExpression(t *testing.T) {
	var input = `3 * (1 + 5) / 6;
3 5 + 3 5 + 4;`

	var l = lexer.New(input)
	var p = New(l)
	var program = p.ParseProgram()
	checkParserErrors(t, p)

	checkStatementString(t, program.Statements, []string{
		"((3 * (1 + 5)) / 6)", // 0
		"3",                   // 1
		"(5 + 3)",             // 2
		"(5 + 4)",             // 3
	})
}

func TestParseBoolean(t *testing.T) {
	var input = "3 > 5 == -false;"

	var l = lexer.New(input)
	var p = New(l)
	var program = p.ParseProgram()
	checkParserErrors(t, p)

	checkStatementString(t, program.Statements, []string{
		"((3 > 5) == (-false))", // 0
	})
}

func TestParsingInfixExpressions(t *testing.T) {
	var input = `1 + 5 + 6;
1 + 5 *6;
5 < 4 != - 3 > - 4;`

	var l = lexer.New(input)
	var p = New(l)
	var program = p.ParseProgram()
	checkParserErrors(t, p)

	checkStatementString(t, program.Statements, []string{
		"((1 + 5) + 6)",              // 0
		"(1 + (5 * 6))",              // 1
		"((5 < 4) != ((-3) > (-4)))", // 2
	})
}

func TestParsingPrefixExpressions(t *testing.T) {
	var input = "!5;-15;"

	var l = lexer.New(input)
	var p = New(l)
	var program = p.ParseProgram()
	checkParserErrors(t, p)

	checkStatementString(t, program.Statements, []string{
		"(!5)",  // 0
		"(-15)", // 1
	})
}

func TestIdentifierExpression(t *testing.T) {
	var input = "foobar;"

	var l = lexer.New(input)
	var p = New(l)
	var program = p.ParseProgram()
	checkParserErrors(t, p)

	require.Equal(t, 1, len(program.Statements))
	assert.Equal(t,
		ast.NewIdentifier(token.New(token.IDENT, "foobar"), "foobar"),
		program.Statements[0].(*ast.ExpressionStatement).Expression.(*ast.Identifier))
}

func TestIntegerLiteralExpression(t *testing.T) {
	var input = `5;`

	var l = lexer.New(input)
	var p = New(l)
	var program = p.ParseProgram()
	checkParserErrors(t, p)

	require.Equal(t, 1, len(program.Statements))
	assert.Equal(t,
		ast.NewIntegerLiteral(token.New(token.INT, "5"), 5),
		program.Statements[0].(*ast.ExpressionStatement).Expression.(*ast.IntegerLiteral))
}

func TestLetStatements(t *testing.T) {
	var input = `
let x = 5;
let y = 10;
let foobar = 838383;
`

	var l = lexer.New(input)
	var p = New(l)
	var program = p.ParseProgram()
	checkParserErrors(t, p)

	checkStatementString(t, program.Statements, []string{
		"let x = 5;",           // 0
		"let y = 10;",          // 1
		"let foobar = 838383;", // 2
	})
}

func TestReturnStatements(t *testing.T) {
	var input = `
return 5;
return 10;
return 993 322;
`

	var l = lexer.New(input)
	var p = New(l)
	var program = p.ParseProgram()
	checkParserErrors(t, p)

	checkStatementString(t, program.Statements, []string{
		"return 5;",   // 0
		"return 10;",  // 1
		"return 993;", // 2
		"322",         // 3
	})
}

func checkParserErrors(t *testing.T, p *Parser) {
	var printErrors = func() string {
		return strings.Join(p.Errors(), "\n")
	}
	require.Equal(t, 0, len(p.Errors()), printErrors())
}

func checkStatementString(
	t *testing.T,
	target []ast.Statement,
	golden []string) {

	assert.Equal(t, len(golden), len(target))
	for i, v := range target {
		assert.Equal(t,
			golden[i],
			v.String(),
			fmt.Sprintf("index: %d", i))
	}
}
