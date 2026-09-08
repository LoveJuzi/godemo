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
	run() ast.Statement
}

type letStatementParser struct{ p *Parser }

func (lsp letStatementParser) run() ast.Statement {
	curToken := lsp.p.getToken()

	stmt := &ast.LetStatement{Token: curToken}

	curToken = lsp.p.getToken()
	if curToken.Type != token.IDENT {
		lsp.p.expectPeekError(curToken, token.IDENT)
		lsp.p.ungetToken(curToken)
		return nil
	}
	stmt.Name = &ast.Identifier{Token: curToken, Value: curToken.Literal}

	stmt.Value = lsp.p.parseExpression(LOWEST)
	if lsp.p.hasErr() {
		return nil
	}

	curToken = lsp.p.getToken()
	if curToken.Type != token.SEMICOLON {
		lsp.p.ungetToken(curToken)
	}

	return stmt
}

type returnStatementParser struct{ p *Parser }

func (rsp returnStatementParser) run() ast.Statement {
	curToken := rsp.p.getToken()
	stmt := &ast.ReturnStatement{Token: curToken}

	stmt.ReturnValue = rsp.p.parseExpression(LOWEST)
	if rsp.p.hasErr() {
		return nil
	}

	curToken = rsp.p.getToken()
	if curToken.Type != token.SEMICOLON {
		rsp.p.ungetToken(curToken)
	}

	return stmt
}

type expressionStatementParser struct{ p *Parser }

func (esp expressionStatementParser) run() ast.Statement {
	curToken := esp.p.getToken()
	stmt := &ast.ExpressionStatement{Token: curToken}
	esp.p.ungetToken(curToken)
	stmt.Expression = esp.p.parseExpression(LOWEST)

	curToken = esp.p.getToken()
	if curToken.Type != token.SEMICOLON {
		esp.p.ungetToken(curToken)
	}

	if esp.p.hasErr() {
		return nil
	}

	return stmt
}

type blockStatementParser struct{ p *Parser }

func (bsp blockStatementParser) run() ast.Statement {
	return bsp.p.parseBlockStatement()
}

type blankStatmentParser struct{ p *Parser }

func (bsp blankStatmentParser) run() ast.Statement {
	bsp.p.getToken()
	return nil
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

type booleanParser struct{ p *Parser }

func (bp booleanParser) run() ast.Expression {
	curToken := bp.p.getToken()
	return &ast.Boolean{Token: curToken, Value: curToken.Type == token.TRUE}
}

type prefixExpressionParser struct{ p *Parser }

func (pep prefixExpressionParser) run() ast.Expression {
	curToken := pep.p.getToken()
	expression := &ast.PrefixExpression{
		Token:    curToken,
		Operator: curToken.Literal,
	}

	expression.Right = pep.p.parseExpression(PREFIX)

	return expression
}

type groupedExpressionParser struct{ p *Parser }

func (gep groupedExpressionParser) run() ast.Expression {
	curToken := gep.p.getToken()
	exp := gep.p.parseExpression(LOWEST)

	curToken = gep.p.getToken()
	if curToken.Type != token.RPAREN {
		gep.p.expectPeekError(curToken, token.RPAREN)
		gep.p.ungetToken(curToken)
		return nil
	}

	return exp
}

type ifExpressionParser struct{ p *Parser }

func (iep ifExpressionParser) run() ast.Expression {
	curToken := iep.p.getToken()
	ifexp := &ast.IfExpression{Token: curToken}

	curToken = iep.p.getToken()
	if curToken.Type != token.LPAREN {
		iep.p.expectPeekError(curToken, token.LPAREN)
		iep.p.ungetToken(curToken)
		return nil
	}

	ifexp.Condition = iep.p.parseExpression(LOWEST)

	curToken = iep.p.getToken()
	if curToken.Type != token.RPAREN {
		iep.p.expectPeekError(curToken, token.RPAREN)
		iep.p.ungetToken(curToken)
		return nil
	}

	curToken = iep.p.getToken()
	if curToken.Type != token.LBRACE {
		iep.p.expectPeekError(curToken, token.RPAREN)
		iep.p.ungetToken(curToken)
		return nil
	}
	iep.p.ungetToken(curToken)

	ifexp.Consequence = iep.p.parseBlockStatement()

	curToken = iep.p.getToken()
	if curToken.Type != token.ELSE {
		iep.p.ungetToken(curToken)
		return ifexp
	}

	curToken = iep.p.getToken()
	if curToken.Type != token.LBRACE {
		iep.p.expectPeekError(curToken, token.RPAREN)
		iep.p.ungetToken(curToken)
		return nil
	}
	iep.p.ungetToken(curToken)

	ifexp.Alternative = iep.p.parseBlockStatement()

	if iep.p.hasErr() {
		return nil
	}

	return ifexp
}

type functionLiteralParser struct{ p *Parser }

func (fll functionLiteralParser) run() ast.Expression {
	curToken := fll.p.getToken()
	fnexp := &ast.FunctionLiteral{Token: curToken}

	curToken = fll.p.getToken()
	if curToken.Type != token.LPAREN {
		fll.p.expectPeekError(curToken, token.LPAREN)
		fll.p.ungetToken(curToken)
		return nil
	}
	fll.p.ungetToken(curToken)

	fnexp.Parameters = fll.p.parseParameters()

	if fll.p.hasErr() {
		return nil
	}

	curToken = fll.p.getToken()
	if curToken.Type != token.LBRACE {
		fll.p.expectPeekError(curToken, token.RPAREN)
		fll.p.ungetToken(curToken)
		return nil
	}
	fll.p.ungetToken(curToken)

	fnexp.Body = fll.p.parseBlockStatement()

	if fll.p.hasErr() {
		return nil
	}

	return fnexp
}

type illegalPrefixParser struct{ p *Parser }

func (ipp illegalPrefixParser) run() ast.Expression {
	curToken := ipp.p.getToken()
	ipp.p.noPrefixParserError(curToken)
	return nil
}

type infixParser interface {
	run(ast.Expression) ast.Expression
}

type infixExpressionParser struct{ p *Parser }

func (iep infixExpressionParser) run(left ast.Expression) ast.Expression {
	curToken := iep.p.getToken()

	expression := &ast.InfixExpression{
		Token:    curToken,
		Operator: curToken.Literal,
		Left:     left,
	}

	precedence := iep.p.getPrecedence(curToken)
	expression.Right = iep.p.parseExpression(precedence)

	return expression
}

type callExpressionParser struct{ p *Parser }

func (cep callExpressionParser) run(function ast.Expression) ast.Expression {
	curToken := cep.p.getToken()
	exp := &ast.CallExpression{Token: curToken, Function: function}

	cep.p.ungetToken(curToken)
	exp.Arguments = cep.p.parseArguments()

	if cep.p.hasErr() {
		return nil
	}

	return exp
}

type illegalInfixParser struct{ p *Parser }

func (iip illegalInfixParser) run(left ast.Expression) ast.Expression {
	return nil
}

type Parser struct {
	l *lexer.Lexer

	errors []error

	err error

	buffer []token.Token
}

func NewParser(l *lexer.Lexer) *Parser {
	return &Parser{l: l}
}

func (p *Parser) Errors() []error {
	return p.errors
}

func (p *Parser) hasErr() bool {
	return p.err != nil
}

func (p *Parser) resetErr() {
	if p.hasErr() {
		p.errors = append(p.errors, p.err)
	}
	p.err = nil
}

func (p *Parser) expectPeekError(curToken token.Token,
	tokenType token.TokenType) {
	p.err = fmt.Errorf(
		"expected next token to be %s, got %s instead",
		tokenType,
		curToken.Type)
}

func (p *Parser) integerLiteralParserError(curToken token.Token) {
	p.err = fmt.Errorf("could not parser %q as integer", curToken.Literal)
}

func (p *Parser) noPrefixParserError(curToken token.Token) {
	p.err = fmt.Errorf("no prefix parse function for %s found", curToken.Type)
}

func (p *Parser) ParserProgram() *ast.Program {
	program := &ast.Program{}

	for {
		p.resetErr() // 保存错误信息
		curToken := p.getToken()
		if curToken.Type == token.EOF {
			break
		}
		p.ungetToken(curToken)
		stmt := p.parseStatement()
		if stmt == nil {
			continue
		}
		program.Statements = append(program.Statements, stmt)
	}

	return program
}

func (p *Parser) getToken() token.Token {
	if len(p.buffer) > 0 {
		last := p.buffer[len(p.buffer)-1]
		p.buffer = p.buffer[:len(p.buffer)-1]
		return last
	}
	return p.l.NextToken()
}

func (p *Parser) ungetToken(curToken token.Token) {
	p.buffer = append(p.buffer, curToken)
}

func (p *Parser) getPrecedence(curToken token.Token) int {
	switch curToken.Type {
	case token.EQ, token.NOT_EQ:
		return EQUALS
	case token.LT, token.GT:
		return LESSGREATER
	case token.PLUS, token.MINUS:
		return SUM
	case token.SLASH, token.ASTERISK:
		return PRODUCT
	default:
		return LOWEST
	}
}

func (p *Parser) parseStatement() ast.Statement {
	return p.getStatementParser().run()
}

func (p *Parser) getStatementParser() statementParser {
	curToken := p.getToken()
	defer func() { p.ungetToken(curToken) }()

	switch curToken.Type {
	case token.LET:
		return letStatementParser{p: p}
	case token.RETURN:
		return returnStatementParser{p: p}
	case token.LBRACE:
		return blockStatementParser{p: p}
	case token.SEMICOLON:
		return blankStatmentParser{p: p}
	default:
		return expressionStatementParser{p: p}
	}
}

func (p *Parser) parseExpression(precedence int) ast.Expression {
	leftExp := p.getPrefixParser().run()
	if leftExp == nil {
		return nil
	}

	for {
		curToken := p.getToken()
		if curToken.Type == token.SEMICOLON {
			break
		}
		p.ungetToken(curToken)
		if precedence >= p.getPrecedence(curToken) {
			break
		}

		newExp := p.getInfixParser().run(leftExp)
		if newExp == nil {
			return nil
		}
		leftExp = newExp
	}

	return leftExp
}

func (p *Parser) getPrefixParser() prefixParser {
	curToken := p.getToken()
	defer func() { p.ungetToken(curToken) }()

	switch curToken.Type {
	case token.IDENT:
		return identifierParser{p: p}
	case token.INT:
		return integerLiteralParser{p: p}
	case token.TRUE, token.FALSE:
		return booleanParser{p: p}
	case token.BANG, token.MINUS:
		return prefixExpressionParser{p: p}
	case token.LPAREN:
		return groupedExpressionParser{p: p}
	case token.IF:
		return ifExpressionParser{p: p}
	case token.FUNCTION:
		return functionLiteralParser{p: p}
	default:
		return illegalPrefixParser{p: p}
	}
}

func (p *Parser) getInfixParser() infixParser {
	curToken := p.getToken()
	defer func() { p.ungetToken(curToken) }()

	switch curToken.Type {
	case token.EQ, token.NOT_EQ,
		token.LT, token.GT,
		token.PLUS, token.MINUS,
		token.SLASH, token.ASTERISK:
		return infixExpressionParser{p: p}
	default:
		return illegalInfixParser{p: p}
	}
}

func (p *Parser) parseParameters() []*ast.Identifier {
	curToken := p.getToken()

	idents := []*ast.Identifier{}

	for {
		curToken = p.getToken()
		if curToken.Type == token.EOF {
			p.ungetToken(curToken)
			break
		}
		if curToken.Type == token.RPAREN {
			return idents
		}

		if curToken.Type != token.IDENT {
			p.ungetToken(curToken)
			break
		}

		ident := &ast.Identifier{Token: curToken, Value: curToken.Literal}

		idents = append(idents, ident)

		curToken = p.getToken()
		if curToken.Type == token.COMMA {
			continue
		}
		if curToken.Type == token.RPAREN {
			return idents
		}
		p.ungetToken(curToken)
		break
	}
	p.expectPeekError(curToken, token.RPAREN)

	return nil
}

func (p *Parser) parseArguments() []ast.Expression {
	curToken := p.getToken()

	args := []ast.Expression{}

	for {
		p.resetErr()
		curToken = p.getToken()
		if curToken.Type == token.EOF {
			p.ungetToken(curToken)
			break
		}

		if curToken.Type == token.RPAREN {
			return args
		}

		p.ungetToken(curToken)

		arg := p.parseExpression(LOWEST)
		if p.hasErr() {
			p.resetErr()
			break
		}
		args = append(args, arg)

		curToken = p.getToken()
		if curToken.Type == token.COMMA {
			continue
		}
		if curToken.Type == token.RPAREN {
			return args
		}
		p.ungetToken(curToken)
		break
	}

	p.expectPeekError(curToken, token.RPAREN)
	return nil
}

func (p *Parser) parseBlockStatement() *ast.BlockStatement {
	curToken := p.getToken()
	block := &ast.BlockStatement{Token: curToken}

	for {
		p.resetErr()
		curToken = p.getToken()
		if curToken.Type == token.EOF {
			p.ungetToken(curToken)
			break
		}
		if curToken.Type == token.RBRACE {
			return block
		}
		p.ungetToken(curToken)

		stmt := p.parseStatement()
		if stmt != nil {
			block.Statements = append(block.Statements, stmt)
		}
	}
	p.expectPeekError(curToken, token.RBRACE)

	return nil
}
