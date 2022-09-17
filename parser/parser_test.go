package parser

import (
	"app/lexer"
	"reflect"
	"strings"
	"testing"
)

func TestParser(t *testing.T) {
	tests := map[string]struct {
		in  string
		out map[string]interface{}
	}{
		"minified sample example": {
			in: `io_mode"async"service(http(web_proxy(listen_addr"127.0.0.1:8080"process(main(command("/usr/local/bin/awesome-srv""server"))mgmt(command("/usr/local/bin/awesome-mgmt""mgmt"))))))`,
			out: map[string]interface{}{
				"io_mode": "async",
				"service": map[string]interface{}{
					"http": map[string]interface{}{
						"web_proxy": map[string]interface{}{
							"listen_addr": "127.0.0.1:8080",
							"process": map[string]interface{}{
								"main": map[string]interface{}{
									"command": []interface{}{
										"/usr/local/bin/awesome-srv",
										"server",
									},
								},
								"mgmt": map[string]interface{}{
									"command": []interface{}{
										"/usr/local/bin/awesome-mgmt",
										"mgmt",
									},
								},
							},
						},
					},
				},
			},
		},
	}

	for title, tt := range tests {
		t.Run(title, func(t *testing.T) {
			l := lexer.New(strings.NewReader(tt.in))
			result, err := New(l).Parse()
			if err != nil {
				t.Fatalf("failed %v", err)
			}
			if !reflect.DeepEqual(result, tt.out) {
				t.Fatalf("exp %#v\ngot %#v", tt.out, result)
			}
		})
	}
}
