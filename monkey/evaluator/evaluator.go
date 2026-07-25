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
	env     *object.Enviroment
}

func (e *Evaluator) Eval() object.Object {
	e.env = object.NewEnviroment()
	return e.EvalImpl(e.program, e.env)
}

func (e *Evaluator) EvalImpl(node ast.Node, env *object.Enviroment) object.Object {
	if evalFn, ok := evalFns[node.Kind()]; ok {
		return evalFn(e, node, env)
	}
	return object.NewError("unknown node type: %d", node.Kind())
}

func (e *Evaluator) nativeBoolToBooleanObject(input bool) object.Object {
	if input {
		return TRUE
	}
	return FALSE
}

func (e *Evaluator) evalProgram(node ast.Node, env *object.Enviroment) object.Object {
	var program = node.(*ast.Program)

	var result object.Object

	for _, statement := range program.Statements {
		result = e.EvalImpl(statement, env)
		if result.Type() == object.RETURN_VALUE_OBJ {
			return result.(*object.ReturnValue).Value
		}
		if result.Type() == object.ERROR_OBJ {
			return result
		}
	}

	return result
}

func (e *Evaluator) evalBlockStatement(node ast.Node, env *object.Enviroment) object.Object {
	var blockStmt = node.(*ast.BlockStatement)

	var result object.Object

	for _, statement := range blockStmt.Statements {
		result = e.EvalImpl(statement, env)
		if result.Type() == object.RETURN_VALUE_OBJ {
			return result
		}
		if result.Type() == object.ERROR_OBJ {
			return result
		}
	}

	return result
}

func (e *Evaluator) evalReturnStatement(node ast.Node, env *object.Enviroment) object.Object {
	var returnStmt = node.(*ast.ReturnStatement)

	var value = e.EvalImpl(returnStmt.ReturnValue, env)
	if value.Type() == object.ERROR_OBJ {
		return value
	}

	return object.NewReturnValue(value)
}

func (e *Evaluator) evalExpressionStatement(node ast.Node, env *object.Enviroment) object.Object {
	var expStmt = node.(*ast.ExpressionStatement)

	return e.EvalImpl(expStmt.Expression, env)
}

func (e *Evaluator) evalPrefixExpression(node ast.Node, env *object.Enviroment) object.Object {
	var prefixExp = node.(*ast.PrefixExpression)

	var right = e.EvalImpl(prefixExp.Right, env)
	if right.Type() == object.ERROR_OBJ {
		return right
	}

	if fn, ok := evalPrefixExprFns[prefixExp.Token.Type]; ok {
		return fn(e, right)
	}
	return object.NewError("unknown operator: %s", prefixExp.Operator)
}

func (e *Evaluator) evalInfixExpression(node ast.Node, env *object.Enviroment) object.Object {
	var infixExp = node.(*ast.InfixExpression)

	var left = e.EvalImpl(infixExp.Left, env)
	if left.Type() == object.ERROR_OBJ {
		return left
	}

	var right = e.EvalImpl(infixExp.Right, env)
	if right.Type() == object.ERROR_OBJ {
		return right
	}

	if fn, ok := evalInfixExprFns[infixExp.Token.Type]; ok {
		return fn(e, left, right)
	}
	return object.NewError("unknown operator: %s %s %s", left.Type(), infixExp.Operator, right.Type())
}

func (e *Evaluator) evalIfExpression(node ast.Node, env *object.Enviroment) object.Object {
	var ifExp = node.(*ast.IfExpression)

	var condition = e.EvalImpl(ifExp.Condition, env)
	if condition.Type() == object.ERROR_OBJ {
		return condition
	}

	if e.isTruthy(condition) {
		return e.EvalImpl(ifExp.Consequence, env)
	} else if ifExp.Alternative != nil {
		return e.EvalImpl(ifExp.Alternative, env)
	} else {
		return NULL
	}
}

func (e *Evaluator) evalLetStatement(node ast.Node, env *object.Enviroment) object.Object {
	var letStmt = node.(*ast.LetStatement)

	var val = e.EvalImpl(letStmt.Value, env)
	if val.Type() == object.ERROR_OBJ {
		return val
	}

	env.Set(letStmt.Name.Value, val)
	return val
}

func (e *Evaluator) evalIdentifier(node ast.Node, env *object.Enviroment) object.Object {
	var ident = node.(*ast.Identifier)

	if val, ok := env.Get(ident.Value); ok {
		return val
	}

	return object.NewError("identifier not found: %s", ident.Value)
}

func (e *Evaluator) evalFunctionLiteral(node ast.Node, env *object.Enviroment) object.Object {
	var funcLiteral = node.(*ast.FunctionLiteral)

	var params = funcLiteral.Parameters
	var body = funcLiteral.Body

	return object.NewFunction(params, body, env)
}

func (e *Evaluator) evalCallExpression(node ast.Node, env *object.Enviroment) object.Object {
	var callExp = node.(*ast.CallExpression)

	var function = e.EvalImpl(callExp.Function, env)
	if function.Type() == object.ERROR_OBJ {
		return function
	}

	var args = []object.Object{}
	for _, arg := range callExp.Arguments {
		var evaluatedArg = e.EvalImpl(arg, env)
		if evaluatedArg.Type() == object.ERROR_OBJ {
			return evaluatedArg
		}
		args = append(args, evaluatedArg)
	}

	return e.applyFunction(function, args)
}

func (e *Evaluator) applyFunction(fn object.Object, args []object.Object) object.Object {
	if fn, ok := fn.(*object.Function); ok {
		var extendedEnv = e.extendFunctionEnv(fn, args)
		var evaluated = e.EvalImpl(fn.Body, extendedEnv)
		return e.unwrapReturnValue(evaluated)
	}
	return object.NewError("not a function: %s", fn.Type())
}

func (e *Evaluator) extendFunctionEnv(fn *object.Function, args []object.Object) *object.Enviroment {
	var env = object.NewEnclosedEnviroment(fn.Env)

	for paramIdx, param := range fn.Parameters {
		env.Set(param.Value, args[paramIdx])
	}

	return env
}

func (e *Evaluator) unwrapReturnValue(obj object.Object) object.Object {
	if returnValue, ok := obj.(*object.ReturnValue); ok {
		return returnValue.Value
	}
	return obj
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

func (e *Evaluator) evalIntegerLiteral(node ast.Node, env *object.Enviroment) object.Object {
	var intLiteral = node.(*ast.IntegerLiteral)

	return object.NewInteger(intLiteral.Value)
}

func (e *Evaluator) evalBoolean(node ast.Node, env *object.Enviroment) object.Object {
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

type (
	evalFn           func(e *Evaluator, node ast.Node, env *object.Enviroment) object.Object
	evalPrefixExprFn func(e *Evaluator, right object.Object) object.Object
	evalInfixExprFn  func(e *Evaluator, left object.Object, right object.Object) object.Object
)

var (
	evalFns           map[ast.NodeType]evalFn
	evalPrefixExprFns map[token.TokenType]evalPrefixExprFn
	evalInfixExprFns  map[token.TokenType]evalInfixExprFn
)

func init() {
	evalFns = map[ast.NodeType]evalFn{
		ast.NodeProgram:             (*Evaluator).evalProgram,
		ast.NodeBlockStatement:      (*Evaluator).evalBlockStatement,
		ast.NodeReturnStatement:     (*Evaluator).evalReturnStatement,
		ast.NodeExpressionStatement: (*Evaluator).evalExpressionStatement,
		ast.NodePrefixExpression:    (*Evaluator).evalPrefixExpression,
		ast.NodeInfixExpression:     (*Evaluator).evalInfixExpression,
		ast.NodeIfExpression:        (*Evaluator).evalIfExpression,
		ast.NodeLetStatement:        (*Evaluator).evalLetStatement,
		ast.NodeIdentifier:          (*Evaluator).evalIdentifier,
		ast.NodeFunctionLiteral:     (*Evaluator).evalFunctionLiteral,
		ast.NodeCallExpression:      (*Evaluator).evalCallExpression,
		ast.NodeIntegerLiteral:      (*Evaluator).evalIntegerLiteral,
		ast.NodeBoolean:             (*Evaluator).evalBoolean,
	}

	evalPrefixExprFns = map[token.TokenType]evalPrefixExprFn{
		token.BANG:  (*Evaluator).evalBangOperatorExpression,
		token.MINUS: (*Evaluator).evalMinusPrefixOperatorExpression,
	}

	evalInfixExprFns = map[token.TokenType]evalInfixExprFn{
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
