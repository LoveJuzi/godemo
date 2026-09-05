package parser

import (
	"fmt"
	"monkeyv2/ast"
	"monkeyv2/lexer"
	"monkeyv2/token"
)

type statementParser interface {
	run(p *Parser) ast.Statement
}

type letStatementParser struct{}

func (letStatementParser) run(p *Parser) ast.Statement {
	curToken := p.getToken()

	stmt := &ast.LetStatement{Token: curToken}

	curToken = p.getToken()
	if curToken.Type != token.IDENT {
		p.expectPeekError(curToken, token.IDENT)
		p.ungetToken(curToken)
		return nil
	}
	stmt.Name = &ast.Identifier{Token: curToken, Value: curToken.Literal}

	// TODO: 跳过对表达式的处理，知道遇见分号
	for {
		curToken = p.getToken()
		if curToken.Type == token.SEMICOLON {
			break
		}
	}

	return stmt
}

type illegalStatementParser struct{}

func (illegalStatementParser) run(p *Parser) ast.Statement {
	p.getToken()
	return nil
}

type Parser struct {
	l *lexer.Lexer

	errors []error

	buffer []token.Token
}

func NewParser(l *lexer.Lexer) *Parser {
	return &Parser{l: l}
}

func (t *Parser) Errors() []error {
	return t.errors
}

func (t *Parser) expectPeekError(curToken token.Token,
	tokenType token.TokenType) {
	msg := fmt.Errorf(
		"expected next token to be %s, got %s instead",
		tokenType,
		curToken.Type)
	t.errors = append(t.errors, msg)
}

func (t *Parser) ParserProgram() *ast.Program {
	program := &ast.Program{}

	for {
		curToken := t.getToken()
		if curToken.Type == token.EOF {
			break
		}
		t.ungetToken(curToken)
		stmt := t.parseStatement()
		if stmt == nil {
			continue
		}
		program.Statements = append(program.Statements, stmt)
	}

	return program
}

func (t *Parser) getToken() token.Token {
	if len(t.buffer) > 0 {
		last := t.buffer[len(t.buffer)-1]
		t.buffer = t.buffer[:len(t.buffer)-1]
		return last
	}
	return t.l.NextToken()
}

func (t *Parser) ungetToken(tokenObj token.Token) {
	t.buffer = append(t.buffer, tokenObj)
}

func (t *Parser) parseStatement() ast.Statement {
	return t.getStatementParser().run(t)
}

func (t *Parser) getStatementParser() statementParser {
	curToken := t.getToken()
	defer func() { t.ungetToken(curToken) }()

	switch curToken.Type {
	case token.LET:
		return letStatementParser{}
	default:
		return illegalStatementParser{}
	}
}
