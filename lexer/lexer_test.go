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
		"complete example": {
			in:  strings.NewReader(`io_mode"async"service(http(web_proxy(listen_addr"127.0.0.1:8080"process(main(command("/usr/local/bin/awesome-app""server"))mgmt(command("/usr/local/bin/awesome-app""mgmt"))))))`),
			out: []token.Token{},
		},
	}

	for title, tt := range tests {
		t.Run(title, func(t *testing.T) {
			l := New(tt.in)
			for _, tok := range tt.out {
				next := l.GetNextToken()
				if next != tok {
					t.Fatalf("at line %d column %d got %v expexted %v", l.Line(), l.Column(), next, tok)
				}
			}
		})
	}
}
