package parser2

import (
	"fmt"
	"monkey/ast"
	"monkey/lexer"
	"monkey/token"
	"strconv"
)

func NewParser(l *lexer.Lexer) *Parser {
	return &Parser{l: l, errors: []string{}}
}

type Parser struct {
	l *lexer.Lexer

	errors []string
}

func (p *Parser) ParseProgram() *ast.Program {
	var program = ast.NewProgram()

	for {
		var tokType = p.peekTokenType()
		if token.SEMICOLON == tokType {
			p.l.GetToken()
			continue
		}
		if token.EOF == tokType {
			p.l.GetToken()
			break
		}
		var stmt = p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
	}

	return program
}

func (p *Parser) Errors() []string { return p.errors }

func (p *Parser) peekTokenType() token.TokenType {
	var tok = p.l.GetToken()
	defer p.l.UngetToken(tok)
	return tok.Type
}

func (p *Parser) parseStatement() ast.Statement {
	var tokType = p.getStatementType(p.peekTokenType())
	if fn, ok := parseStatementFns[tokType]; ok {
		return fn(p)
	}
	panic(fmt.Sprintf("unknown statement type %v", tokType))
}

func (p *Parser) getStatementType(tokType token.TokenType) token.TokenType {
	if token.LET == tokType {
		return token.LET
	}
	if token.RETURN == tokType {
		return token.RETURN
	}
	return token.EXPRESSION
}

func (p *Parser) parseLetStatement() ast.Statement {
	return nil
}

func (p *Parser) parseReturnStatement() ast.Statement {
	return nil
}

func (p *Parser) parseExpressionStatement() ast.Statement {
	var tok = token.New(token.EXPRESSION, "")
	var expression = p.parseExpression(LOWEST)
	return ast.NewExpressionStatement(tok, expression)
}

func (p *Parser) parseExpression(precedence int) ast.Expression {
	var tokType = p.peekTokenType()
	if fn, ok := prefixParseFns[tokType]; ok {
		var leftExp = fn(p)

		leftExp = p.parsePrecedence(leftExp, precedence)

		return leftExp
	}
	p.errors = append(p.errors, fmt.Sprintf("no prefix parse function for %s found", tokType))
	return nil
}

func (p *Parser) parsePrecedence(leftExp ast.Expression, precedence int) ast.Expression {
	for {
		var tokType = p.peekTokenType()
		if token.SEMICOLON == tokType || precedence >= precedences[tokType] {
			break
		}
		if fn, ok := infixParseFns[tokType]; ok {
			leftExp = fn(p, leftExp)
			continue
		}
		break
	}

	return leftExp
}

func (p *Parser) parseIdentifier() ast.Expression {
	var tok = p.l.GetToken()
	return ast.NewIdentifier(tok, tok.Literal)
}

func (p *Parser) parseIntegerLiteral() ast.Expression {
	var tok = p.l.GetToken()

	var value = int64(0)

	if tmpVal, ok := strconv.ParseInt(tok.Literal, 0, 64); nil == ok {
		value = tmpVal
		return ast.NewIntegerLiteral(tok, value)
	}

	p.errors = append(p.errors, fmt.Sprintf("could not parse %q as integer", tok.Literal))
	return nil
}

func (p *Parser) parsePrefixExpression() ast.Expression {
	var tok = p.l.GetToken()
	var operator = tok.Literal
	var right = p.parseExpression(PREFIX)
	return ast.NewPrefixExpression(tok, operator, right)
}

func (p *Parser) parseBoolean() ast.Expression {
	var tok = p.l.GetToken()
	return ast.NewBoolean(tok, token.TRUE == tok.Type)
}

func (p *Parser) parseGroupedExpression() ast.Expression {
	var btok = p.l.GetToken()
	if token.LPAREN != btok.Type {
		p.l.UngetToken(btok)
		p.errors = append(p.errors, fmt.Sprintf("expected LPAREN as begin token"))
		return nil
	}

	var exp = p.parseExpression(LOWEST)

	var etok = p.l.GetToken()
	if token.RPAREN != etok.Type {
		p.l.UngetToken(etok)
		p.errors = append(p.errors, fmt.Sprintf("expected RPAREN as end token"))
		return nil
	}

	return exp
}

func (p *Parser) parseIfExpression() ast.Expression {
	var tok = p.l.GetToken()

	var condition = p.parseCondition()

	if nil == condition {
		return nil
	}

	var consequence = p.parseBlockStatement()

	if nil == consequence {
		return nil
	}

	var alternative = p.parseElseBlockStatement()

	return ast.NewIfExpression(tok, condition, consequence, alternative)
}

func (p *Parser) parseCondition() ast.Expression {
	var btok = p.l.GetToken()
	if token.LPAREN != btok.Type {
		p.l.UngetToken(btok)
		p.errors = append(p.errors, fmt.Sprintf("expected LPAREN as begin token"))
		return nil
	}

	var condition = p.parseExpression(LOWEST)

	var etok = p.l.GetToken()
	if token.RPAREN != etok.Type {
		p.l.UngetToken(etok)
		p.errors = append(p.errors, fmt.Sprintf("expected RPAREN as end token"))
		return nil
	}

	return condition
}

func (p *Parser) parseBlockStatement() *ast.BlockStatement {
	var btok = p.l.GetToken()
	if token.LBRACE != btok.Type {
		p.errors = append(p.errors, fmt.Sprintf("expected LBRACE as begin token"))
		return nil
	}

	var statements = []ast.Statement{}

	for {
		var tokType = p.peekTokenType()
		if token.RBRACE == tokType || token.EOF == tokType {
			break
		}
		var stmt = p.parseStatement()
		if nil == stmt {
			continue
		}
		statements = append(statements, stmt)
	}

	var etok = p.l.GetToken()
	if token.RBRACE != etok.Type {
		p.errors = append(p.errors, fmt.Sprintf("expected RBRACE as end token"))
		return nil
	}

	return ast.NewBlockStatement(btok, statements)
}

func (p *Parser) parseElseBlockStatement() *ast.BlockStatement {
	var tok = p.l.GetToken()
	if token.ELSE != tok.Type {
		p.l.UngetToken(tok)
		return nil
	}

	return p.parseBlockStatement()
}

func (p *Parser) parseInfixExpression(left ast.Expression) ast.Expression {
	var tok = p.l.GetToken()
	var operator = tok.Literal
	var precedence = precedences[tok.Type]
	var right = p.parseExpression(precedence)
	return ast.NewInfixExpression(tok, left, operator, right)
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

type (
	parseStatementFn func(p *Parser) ast.Statement
	prefixParseFn    func(p *Parser) ast.Expression
	infixParseFn     func(p *Parser, left ast.Expression) ast.Expression
)

var (
	parseStatementFns map[token.TokenType]parseStatementFn
	prefixParseFns    map[token.TokenType]prefixParseFn
	infixParseFns     map[token.TokenType]infixParseFn
)

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

func init() {
	parseStatementFns = map[token.TokenType]parseStatementFn{
		token.LET:        (*Parser).parseLetStatement,
		token.RETURN:     (*Parser).parseReturnStatement,
		token.EXPRESSION: (*Parser).parseExpressionStatement,
	}

	prefixParseFns = map[token.TokenType]prefixParseFn{
		token.IDENT:  (*Parser).parseIdentifier,
		token.INT:    (*Parser).parseIntegerLiteral,
		token.BANG:   (*Parser).parsePrefixExpression,
		token.MINUS:  (*Parser).parsePrefixExpression,
		token.TRUE:   (*Parser).parseBoolean,
		token.FALSE:  (*Parser).parseBoolean,
		token.LPAREN: (*Parser).parseGroupedExpression,
		token.IF:     (*Parser).parseIfExpression,
		// token.FUNCTION: (*Parser).parseFunctionLiteral,
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

		// token.LPAREN: (*Parser).parseCallExpression,
	}
}
