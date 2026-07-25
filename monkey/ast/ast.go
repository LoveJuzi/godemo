package ast

import (
	"bytes"
	"monkey/token"
	"strings"
)

type NodeType int

const (
	NodeUnknown NodeType = iota
	NodeProgram
	NodeReturnStatement
	NodeLetStatement
	NodeExpressionStatement
	NodeIdentifier
	NodeBoolean
	NodeIfExpression
	NodeFunctionLiteral
	NodeCallExpression
	NodeBlockStatement
	NodePrefixExpression
	NodeInfixExpression
	NodeIntegerLiteral
)

type Node interface {
	TokenLiteral() string
	String() string
	Kind() NodeType
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

func NewProgram() *Program {
	return &Program{Statements: []Statement{}}
}

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}
func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}
func (p *Program) Kind() NodeType {
	return NodeProgram
}

func NewLetStatement(
	tokenObj token.Token,
	name *Identifier,
	value Expression) *LetStatement {
	return &LetStatement{
		Token: tokenObj,
		Name:  name,
		Value: value,
	}
}

type LetStatement struct {
	Token token.Token // token.LET 词法单元
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {}
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}
func (ls *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral())
	out.WriteString(" ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")

	if nil != ls.Value {
		out.WriteString(ls.Value.String())
	}

	out.WriteString(";")

	return out.String()
}
func (ls *LetStatement) Kind() NodeType {
	return NodeLetStatement
}

func NewReturnStatement(
	tokenObj token.Token,
	value Expression) *ReturnStatement {
	return &ReturnStatement{Token: tokenObj, ReturnValue: value}
}

type ReturnStatement struct {
	Token       token.Token // 'return'词法单元
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode() {}
func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}
func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral())
	out.WriteString(" ")

	if nil != rs.ReturnValue {
		out.WriteString(rs.ReturnValue.String())
	}

	out.WriteString(";")

	return out.String()
}
func (rs *ReturnStatement) Kind() NodeType {
	return NodeReturnStatement
}

func NewExpressionStatement(
	tokenObj token.Token,
	expression Expression) *ExpressionStatement {
	return &ExpressionStatement{Token: tokenObj, Expression: expression}
}

type ExpressionStatement struct {
	Token      token.Token // 表达式中的第一个词法单元
	Expression Expression
}

func (es *ExpressionStatement) statementNode() {}
func (es *ExpressionStatement) TokenLiteral() string {
	return es.Token.Literal
}
func (es *ExpressionStatement) String() string {
	if nil != es.Expression {
		return es.Expression.String()
	}
	return ""
}
func (es *ExpressionStatement) Kind() NodeType {
	return NodeExpressionStatement
}

//////////////////////////////////////////////////////////////////////

func NewIdentifier(tokenObj token.Token, value string) *Identifier {
	return &Identifier{Token: tokenObj, Value: value}
}

type Identifier struct {
	Token token.Token // token.IDENT词法单元
	Value string
}

func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}
func (i *Identifier) String() string { return i.Value }
func (i *Identifier) Kind() NodeType {
	return NodeIdentifier
}

func NewIntegerLiteral(tokenObj token.Token, value int64) *IntegerLiteral {
	return &IntegerLiteral{Token: tokenObj, Value: value}
}

type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (il *IntegerLiteral) expressionNode() {}
func (il *IntegerLiteral) TokenLiteral() string {
	return il.Token.Literal
}
func (il *IntegerLiteral) String() string {
	return il.Token.Literal
}
func (il *IntegerLiteral) Kind() NodeType {
	return NodeIntegerLiteral
}

func NewPrefixExpression(
	tokenObj token.Token,
	operator string,
	right Expression) *PrefixExpression {
	return &PrefixExpression{
		Token:    tokenObj,
		Operator: operator,
		Right:    right}
}

type PrefixExpression struct {
	Token    token.Token //前缀词法单元
	Operator string
	Right    Expression
}

func (pe *PrefixExpression) expressionNode() {}
func (pe *PrefixExpression) TokenLiteral() string {
	return pe.Token.Literal
}
func (pe *PrefixExpression) String() string {
	var out bytes.Buffer

	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())

	return out.String()
}
func (pe *PrefixExpression) Kind() NodeType {
	return NodePrefixExpression
}

func NewInfixExpression(
	tokenObj token.Token,
	left Expression,
	operator string,
	right Expression) *InfixExpression {
	return &InfixExpression{
		Token:    tokenObj,
		Left:     left,
		Operator: operator,
		Right:    right,
	}
}

type InfixExpression struct {
	Token    token.Token
	Left     Expression
	Operator string
	Right    Expression
}

func (ie *InfixExpression) expressionNode() {}
func (ie *InfixExpression) TokenLiteral() string {
	return ie.Token.Literal
}
func (ie *InfixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(ie.Operator)
	out.WriteString(" ")
	out.WriteString(ie.Left.String())
	out.WriteString(" ")
	out.WriteString(ie.Right.String())
	out.WriteString(")")

	return out.String()
}
func (ie *InfixExpression) Kind() NodeType {
	return NodeInfixExpression
}

func NewBoolean(tokenObj token.Token, value bool) *Boolean {
	return &Boolean{
		Token: tokenObj,
		Value: value,
	}
}

type Boolean struct {
	Token token.Token
	Value bool
}

func (b *Boolean) expressionNode() {}
func (b *Boolean) TokenLiteral() string {
	return b.Token.Literal
}
func (b *Boolean) String() string {
	return b.Token.Literal
}
func (b *Boolean) Kind() NodeType {
	return NodeBoolean
}

func NewIfExpression(
	tokenObj token.Token,
	condition Expression,
	consequence *BlockStatement,
	alternative *BlockStatement) *IfExpression {
	return &IfExpression{
		Token:       tokenObj,
		Condition:   condition,
		Consequence: consequence,
		Alternative: alternative,
	}
}

type IfExpression struct {
	Token       token.Token
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

func (ie *IfExpression) expressionNode() {}
func (ie *IfExpression) TokenLiteral() string {
	return ie.Token.Literal
}
func (ie *IfExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(if ")
	out.WriteString(ie.Condition.String())
	out.WriteString(" ")
	out.WriteString(ie.Consequence.String())

	if nil != ie.Alternative {
		out.WriteString(" else ")
		out.WriteString(ie.Alternative.String())
	}
	out.WriteString(")")

	return out.String()
}
func (ie *IfExpression) Kind() NodeType {
	return NodeIfExpression
}

func NewBlockStatement(
	tokenObj token.Token,
	statements []Statement) *BlockStatement {
	return &BlockStatement{Token: tokenObj, Statements: statements}
}

type BlockStatement struct {
	Token      token.Token
	Statements []Statement
}

func (bs *BlockStatement) statementNode() {}
func (bs *BlockStatement) TokenLiteral() string {
	return bs.Token.Literal
}
func (bs *BlockStatement) String() string {
	var out bytes.Buffer

	var statements = []string{}

	for _, s := range bs.Statements {
		statements = append(statements, s.String())
	}

	out.WriteString("(block ")
	out.WriteString(strings.Join(statements, " "))
	out.WriteString(")")

	return out.String()
}
func (bs *BlockStatement) Kind() NodeType {
	return NodeBlockStatement
}

func NewFunctionLiteral(
	tokenObj token.Token,
	parameters []*Identifier,
	body *BlockStatement) *FunctionLiteral {
	return &FunctionLiteral{
		Token:      tokenObj,
		Parameters: parameters,
		Body:       body,
	}
}

type FunctionLiteral struct {
	Token      token.Token
	Parameters []*Identifier
	Body       *BlockStatement
}

func (fl *FunctionLiteral) expressionNode() {}
func (fl *FunctionLiteral) TokenLiteral() string {
	return fl.Token.Literal
}
func (fl *FunctionLiteral) String() string {
	var out bytes.Buffer

	var params = []string{}
	for _, param := range fl.Parameters {
		params = append(params, param.String())
	}

	out.WriteString("fn ")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") ")
	out.WriteString(fl.Body.String())

	return out.String()
}
func (fl *FunctionLiteral) Kind() NodeType {
	return NodeFunctionLiteral
}

func NewCallExpression(
	tokenObj token.Token,
	function Expression,
	arguments []Expression) *CallExpression {
	return &CallExpression{
		Token:     tokenObj,
		Function:  function,
		Arguments: arguments,
	}
}

type CallExpression struct {
	Token     token.Token // '('词法单元
	Function  Expression  // 标识符或函数字面量
	Arguments []Expression
}

func (ce *CallExpression) expressionNode() {}
func (ce *CallExpression) TokenLiteral() string {
	return ce.Token.Literal
}
func (ce *CallExpression) String() string {
	var out bytes.Buffer

	var args = []string{}

	for _, a := range ce.Arguments {
		args = append(args, a.String())
	}

	out.WriteString(ce.Function.String())
	out.WriteString("(")
	out.WriteString(strings.Join(args, ", "))
	out.WriteString(")")

	return out.String()
}
func (ce *CallExpression) Kind() NodeType {
	return NodeCallExpression
}
