package main

import (
	"ultrago/lexer"
	"ultrago/parser"
)

func main() {
	aLexer := lexer.Lexer{}
	aParser := parser.Parser{aLexer}
	aParser.Parse()
}
