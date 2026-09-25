package evaluator

import (
	"monkeyv2/ast"
	"monkeyv2/object"
)

var (
	NULL  = &object.Null{}
	TRUE  = &object.Boolean{Value: true}
	FALSE = &object.Boolean{Value: false}
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
	case *ast.Boolean:
		return &evalBoolean{node: node}
	case *ast.PrefixExpression:
		return &evalPrefixExpression{node: node}
	case *ast.InfixExpression:
		return &evalInfixExpression{node: node}
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

type evalBoolean struct{ node *ast.Boolean }

func (e *evalBoolean) run() object.Object {
	return nativeBoolToBooleanObject(e.node.Value)
}

type evalPrefixExpression struct{ node *ast.PrefixExpression }

func (e *evalPrefixExpression) run() object.Object {
	return e.getEvalObj(Eval(e.node.Right)).run()
}

func (e *evalPrefixExpression) getEvalObj(obj object.Object) eval {
	switch e.node.Operator {
	case "!":
		return &evalBangOperatorExpression{obj: obj}
	case "-":
		return &evalMiusPrefixOperatorExpression{obj: obj}
	default:
		return &evalUnkonwn{node: e.node}
	}
}

type evalInfixExpression struct{ node *ast.InfixExpression }

func (e *evalInfixExpression) run() object.Object {
	return e.getEvalObj(Eval(e.node.Left), Eval(e.node.Right)).run()
}

func (e *evalInfixExpression) getEvalObj(
	leftObj object.Object,
	rightObj object.Object) eval {

	base := evalInfixBase{
		leftObj:  leftObj,
		rightObj: rightObj,
	}

	switch e.node.Operator {
	case "+":
		return &evalPlusInfixExpression{
			evalInfixBase: base,
		}

	case "-":
		return &evalMinusInfixExpression{
			evalInfixBase: base,
		}

	case "*":
		return &evalMultiplyInfixExpression{
			evalInfixBase: base,
		}
	case "/":
		return &evalDivideInfixExpression{
			evalInfixBase: base,
		}

	case "<":
		return &evalLessInfixExpression{
			evalInfixBase: base,
		}

	case ">":
		return &evalGreaterInfixExpression{
			evalInfixBase: base,
		}

	case "==":
		return &evalEqInfixExpression{
			evalInfixBase: base,
		}

	case "!=":
		return &evalNotEqInfixExpression{
			evalInfixBase: base,
		}

	default:
		return &evalUnkonwn{node: e.node}
	}
}

type evalInfixBase struct {
	leftObj  object.Object
	rightObj object.Object
}

func (e *evalInfixBase) integers() (*object.Integer, *object.Integer, bool) {
	leftIntegerObj, leftOk := e.leftObj.(*object.Integer)
	rightIntegerObj, rightOk := e.rightObj.(*object.Integer)

	return leftIntegerObj, rightIntegerObj, leftOk && rightOk
}

type evalPlusInfixExpression struct {
	evalInfixBase
}

func (e *evalPlusInfixExpression) run() object.Object {
	if leftIntegerObj, rightIntegerObj, ok := e.integers(); ok {
		return &object.Integer{Value: leftIntegerObj.Value + rightIntegerObj.Value}
	}

	return NULL
}

type evalMinusInfixExpression struct {
	evalInfixBase
}

func (e *evalMinusInfixExpression) run() object.Object {
	if leftIntegerObj, rightIntegerObj, ok := e.integers(); ok {
		return &object.Integer{Value: leftIntegerObj.Value - rightIntegerObj.Value}
	}

	return NULL
}

type evalMultiplyInfixExpression struct {
	evalInfixBase
}

func (e *evalMultiplyInfixExpression) run() object.Object {
	if leftIntegerObj, rightIntegerObj, ok := e.integers(); ok {
		return &object.Integer{Value: leftIntegerObj.Value * rightIntegerObj.Value}
	}

	return NULL
}

type evalDivideInfixExpression struct {
	evalInfixBase
}

func (e *evalDivideInfixExpression) run() object.Object {
	if leftIntegerObj, rightIntegerObj, ok := e.integers(); ok {
		return &object.Integer{Value: leftIntegerObj.Value / rightIntegerObj.Value}
	}

	return NULL
}

type evalLessInfixExpression struct {
	evalInfixBase
}

func (e *evalLessInfixExpression) run() object.Object {
	if leftIntegerObj, rightIntegerObj, ok := e.integers(); ok {
		return nativeBoolToBooleanObject(leftIntegerObj.Value < rightIntegerObj.Value)
	}

	return NULL
}

type evalGreaterInfixExpression struct {
	evalInfixBase
}

func (e *evalGreaterInfixExpression) run() object.Object {
	if leftIntegerObj, rightIntegerObj, ok := e.integers(); ok {
		return nativeBoolToBooleanObject(leftIntegerObj.Value > rightIntegerObj.Value)
	}

	return NULL
}

type evalEqInfixExpression struct {
	evalInfixBase
}

func (e *evalEqInfixExpression) run() object.Object {
	if leftIntegerObj, rightIntegerObj, ok := e.integers(); ok {
		return nativeBoolToBooleanObject(leftIntegerObj.Value == rightIntegerObj.Value)
	}

	return nativeBoolToBooleanObject(e.leftObj == e.rightObj)
}

type evalNotEqInfixExpression struct {
	evalInfixBase
}

func (e *evalNotEqInfixExpression) run() object.Object {
	if leftIntegerObj, rightIntegerObj, ok := e.integers(); ok {
		return nativeBoolToBooleanObject(leftIntegerObj.Value != rightIntegerObj.Value)
	}

	return nativeBoolToBooleanObject(e.leftObj != e.rightObj)
}

type evalBangOperatorExpression struct{ obj object.Object }

func (e *evalBangOperatorExpression) run() object.Object {
	switch e.obj {
	case TRUE:
		return FALSE
	case FALSE:
		return TRUE
	case NULL:
		return TRUE
	default:
		return FALSE
	}
}

type evalMiusPrefixOperatorExpression struct{ obj object.Object }

func (e *evalMiusPrefixOperatorExpression) run() object.Object {
	switch obj := e.obj.(type) {
	case *object.Integer:
		return &object.Integer{Value: -obj.Value}
	default:
		return NULL
	}
}

func evalStatements(stmts []ast.Statement) object.Object {
	var result object.Object

	for _, statement := range stmts {
		result = Eval(statement)
	}

	return result
}

func nativeBoolToBooleanObject(input bool) *object.Boolean {
	if input {
		return TRUE
	}
	return FALSE
}
