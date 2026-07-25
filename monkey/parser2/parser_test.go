package parser2

import (
	"fmt"
	"monkey/ast"
	"monkey/lexer"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestExpressionStatement(t *testing.T) {
	var input = `foobar; abc; def; 
1 + 2;
1 - 2;
true != false;
-3 - 2 * 4;
-(3 + 4) * 6;
if (!false) {
	x
} else {
	3 + 4
}
`

	var l = lexer.New(input)
	var p = NewParser(l)
	var program = p.ParseProgram()
	checkParserErrors(t, p)

	checkStatementString(t, program.Statements, []string{
		"foobar",
		"abc",
		"def",
		"(+ 1 2)",
		"(- 1 2)",
		"(!= true false)",
		"(- -3 (* 2 4))",
		"(* -(+ 3 4) 6)",
		"(if !false (block x) else (block (+ 3 4)))",
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
