package parser

import (
	"fmt"
	"monkeyv2/ast"
	"monkeyv2/lexer"
	"monkeyv2/token"
	"strconv"
)

const (
	_ int = iota
	LOWEST
	EQUALS
	LESSGREATER
	SUM
	PRODUCT
	PREFIX
	CALL
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

type returnStatementParser struct{}

func (returnStatementParser) run(p *Parser) ast.Statement {
	curToken := p.getToken()
	stmt := &ast.ReturnStatement{Token: curToken}

	// TODO: 跳过对表达式的处理，知道遇见分号
	for {
		curToken = p.getToken()
		if curToken.Type == token.SEMICOLON {
			break
		}
	}

	return stmt
}

type expressionStatementParser struct{}

func (expressionStatementParser) run(p *Parser) ast.Statement {
	curToken := p.getToken()
	stmt := &ast.ExpressionStatement{Token: curToken}
	stmt.Expression = p.parseExpression(LOWEST)

	curToken = p.getToken()
	if curToken.Type != token.SEMICOLON {
		p.ungetToken(curToken)
	}

	return stmt
}

type illegalStatementParser struct{}

func (illegalStatementParser) run(p *Parser) ast.Statement {
	p.getToken()
	return nil
}

type prefixParser interface {
	run() ast.Expression
}

type identifierParser struct{ p *Parser }

func (ip identifierParser) run() ast.Expression {
	curToken := ip.p.getToken()
	return &ast.Identifier{Token: curToken, Value: curToken.Literal}
}

type integerLiteralParser struct{ p *Parser }

func (ill integerLiteralParser) run() ast.Expression {
	curToken := ill.p.getToken()
	lit := &ast.IntegerLiteral{Token: curToken}

	value, err := strconv.ParseInt(curToken.Literal, 0, 64)
	if err != nil {
		ill.p.integerLiteralParserError(curToken)
		return nil
	}
	lit.Value = value
	return lit
}

type illegalPrefixParser struct{ p *Parser }

func (ipp illegalPrefixParser) run() ast.Expression {
	ipp.p.getToken()
	return nil
}

type infixParser interface {
	run(ast.Expression) ast.Expression
}

type Parser struct {
	l *lexer.Lexer

	errors []error

	buffer []token.Token
}

func NewParser(l *lexer.Lexer) *Parser {
	return &Parser{l: l}
}

func (p *Parser) Errors() []error {
	return p.errors
}

func (p *Parser) expectPeekError(curToken token.Token,
	tokenType token.TokenType) {
	msg := fmt.Errorf(
		"expected next token to be %s, got %s instead",
		tokenType,
		curToken.Type)
	p.errors = append(p.errors, msg)
}

func (p *Parser) integerLiteralParserError(curToken token.Token) {
	msg := fmt.Errorf("could not parser %q as integer", curToken.Literal)
	p.errors = append(p.errors, msg)
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
	case token.RETURN:
		return returnStatementParser{}
	default:
		return expressionStatementParser{}
	}
}

func (p *Parser) parseExpression(precedence int) ast.Expression {
	leftExp := p.getPrefixParser().run()
	return leftExp
}

func (p *Parser) getPrefixParser() prefixParser {
	curToken := p.getToken()
	switch curToken.Type {
	case token.IDENT:
		return identifierParser{p: p}
	case token.INT:
		return integerLiteralParser{p: p}
	default:
		return illegalPrefixParser{p: p}
	}
}
