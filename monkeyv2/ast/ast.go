package ast

import "monkeyv2/token"

type Node interface {
	Output() string
}

type Statement = Node
type Expression = Node

type Program struct {
	Statements []Statement
}

func (t *Program) Output() string {
	if len(t.Statements) > 0 {
		return t.Statements[0].Output()
	} else {
		return ""
	}
}

type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (t *LetStatement) Output() string {
	return t.Token.Literal
}

type Identifier struct {
	Token token.Token
	Value string
}

func (t *Identifier) Output() string {
	return t.Token.Literal
}
