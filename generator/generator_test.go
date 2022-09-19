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
			in: map[string]interface{}{"key": "value"},
			out: `key "value"
`,
		},
		"one key value with quoted quote": {
			in: map[string]interface{}{"key": "va\"lue"},
			out: `key "va\"lue"
`,
		},
		"multi key value with scalar value": {
			in: map[string]interface{}{
				"key_str":   "value",
				"key_int":   -13,
				"key_float": 12.34,
				"key_true":  true,
				"key_false": false,
			},
			out: `key_str "value"
key_int -13
key_float 12.34
key_true true
key_false false
`,
		},
		"sub map": {
			in: map[string]interface{}{
				"key_map": map[string]interface{}{
					"key_int":   -13,
					"key_float": 12.34,
					"key_true":  true,
					"key_false": false,
				},
			},
			indent: "\t",
			out: `key_map (
	key_int -13
	key_float 12.34
	key_true true
	key_false false
)
`,
		},
		"sub array": {
			in: map[string]interface{}{
				"key_array": []interface{}{
					-13,
					12.34,
					true,
					false,
				},
			},
			indent: "\t",
			out: `key_array (
	-13
	12.34
	true
	false
)
`,
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
