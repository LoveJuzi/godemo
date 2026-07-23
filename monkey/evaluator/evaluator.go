package evaluator

import (
	"monkey/ast"
	"monkey/object"
	"monkey/token"
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
	return object.NewError("unknown node type: %d", node.Kind())
}

func (e *Evaluator) nativeBoolToBooleanObject(input bool) object.Object {
	if input {
		return TRUE
	}
	return FALSE
}

func (e *Evaluator) evalProgram(node ast.Node) object.Object {
	var program = node.(*ast.Program)

	var result object.Object

	for _, statement := range program.Statements {
		result = e.EvalImpl(statement)
		if result.Type() == object.RETURN_VALUE_OBJ {
			return result.(*object.ReturnValue).Value
		}
		if result.Type() == object.ERROR_OBJ {
			return result
		}
	}

	return result
}

func (e *Evaluator) evalBlockStatement(node ast.Node) object.Object {
	var blockStmt = node.(*ast.BlockStatement)

	var result object.Object

	for _, statement := range blockStmt.Statements {
		result = e.EvalImpl(statement)
		if result.Type() == object.RETURN_VALUE_OBJ {
			return result
		}
		if result.Type() == object.ERROR_OBJ {
			return result
		}
	}

	return result
}

func (e *Evaluator) evalReturnStatement(node ast.Node) object.Object {
	var returnStmt = node.(*ast.ReturnStatement)

	var value = e.EvalImpl(returnStmt.ReturnValue)
	if value.Type() == object.ERROR_OBJ {
		return value
	}

	return object.NewReturnValue(value)
}

func (e *Evaluator) evalExpressionStatement(node ast.Node) object.Object {
	var expStmt = node.(*ast.ExpressionStatement)

	return e.EvalImpl(expStmt.Expression)
}

func (e *Evaluator) evalPrefixExpression(node ast.Node) object.Object {
	var prefixExp = node.(*ast.PrefixExpression)

	var right = e.EvalImpl(prefixExp.Right)
	if right.Type() == object.ERROR_OBJ {
		return right
	}

	if evalPrefixExprFunc, ok := evalPrefixExprFuncTable[prefixExp.Token.Type]; ok {
		return evalPrefixExprFunc(e, right)
	}
	return object.NewError("unknown operator: %s", prefixExp.Operator)
}

func (e *Evaluator) evalInfixExpression(node ast.Node) object.Object {
	var infixExp = node.(*ast.InfixExpression)

	var left = e.EvalImpl(infixExp.Left)
	if left.Type() == object.ERROR_OBJ {
		return left
	}

	var right = e.EvalImpl(infixExp.Right)
	if right.Type() == object.ERROR_OBJ {
		return right
	}

	if evalInfixExprFunc, ok := evalInfixExprFuncTable[infixExp.Token.Type]; ok {
		return evalInfixExprFunc(e, left, right)
	}
	return object.NewError("unknown operator: %s %s %s", left.Type(), infixExp.Operator, right.Type())
}

func (e *Evaluator) evalIfExpression(node ast.Node) object.Object {
	var ifExp = node.(*ast.IfExpression)

	var condition = e.EvalImpl(ifExp.Condition)
	if condition.Type() == object.ERROR_OBJ {
		return condition
	}

	if e.isTruthy(condition) {
		return e.EvalImpl(ifExp.Consequence)
	} else if ifExp.Alternative != nil {
		return e.EvalImpl(ifExp.Alternative)
	} else {
		return NULL
	}
}

func (e *Evaluator) isTruthy(obj object.Object) bool {
	switch obj {
	case TRUE:
		return true
	case FALSE:
		return false
	case NULL:
		return false
	default:
		return true
	}
}

func (e *Evaluator) evalIntegerLiteral(node ast.Node) object.Object {
	var intLiteral = node.(*ast.IntegerLiteral)

	return object.NewInteger(intLiteral.Value)
}

func (e *Evaluator) evalBoolean(node ast.Node) object.Object {
	var boolLiteral = node.(*ast.Boolean)
	return e.nativeBoolToBooleanObject(boolLiteral.Value)
}

func (e *Evaluator) evalBangOperatorExpression(right object.Object) object.Object {
	switch right {
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

func (e *Evaluator) evalMinusPrefixOperatorExpression(right object.Object) object.Object {
	if right.Type() == object.INTEGER_OBJ {
		var value = right.(*object.Integer).Value
		return object.NewInteger(-value)
	}
	return object.NewError("unknown operator: -%s", right.Type())
}

func (e *Evaluator) evalPlusInfixExpression(left object.Object, right object.Object) object.Object {
	if left.Type() == object.INTEGER_OBJ && right.Type() == object.INTEGER_OBJ {
		var leftValue = left.(*object.Integer).Value
		var rightValue = right.(*object.Integer).Value
		return object.NewInteger(leftValue + rightValue)
	}
	return object.NewError("unknown operator: %s + %s", left.Type(), right.Type())
}

func (e *Evaluator) evalMinusInfixExpression(left object.Object, right object.Object) object.Object {
	if left.Type() == object.INTEGER_OBJ && right.Type() == object.INTEGER_OBJ {
		var leftValue = left.(*object.Integer).Value
		var rightValue = right.(*object.Integer).Value
		return object.NewInteger(leftValue - rightValue)
	}
	return object.NewError("unknown operator: %s - %s", left.Type(), right.Type())
}

func (e *Evaluator) evalDivideInfixExpression(left object.Object, right object.Object) object.Object {
	if left.Type() == object.INTEGER_OBJ && right.Type() == object.INTEGER_OBJ {
		var leftValue = left.(*object.Integer).Value
		var rightValue = right.(*object.Integer).Value
		return object.NewInteger(leftValue / rightValue)
	}
	return object.NewError("unknown operator: %s / %s", left.Type(), right.Type())
}

func (e *Evaluator) evalMultiplyInfixExpression(left object.Object, right object.Object) object.Object {
	if left.Type() == object.INTEGER_OBJ && right.Type() == object.INTEGER_OBJ {
		var leftValue = left.(*object.Integer).Value
		var rightValue = right.(*object.Integer).Value
		return object.NewInteger(leftValue * rightValue)
	}
	return object.NewError("unknown operator: %s * %s", left.Type(), right.Type())
}

func (e *Evaluator) evalEqualInfixExpression(left object.Object, right object.Object) object.Object {
	if left.Type() == object.INTEGER_OBJ && right.Type() == object.INTEGER_OBJ {
		var leftValue = left.(*object.Integer).Value
		var rightValue = right.(*object.Integer).Value
		return e.nativeBoolToBooleanObject(leftValue == rightValue)
	}
	if left.Type() == object.BOOLEAN_OBJ && right.Type() == object.BOOLEAN_OBJ {
		return e.nativeBoolToBooleanObject(left == right)
	}
	return object.NewError("unknown operator: %s == %s", left.Type(), right.Type())
}

func (e *Evaluator) evalNotEqualInfixExpression(left object.Object, right object.Object) object.Object {
	if left.Type() == object.INTEGER_OBJ && right.Type() == object.INTEGER_OBJ {
		var leftValue = left.(*object.Integer).Value
		var rightValue = right.(*object.Integer).Value
		return e.nativeBoolToBooleanObject(leftValue != rightValue)
	}
	if left.Type() == object.BOOLEAN_OBJ && right.Type() == object.BOOLEAN_OBJ {
		return e.nativeBoolToBooleanObject(left != right)
	}
	return object.NewError("unknown operator: %s != %s", left.Type(), right.Type())
}

func (e *Evaluator) evalLessThanInfixExpression(left object.Object, right object.Object) object.Object {
	if left.Type() == object.INTEGER_OBJ && right.Type() == object.INTEGER_OBJ {
		var leftValue = left.(*object.Integer).Value
		var rightValue = right.(*object.Integer).Value
		return e.nativeBoolToBooleanObject(leftValue < rightValue)
	}
	return object.NewError("unknown operator: %s < %s", left.Type(), right.Type())
}

func (e *Evaluator) evalGreaterThanInfixExpression(left object.Object, right object.Object) object.Object {
	if left.Type() == object.INTEGER_OBJ && right.Type() == object.INTEGER_OBJ {
		var leftValue = left.(*object.Integer).Value
		var rightValue = right.(*object.Integer).Value
		return e.nativeBoolToBooleanObject(leftValue > rightValue)
	}
	return object.NewError("unknown operator: %s > %s", left.Type(), right.Type())
}

type evalFunc func(e *Evaluator, node ast.Node) object.Object

var evalFuncTable map[ast.NodeType]evalFunc

type evalPrefixExprFunc func(e *Evaluator, right object.Object) object.Object

var evalPrefixExprFuncTable map[token.TokenType]evalPrefixExprFunc

type evalInfixExprFunc func(e *Evaluator, left object.Object, right object.Object) object.Object

var evalInfixExprFuncTable map[token.TokenType]evalInfixExprFunc

func init() {
	evalFuncTable = map[ast.NodeType]evalFunc{
		ast.NodeProgram:             (*Evaluator).evalProgram,
		ast.NodeBlockStatement:      (*Evaluator).evalBlockStatement,
		ast.NodeReturnStatement:     (*Evaluator).evalReturnStatement,
		ast.NodeExpressionStatement: (*Evaluator).evalExpressionStatement,
		ast.NodePrefixExpression:    (*Evaluator).evalPrefixExpression,
		ast.NodeInfixExpression:     (*Evaluator).evalInfixExpression,
		ast.NodeIfExpression:        (*Evaluator).evalIfExpression,
		// ast.NodeLetStatement:        (*Evaluator).evalLetStatement,
		// ast.NodeIdentifier:          (*Evaluator).evalIdentifier,
		// ast.NodeFunctionLiteral:     (*Evaluator).evalFunctionLiteral,
		// ast.NodeCallExpression:      (*Evaluator).evalCallExpression,
		ast.NodeIntegerLiteral: (*Evaluator).evalIntegerLiteral,
		ast.NodeBoolean:        (*Evaluator).evalBoolean,
	}

	evalPrefixExprFuncTable = map[token.TokenType]evalPrefixExprFunc{
		token.BANG:  (*Evaluator).evalBangOperatorExpression,
		token.MINUS: (*Evaluator).evalMinusPrefixOperatorExpression,
	}

	evalInfixExprFuncTable = map[token.TokenType]evalInfixExprFunc{
		token.PLUS:     (*Evaluator).evalPlusInfixExpression,
		token.MINUS:    (*Evaluator).evalMinusInfixExpression,
		token.SLASH:    (*Evaluator).evalDivideInfixExpression,
		token.ASTERISK: (*Evaluator).evalMultiplyInfixExpression,
		token.EQ:       (*Evaluator).evalEqualInfixExpression,
		token.NOT_EQ:   (*Evaluator).evalNotEqualInfixExpression,
		token.LT:       (*Evaluator).evalLessThanInfixExpression,
		token.GT:       (*Evaluator).evalGreaterThanInfixExpression,
	}
}
