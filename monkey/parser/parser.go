package parser

import (
	"fmt"
	"monkey/ast"
	"monkey/lexer"
	"monkey/token"
	"strconv"
)

func New(l *lexer.Lexer) *Parser {
	var p = &Parser{l: l, errors: []string{}}

	// 读取两个词法单元，以设置curToken和peekToken
	p.nextToken()
	p.nextToken()

	return p
}

const (
	_ int = iota
	LOWEST
	EQUALS      // ==
	LESSGREATER // > or <
	SUM         // +
	PRODUCT     // *
	PREFIX      // -X or !X
	CALL        // myFunction(X)
)

type Parser struct {
	l *lexer.Lexer

	errors []string

	curToken  token.Token
	peekToken token.Token
}

func (p *Parser) ParseProgram() *ast.Program {
	var program = &ast.Program{}

	program.Statements = []ast.Statement{}

	for token.EOF != p.curToken.Type {
		var stmt = p.parseStatement()
		if nil != stmt {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}

	return program
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.GetToken()
}

func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) checkPrecedence() int {
	if p, ok := precedences[p.peekToken.Type]; ok {
		return p
	}
	return LOWEST
}

func (p *Parser) curPrecedence() int {
	if p, ok := precedences[p.curToken.Type]; ok {
		return p
	}
	return LOWEST
}

func (p *Parser) peekPrecedence() int {
	if p, ok := precedences[p.peekToken.Type]; ok {
		return p
	}

	return LOWEST
}

func (p *Parser) peekError(t token.TokenType) {
	var msg = fmt.Sprintf("expected next token to be %s, got %s instead", t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}

func (p *Parser) expectNextToken(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		p.peekError(t)
		return false
	}
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	default:
		return p.parseExpressionStatement()
	}
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	var tokenObj = p.curToken

	if !p.expectNextToken(token.IDENT) {
		return nil
	}

	var name = ast.NewIdentifier(p.curToken, p.curToken.Literal)

	if !p.expectNextToken(token.ASSIGN) {
		return nil
	}

	p.nextToken()
	var value = p.parseExpression(LOWEST)

	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return ast.NewLetStatement(tokenObj, name, value)
}

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	var tokenObj = p.curToken

	p.nextToken()
	var returnValue = p.parseExpression(LOWEST)

	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return ast.NewReturnStatement(tokenObj, returnValue)
}

func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	var tokenObj = p.curToken
	var expression = p.parseExpression(LOWEST)

	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return ast.NewExpressionStatement(tokenObj, expression)
}

func (p *Parser) noPrefixParseFnError(t token.TokenType) {
	p.errors = append(p.errors, fmt.Sprintf("no prefix parse function for %s found", t))
}

func (p *Parser) parseExpression(precedence int) ast.Expression {
	var prefix = prefixParseFns[p.curToken.Type]
	if nil == prefix {
		p.noPrefixParseFnError(p.curToken.Type)
		return nil
	}

	var leftExp = prefix(p)

	for !p.peekTokenIs(token.SEMICOLON) && precedence < p.peekPrecedence() {
		var infix = infixParseFns[p.peekToken.Type]
		if nil == infix {
			return leftExp
		}
		p.nextToken() // skip peekToken

		leftExp = infix(p, leftExp)
	}

	return leftExp
}

func (p *Parser) parseIdentifier() ast.Expression {
	return ast.NewIdentifier(p.curToken, p.curToken.Literal)
}

func (p *Parser) parseIntegerLiteral() ast.Expression {
	var tokenObj = p.curToken

	var value = int64(0)
	if tmpVal, err := strconv.ParseInt(p.curToken.Literal, 0, 64); nil != err {
		p.errors = append(p.errors, fmt.Sprintf("could not parse %q as integer", p.curToken.Literal))
		return nil
	} else {
		value = tmpVal
	}

	return ast.NewIntegerLiteral(tokenObj, value)
}

func (p *Parser) parsePrefixExpression() ast.Expression {
	var tokenObj = p.curToken
	var operator = p.curToken.Literal
	p.nextToken()
	var right = p.parseExpression(PREFIX)
	return ast.NewPrefixExpression(tokenObj, operator, right)
}

func (p *Parser) parseBoolean() ast.Expression {
	return ast.NewBoolean(p.curToken, p.curTokenIs(token.TRUE))
}

func (p *Parser) parseGroupedExpression() ast.Expression {
	p.nextToken()

	var exp = p.parseExpression(LOWEST)

	if !p.expectNextToken(token.RPAREN) {
		return nil
	}

	return exp
}

func (p *Parser) parseIfExpression() ast.Expression {
	var tokenObj = p.curToken

	var condition = p.parseCondition()

	if nil == condition {
		return nil
	}

	var consequence = p.parseBlockStatement()

	if nil == consequence {
		return nil
	}

	var alternative = p.parseElseBlockStatement()

	return ast.NewIfExpression(tokenObj, condition, consequence, alternative)
}

func (p *Parser) parseCondition() ast.Expression {
	if !p.expectNextToken(token.LPAREN) {
		return nil
	}
	p.nextToken()

	var condition = p.parseExpression(LOWEST)

	if !p.expectNextToken(token.RPAREN) {
		return nil
	}

	return condition
}

func (p *Parser) parseBlockStatement() *ast.BlockStatement {
	if !p.expectNextToken(token.LBRACE) {
		return nil
	}

	var tokenObj = p.curToken

	var statements = []ast.Statement{}

	p.nextToken()
	for !p.curTokenIs(token.RBRACE) && !p.curTokenIs(token.EOF) {
		var stmt = p.parseStatement()
		if nil != stmt {
			statements = append(statements, stmt)
		}
		p.nextToken()
	}

	if p.curTokenIs(token.EOF) {
		p.peekError(token.RBRACE)
		return nil
	}

	return ast.NewBlockStatement(tokenObj, statements)
}

func (p *Parser) parseElseBlockStatement() *ast.BlockStatement {
	if !p.peekTokenIs(token.ELSE) {
		return nil
	}
	p.nextToken()

	return p.parseBlockStatement()
}

func (p *Parser) parseFunctionLiteral() ast.Expression {
	var tokenObj = p.curToken

	var parameters = p.parseParameters()
	if nil == parameters {
		return nil
	}

	var body = p.parseBlockStatement()

	if nil == body {
		return nil
	}

	return ast.NewFunctionLiteral(tokenObj, parameters, body)
}

func (p *Parser) parseParameters() []*ast.Identifier {
	if !p.expectNextToken(token.LPAREN) {
		return nil
	}

	var parameters = []*ast.Identifier{}

	for {
		if !p.peekTokenIs(token.IDENT) {
			break
		}
		p.nextToken()

		parameters = append(parameters,
			ast.NewIdentifier(p.curToken, p.curToken.Literal))

		if p.peekTokenIs(token.COMMA) {
			p.nextToken()
			if !p.peekTokenIs(token.IDENT) {
				p.peekError(token.IDENT)
				return nil
			}
		} else {
			break
		}
	}

	if !p.expectNextToken(token.RPAREN) {
		return nil
	}

	return parameters
}

func (p *Parser) parseInfixExpression(left ast.Expression) ast.Expression {
	var tokenObj = p.curToken
	var operator = p.curToken.Literal
	var precedence = p.curPrecedence()
	p.nextToken()
	var right = p.parseExpression(precedence)
	return ast.NewInfixExpression(tokenObj, left, operator, right)
}

func (p *Parser) parseCallExpression(function ast.Expression) ast.Expression {
	var tokenObj = p.curToken

	var arguments = p.parseCallArguments()

	return ast.NewCallExpression(tokenObj, function, arguments)
}

func (p *Parser) parseCallArguments() []ast.Expression {
	var args = []ast.Expression{}

	if p.peekTokenIs(token.RPAREN) {
		p.nextToken()
		return args
	}

	for {
		p.nextToken()
		args = append(args, p.parseExpression(LOWEST))

		if p.peekTokenIs(token.COMMA) {
			p.nextToken()
			if p.peekTokenIs(token.RPAREN) {
				return []ast.Expression{}
			}
			continue
		}

		if p.peekTokenIs(token.RPAREN) {
			p.nextToken()
			break
		}
	}

	return args
}

///////////////////////////////////////////////////////////////////////

var precedences = map[token.TokenType]int{
	token.EQ:       EQUALS,
	token.NOT_EQ:   EQUALS,
	token.LT:       LESSGREATER,
	token.GT:       LESSGREATER,
	token.PLUS:     SUM,
	token.MINUS:    SUM,
	token.SLASH:    PRODUCT,
	token.ASTERISK: PRODUCT,
	token.LPAREN:   CALL,
}

type (
	prefixParseFn func(p *Parser) ast.Expression
	infixParseFn  func(p *Parser, left ast.Expression) ast.Expression
)

var prefixParseFns map[token.TokenType]prefixParseFn
var infixParseFns map[token.TokenType]infixParseFn

func init() {
	prefixParseFns = map[token.TokenType]prefixParseFn{
		token.IDENT:    (*Parser).parseIdentifier,
		token.INT:      (*Parser).parseIntegerLiteral,
		token.BANG:     (*Parser).parsePrefixExpression,
		token.MINUS:    (*Parser).parsePrefixExpression,
		token.TRUE:     (*Parser).parseBoolean,
		token.FALSE:    (*Parser).parseBoolean,
		token.LPAREN:   (*Parser).parseGroupedExpression,
		token.IF:       (*Parser).parseIfExpression,
		token.FUNCTION: (*Parser).parseFunctionLiteral,
	}

	infixParseFns = map[token.TokenType]infixParseFn{
		token.PLUS:     (*Parser).parseInfixExpression,
		token.MINUS:    (*Parser).parseInfixExpression,
		token.SLASH:    (*Parser).parseInfixExpression,
		token.ASTERISK: (*Parser).parseInfixExpression,

		token.EQ:     (*Parser).parseInfixExpression,
		token.NOT_EQ: (*Parser).parseInfixExpression,
		token.LT:     (*Parser).parseInfixExpression,
		token.GT:     (*Parser).parseInfixExpression,

		token.LPAREN: (*Parser).parseCallExpression,
	}
}
