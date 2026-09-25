package evaluator

import (
	"monkeyv2/ast"
	"monkeyv2/object"
)

func Eval(node ast.Node) object.Object {
	return getEvalObj(node).run()
}

func getEvalObj(node ast.Node) eval {
	switch node := node.(type) {
	case *ast.Program:
		return &evalProgram{node: node}
	case *ast.ExpressionStatement:
		return &evalExpressionStatement{node: node}
	case *ast.IntegerLiteral:
		return &evalIntegerLiteral{node: node}
	default:
		return &evalUnkonwn{node: node}
	}
}

type eval interface{ run() object.Object }

type evalUnkonwn struct{ node ast.Node }

func (e *evalUnkonwn) run() object.Object {
	return nil
}

type evalProgram struct{ node *ast.Program }

func (e *evalProgram) run() object.Object {
	return evalStatements(e.node.Statements)
}

type evalExpressionStatement struct{ node *ast.ExpressionStatement }

func (e *evalExpressionStatement) run() object.Object {
	return Eval(e.node.Expression)
}

type evalIntegerLiteral struct{ node *ast.IntegerLiteral }

func (e *evalIntegerLiteral) run() object.Object {
	return &object.Integer{Value: e.node.Value}
}

func evalStatements(stmts []ast.Statement) object.Object {
	var result object.Object

	for _, statement := range stmts {
		result = Eval(statement)
	}

	return result
}
