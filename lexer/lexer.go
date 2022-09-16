package lexer

import (
	"app/token"
	"io"
)

type Lexer struct {
	stream io.Reader
	line   int
	column int
}

func New(r io.Reader) *Lexer {
	return &Lexer{
		stream: r,
	}
}

func (l *Lexer) Line() int   { return l.line }
func (l *Lexer) Column() int { return l.column }

func (l *Lexer) GetNextToken() token.Token {
	return token.Token{}
}
