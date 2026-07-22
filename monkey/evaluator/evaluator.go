package evaluator

import (
	"fmt"
	"monkey/ast"
	"monkey/object"
)

var (
	NULL  = object.NewNull()
	TRUE  = object.NewBoolean(true)
	FALSE = object.NewBoolean(false)
)

func New(program *ast.Program) *Evaluator {
	return &Evaluator{
		program: program,
	}
}

type Evaluator struct {
	program *ast.Program
}

func (e *Evaluator) Eval() object.Object {
	return e.EvalImpl(e.program)
}

func (e *Evaluator) EvalImpl(node ast.Node) object.Object {
	if evalFunc, ok := evalFuncTable[node.Kind()]; ok {
		return evalFunc(e, node)
	}
	panic(fmt.Sprintf("unknown node type: %d", node.Kind()))
}

func (e *Evaluator) evalProgram(node ast.Node) object.Object {
	var program = node.(*ast.Program)

	var result object.Object

	for _, statement := range program.Statements {
		result = e.EvalImpl(statement)
	}

	return result
}

func (e *Evaluator) evalExpressionStatement(node ast.Node) object.Object {
	var expStmt = node.(*ast.ExpressionStatement)

	return e.EvalImpl(expStmt.Expression)
}

func (e *Evaluator) evalIntegerLiteral(node ast.Node) object.Object {
	var intLiteral = node.(*ast.IntegerLiteral)

	return object.NewInteger(intLiteral.Value)
}

type evalFunc func(e *Evaluator, node ast.Node) object.Object

var evalFuncTable map[ast.NodeType]evalFunc

func init() {
	evalFuncTable = map[ast.NodeType]evalFunc{
		ast.NodeProgram:             (*Evaluator).evalProgram,
		ast.NodeExpressionStatement: (*Evaluator).evalExpressionStatement,
		// ast.NodePrefixExpression:    (*Evaluator).evalPrefixExpression,
		// ast.NodeInfixExpression:     (*Evaluator).evalInfixExpression,
		// ast.NodeBoolean:             (*Evaluator).evalBoolean,
		// ast.NodeIfExpression:        (*Evaluator).evalIfExpression,
		// ast.NodeBlockStatement:      (*Evaluator).evalBlockStatement,
		// ast.NodeReturnStatement:     (*Evaluator).evalReturnStatement,
		// ast.NodeLetStatement:        (*Evaluator).evalLetStatement,
		// ast.NodeIdentifier:          (*Evaluator).evalIdentifier,
		// ast.NodeFunctionLiteral:     (*Evaluator).evalFunctionLiteral,
		// ast.NodeCallExpression:      (*Evaluator).evalCallExpression,
		ast.NodeIntegerLiteral: (*Evaluator).evalIntegerLiteral,
	}
}
