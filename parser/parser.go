package parser

import (
	"app/lexer"
	"app/token"
	"fmt"
	"strconv"
)

type Parser struct {
	lexer   *lexer.Lexer
	current token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		lexer: l,
	}
	p.getNextToken()
	return p
}

func (p *Parser) getNextToken() {
	p.current = p.lexer.GetNextToken()
}

func (p *Parser) Parse() (map[string]interface{}, error) {
	out, err := p.parseObject()
	if err != nil {
		return nil, fmt.Errorf("at line %d:%d %w", p.lexer.Line(), p.lexer.Column(), err)
	}
	return out, nil
}

func (p *Parser) parseObject() (map[string]interface{}, error) {
	out := map[string]interface{}{}
	for p.current.Kind == token.KindIdentifier {
		key := p.current.Literal
		p.getNextToken()
		value, err := p.parseValue()
		if err != nil {
			return nil, fmt.Errorf("cannot parse value : %w", err)
		}
		out[key] = value
	}
	if p.current.Kind != token.KindRightParenthesis && p.current.Kind != token.KindEOF {
		return nil, fmt.Errorf("cannot parse object expected ')' got %v", p.current.Literal)
	}
	p.getNextToken()
	return out, nil
}

func (p *Parser) parseArray() ([]interface{}, error) {
	out := []interface{}{}
	for p.current.Kind != token.KindRightParenthesis {
		val, err := p.parseValue()
		if err != nil {
			return nil, err
		}
		out = append(out, val)
	}
	p.getNextToken()
	return out, nil
}

func (p *Parser) parseValue() (interface{}, error) {
	var value interface{}
	switch p.current.Kind {
	case token.KindString:
		value = p.current.Literal
		p.getNextToken()
	case token.KindInt:
		var err error
		value, err = strconv.ParseInt(p.current.Literal, 0, 64)
		if err != nil {
			return nil, fmt.Errorf("cannot parse int : %w", err)
		}
		p.getNextToken()
	case token.KindFloat:
		var err error
		value, err = strconv.ParseFloat(p.current.Literal, 64)
		if err != nil {
			return nil, fmt.Errorf("cannot parse float : %w", err)
		}
		p.getNextToken()
	case token.KindTrue:
		value = true
		p.getNextToken()
	case token.KindFalse:
		value = false
		p.getNextToken()
	case token.KindLeftParenthesis:
		p.getNextToken()
		var err error
		if p.current.Kind == token.KindIdentifier {
			value, err = p.parseObject()
			if err != nil {
				return nil, err
			}
		} else {
			value, err = p.parseArray()
			if err != nil {
				return nil, err
			}
		}
	default:
		return nil, fmt.Errorf("cannot parse '%s' as value", p.current.Literal)
	}
	return value, nil
}
