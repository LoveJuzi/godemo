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
	var inputs = []string{
		"true",
		"false",
	}
	var golden = []string{
		"true",
		"false",
	}
	for i, input := range inputs {
		l := lexer.New(input)
		p := parser.New(l)
		obj := New(p.ParseProgram()).Eval()
		assert.Equal(t, golden[i], obj.Inspect(), "input: %s", input)
	}
}

func TestEvalBangOperatorCase1(t *testing.T) {
	var inputs = []string{
		"!true",
		"!false",
		"!!true",
		"!!false",
		"!5",
		"!!5",
	}
	var golden = []string{
		"false",
		"true",
		"true",
		"false",
		"false",
		"true",
	}
	for i, input := range inputs {
		l := lexer.New(input)
		p := parser.New(l)
		obj := New(p.ParseProgram()).Eval()
		assert.Equal(t, golden[i], obj.Inspect(), "input: %s", input)
	}
}

func TestEvalMinusPrefixOperatorCase1(t *testing.T) {
	var inputs = []string{
		"-5",
		"-10",
		"--10",
	}
	var golden = []string{
		"-5",
		"-10",
		"10",
	}
	for i, input := range inputs {
		l := lexer.New(input)
		p := parser.New(l)
		obj := New(p.ParseProgram()).Eval()
		assert.Equal(t, golden[i], obj.Inspect(), "input: %s", input)
	}
}
