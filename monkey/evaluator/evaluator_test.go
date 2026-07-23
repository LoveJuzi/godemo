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

func TestEvalBangOperator(t *testing.T) {
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

func TestEvalMinusPrefixOperator(t *testing.T) {
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

func TestEvalInfixOperator(t *testing.T) {
	var inputs = []string{
		"5 + 10",
		"5 - 10",
		"5 * 10",
		"10 / 2",
		"5 > 10",
		"5 < 10",
		"5 == 5",
		"5 != 10",
		"5 == 10",
		"5 != 5",
		"true == true",
		"false == false",
		"true != false",
		"false != false",
	}
	var golden = []string{
		"15",
		"-5",
		"50",
		"5",
		"false",
		"true",
		"true",
		"true",
		"false",
		"false",
		"true",
		"true",
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

func TestEvalIfExpression(t *testing.T) {
	var inputs = []string{
		"if (true) { 10 }",
		"if (false) { 10 } else { 20 }",
		"if (5 < 10) { 10 } else { 20 }",
		"if (false) { 10 }",
	}
	var golden = []string{
		"10",
		"20",
		"10",
		"null",
	}
	for i, input := range inputs {
		l := lexer.New(input)
		p := parser.New(l)
		obj := New(p.ParseProgram()).Eval()
		assert.Equal(t, golden[i], obj.Inspect(), "input: %s", input)
	}
}
