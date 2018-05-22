package parser

import (
	"monkey/lexer"
	"monkey/token"
)

type Parser struct {
	l *lexer.Lexer

	curToken token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser{
	p := &Parser{l: l	}

	// Read two token, so curToken and peekToken are both set
	p.NextToken()
	p.NextToken()
	return p
}

func (p *Parser) nextToken(){
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParserProgram() *ast.Program {
	return nil
}
