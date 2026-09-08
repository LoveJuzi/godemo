package ast

import (
	"bytes"
	"fmt"
	"monkeyv2/token"
	"strings"
)

type Node interface {
	TokenLiteral() string
	String() string
}

type Statement = Node
type Expression = Node

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].String()
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

type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

func (ls *LetStatement) String() string {
	var out bytes.Buffer

	fmt.Fprintf(&out, "%s %s = ", ls.TokenLiteral(), ls.Name.String())

	if ls.Value != nil {
		fmt.Fprintf(&out, "%s", ls.Value.String())
	}

	fmt.Fprintf(&out, ";")

	return out.String()
}

type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}

func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	fmt.Fprintf(&out, "%s ", rs.TokenLiteral())

	if rs.ReturnValue != nil {
		fmt.Fprintf(&out, "%s", rs.ReturnValue.String())
	}

	fmt.Fprintf(&out, ";")

	return out.String()
}

type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

func (es *ExpressionStatement) TokenLiteral() string {
	return es.Token.Literal
}

func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}

type Identifier struct {
	Token token.Token
	Value string
}

func (ident *Identifier) TokenLiteral() string {
	return ident.Token.Literal
}

func (ident *Identifier) String() string {
	return ident.Value
}

type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (il *IntegerLiteral) TokenLiteral() string { return il.Token.Literal }
func (il *IntegerLiteral) String() string       { return il.Token.Literal }

type PrefixExpression struct {
	Token    token.Token
	Operator string
	Right    Expression
}

func (pe *PrefixExpression) TokenLiteral() string { return pe.Token.Literal }
func (pe *PrefixExpression) String() string {
	var out bytes.Buffer

	fmt.Fprintf(&out, "(%s%s)", pe.Operator, pe.Right.String())

	return out.String()
}

type Boolean struct {
	Token token.Token
	Value bool
}

func (b *Boolean) TokenLiteral() string { return b.Token.Literal }
func (b *Boolean) String() string       { return b.Token.Literal }

type InfixExpression struct {
	Token    token.Token
	Left     Expression
	Operator string
	Right    Expression
}

func (ie *InfixExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *InfixExpression) String() string {
	var out bytes.Buffer

	fmt.Fprintf(&out, "(%s %s %s)",
		ie.Left.String(), ie.Operator, ie.Right.String())

	return out.String()
}

type IfExpression struct {
	Token       token.Token
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

func (ie *IfExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *IfExpression) String() string {
	var out bytes.Buffer

	fmt.Fprintf(&out, "if %s %s ",
		ie.Condition.String(), ie.Consequence.String())

	if ie.Alternative != nil {
		fmt.Fprintf(&out, "else %s", ie.Alternative.String())
	}

	return out.String()
}

type FunctionLiteral struct {
	Token      token.Token
	Parameters []*Identifier
	Body       *BlockStatement
}

func (fl *FunctionLiteral) TokenLiteral() string { return fl.Token.Literal }
func (fl *FunctionLiteral) String() string {
	var out bytes.Buffer

	params := []string{}

	for _, p := range fl.Parameters {
		params = append(params, p.String())
	}

	fmt.Fprintf(&out, "%s(%s) %s",
		fl.Token.Literal,
		strings.Join(params, ", "),
		fl.Body.String())

	return out.String()
}

type CallExpression struct {
	Token     token.Token
	Function  Expression
	Arguments []Expression
}

func (ce CallExpression) TokenLiteral() string { return ce.Token.Literal }
func (ce CallExpression) String() string {
	var out bytes.Buffer

	args := []string{}

	for _, a := range ce.Arguments {
		args = append(args, a.String())
	}

	fmt.Fprintf(&out, "%s(%s)",
		ce.Function.String(),
		strings.Join(args, ", "))

	return out.String()
}

type BlockStatement struct {
	Token      token.Token
	Statements []Statement
}

func (bs *BlockStatement) TokenLiteral() string { return bs.Token.Literal }
func (bs *BlockStatement) String() string {
	var out bytes.Buffer

	for _, s := range bs.Statements {
		fmt.Fprintf(&out, "%s", s)
	}

	return out.String()
}
