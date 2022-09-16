package lexer

import (
	"app/token"
	"io"
	"strings"
	"testing"
)

func TestLexer(t *testing.T) {
	tests := map[string]struct {
		in  io.Reader
		out []token.Token
	}{
		"minified sample example": {
			in: strings.NewReader(`io_mode"async"service(http(web_proxy(listen_addr"127.0.0.1:8080"process(main(command("/usr/local/bin/awesome-srv""server"))mgmt(command("/usr/local/bin/awesome-mgmt""mgmt"))))))`),
			out: []token.Token{
				{Kind: token.KindIdentifier, Literal: "io_mode"},
				{Kind: token.KindString, Literal: "async"},
				{Kind: token.KindIdentifier, Literal: "service"},
				{Kind: token.KindLeftParenthesis, Literal: "("},
				{Kind: token.KindIdentifier, Literal: "http"},
				{Kind: token.KindLeftParenthesis, Literal: "("},
				{Kind: token.KindIdentifier, Literal: "web_proxy"},
				{Kind: token.KindLeftParenthesis, Literal: "("},
				{Kind: token.KindIdentifier, Literal: "listen_addr"},
				{Kind: token.KindString, Literal: "127.0.0.1:8080"},
				{Kind: token.KindIdentifier, Literal: "process"},
				{Kind: token.KindLeftParenthesis, Literal: "("},
				{Kind: token.KindIdentifier, Literal: "main"},
				{Kind: token.KindLeftParenthesis, Literal: "("},
				{Kind: token.KindIdentifier, Literal: "command"},
				{Kind: token.KindLeftParenthesis, Literal: "("},
				{Kind: token.KindString, Literal: "/usr/local/bin/awesome-srv"},
				{Kind: token.KindString, Literal: "server"},
				{Kind: token.KindRightParenthesis, Literal: ")"},
				{Kind: token.KindRightParenthesis, Literal: ")"},
				{Kind: token.KindIdentifier, Literal: "mgmt"},
				{Kind: token.KindLeftParenthesis, Literal: "("},
				{Kind: token.KindIdentifier, Literal: "command"},
				{Kind: token.KindLeftParenthesis, Literal: "("},
				{Kind: token.KindString, Literal: "/usr/local/bin/awesome-mgmt"},
				{Kind: token.KindString, Literal: "mgmt"},
				{Kind: token.KindRightParenthesis, Literal: ")"},
				{Kind: token.KindRightParenthesis, Literal: ")"},
				{Kind: token.KindRightParenthesis, Literal: ")"},
				{Kind: token.KindRightParenthesis, Literal: ")"},
				{Kind: token.KindRightParenthesis, Literal: ")"},
				{Kind: token.KindRightParenthesis, Literal: ")"},
				{Kind: token.KindEOF, Literal: "\x00"},
			},
		},
		"full sample example": {
			in: strings.NewReader(`io_mode "async"
			service (
				http (
					web_proxy (
						listen_addr "127.0.0.1:8080"
						process (
							main (
								command ("/usr/local/bin/awesome-srv" "server")
							)
							mgmt (
								command ("/usr/local/bin/awesome-mgmt" "mgmt")
							)
						)
					)
				)
			)`),
			out: []token.Token{
				{Kind: token.KindIdentifier, Literal: "io_mode"},
				{Kind: token.KindString, Literal: "async"},
				{Kind: token.KindIdentifier, Literal: "service"},
				{Kind: token.KindLeftParenthesis, Literal: "("},
				{Kind: token.KindIdentifier, Literal: "http"},
				{Kind: token.KindLeftParenthesis, Literal: "("},
				{Kind: token.KindIdentifier, Literal: "web_proxy"},
				{Kind: token.KindLeftParenthesis, Literal: "("},
				{Kind: token.KindIdentifier, Literal: "listen_addr"},
				{Kind: token.KindString, Literal: "127.0.0.1:8080"},
				{Kind: token.KindIdentifier, Literal: "process"},
				{Kind: token.KindLeftParenthesis, Literal: "("},
				{Kind: token.KindIdentifier, Literal: "main"},
				{Kind: token.KindLeftParenthesis, Literal: "("},
				{Kind: token.KindIdentifier, Literal: "command"},
				{Kind: token.KindLeftParenthesis, Literal: "("},
				{Kind: token.KindString, Literal: "/usr/local/bin/awesome-srv"},
				{Kind: token.KindString, Literal: "server"},
				{Kind: token.KindRightParenthesis, Literal: ")"},
				{Kind: token.KindRightParenthesis, Literal: ")"},
				{Kind: token.KindIdentifier, Literal: "mgmt"},
				{Kind: token.KindLeftParenthesis, Literal: "("},
				{Kind: token.KindIdentifier, Literal: "command"},
				{Kind: token.KindLeftParenthesis, Literal: "("},
				{Kind: token.KindString, Literal: "/usr/local/bin/awesome-mgmt"},
				{Kind: token.KindString, Literal: "mgmt"},
				{Kind: token.KindRightParenthesis, Literal: ")"},
				{Kind: token.KindRightParenthesis, Literal: ")"},
				{Kind: token.KindRightParenthesis, Literal: ")"},
				{Kind: token.KindRightParenthesis, Literal: ")"},
				{Kind: token.KindRightParenthesis, Literal: ")"},
				{Kind: token.KindRightParenthesis, Literal: ")"},
				{Kind: token.KindEOF, Literal: "\x00"},
			},
		},
		"number and other value": {
			in: strings.NewReader(`1 1.23 -1 -1.23 true false`),
			out: []token.Token{
				{Kind: token.KindInt, Literal: "1"},
				{Kind: token.KindFloat, Literal: "1.23"},
				{Kind: token.KindInt, Literal: "-1"},
				{Kind: token.KindFloat, Literal: "-1.23"},
				// {Kind: token.KindTrue, Literal: "true"},
				// {Kind: token.KindFalse, Literal: "false"},
				// {Kind: token.KindEOF, Literal: "\x00"},
			},
		},
	}

	for title, tt := range tests {
		t.Run(title, func(t *testing.T) {
			l := New(tt.in)
			for _, tok := range tt.out {
				next := l.GetNextToken()
				if next != tok {
					t.Fatalf("at line %d column %d got %#v expexted %#v", l.Line(), l.Column(), next, tok)
				}
			}
		})
	}
}
