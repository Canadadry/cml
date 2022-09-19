package generator

import (
	"bytes"
	"testing"
)

func TestGenerate(t *testing.T) {
	tests := map[string]struct {
		in     map[string]interface{}
		prefix string
		indent string
		out    string
	}{
		"one key value": {
			in:  map[string]interface{}{"key": "value"},
			out: `key "value"`,
		},
		"one key value with quoted quote": {
			in:  map[string]interface{}{"key": "va\"lue"},
			out: `key "va\"lue"`,
		},
	}

	for title, tt := range tests {
		t.Run(title, func(t *testing.T) {
			buf := bytes.Buffer{}
			err := Generate(tt.in, &buf, tt.prefix, tt.indent)
			if err != nil {
				t.Fatalf("failed %v", err)
			}
			if buf.String() != tt.out {
				t.Fatalf("exp\n%s\n%v\ngot\n%s\n%v", tt.out, []byte(tt.out), buf.String(), buf.Bytes())
			}
		})
	}
}
