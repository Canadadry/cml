package lexer

import (
	"app/token"
	"io"
)

type Lexer struct {
	stream io.Reader
	ch     [2]byte
	line   int
	column int
}

func New(r io.Reader) *Lexer {
	l := &Lexer{
		stream: r,
		line:   1,
		column: 1,
		ch:     [2]byte{0, 0},
	}
	l.readChar()
	l.readChar()
	return l
}

func (l *Lexer) Line() int   { return l.line }
func (l *Lexer) Column() int { return l.column }

func (l *Lexer) GetNextToken() token.Token {
	tok := token.Token{
		Literal: string(l.ch[0]),
	}
	switch l.ch[0] {
	case '"':
		l.readChar()
		tok.Kind = token.KindString
		tok.Literal = l.readString()
	case '(':
		tok.Kind = token.KindLeftParenthesis
		l.readChar()
	default:
		switch {
		case isLetter(l.ch[0]):
			tok.Kind = token.KindIdentifier
			tok.Literal = l.readIdentifier()
		default:
			tok.Kind = token.KindInvalid
		}
	}
	return tok
}

func (l *Lexer) readChar() {

	l.column += 1

	if l.ch[0] == '\n' {
		l.column = 1
		l.line += 1
	}

	read := []byte{0}
	n, err := l.stream.Read(read)
	if err != nil || n != 1 {
		read[0] = 0
	}

	l.ch[0] = l.ch[1]
	l.ch[1] = read[0]
}

func isWhiteSpace(ch byte) bool {
	return ch == ' ' || ch == '\t' || ch == '\r' || ch == '\n'
}

func isLetter(ch byte) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || ch == '_'
}

func isNumeric(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

func (l *Lexer) readIdentifier() string {
	litteral := ""
	for isLetter(l.ch[0]) || isNumeric(l.ch[0]) {
		litteral += string(l.ch[0])
		l.readChar()
	}
	return litteral
}

func (l *Lexer) readString() string {
	litteral := ""
	for l.ch[0] != '"' && l.ch[0] != 0 {
		litteral += string(l.ch[0])
		l.readChar()
	}
	l.readChar()
	return litteral
}
