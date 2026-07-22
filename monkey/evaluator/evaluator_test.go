package evaluator

import (
	"monkey/lexer"
	"monkey/parser"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEvalIntegerLiteral(t *testing.T) {
	var input = "15"

	l := lexer.New(input)
	p := parser.New(l)
	eval := New(p.ParseProgram())
	obj := eval.Eval()
	assert.Equal(t, "15", obj.Inspect())
}

func TestEvalBoolean(t *testing.T) {
	var input = "true"
	l := lexer.New(input)
	p := parser.New(l)
	eval := New(p.ParseProgram())
	obj := eval.Eval()
	assert.Equal(t, "true", obj.Inspect())
}
