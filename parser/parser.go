package parser

import (
	"ultrago/lexer"
)

type Parser struct {
	LexerInstance lexer.Lexer
}

func (self Parser) Parse() {
	self.LexerInstance.GetNextToken()
}
