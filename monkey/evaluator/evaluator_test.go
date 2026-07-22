package evaluator

import (
	"monkey/lexer"
	"monkey/parser"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEval(t *testing.T) {
	var input = "15"

	l := lexer.New(input)
	p := parser.New(l)
	eval := New(p.ParseProgram())
	obj := eval.Eval()
	assert.Equal(t, "15", obj.Inspect())
}
