package lexer

import "fmt"

type Lexer struct {
	Pos       int
	Text      string
	CurrToken string
}

func (Lexer) GetNextToken() {
	fmt.Println("deu certo")
}
