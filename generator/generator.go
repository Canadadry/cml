package generator

import (
	"fmt"
	"io"
	"strings"
)

func Generate(in map[string]interface{}, w io.Writer, prefix, indent string) error {
	for key, value := range in {
		switch value.(type) {
		case map[string]interface{}:
			err := Generate(value.(map[string]interface{}), w, prefix, prefix+indent)
			if err != nil {
				return err
			}
		case []interface{}:
			err := generateSlice(value.([]interface{}), w, prefix, prefix+indent)
			if err != nil {
				return err
			}
		case string:
			str := value.(string)
			str = strings.ReplaceAll(str, "\"", "\\\"")
			fmt.Fprintf(w, "%s%s%s \"%s\"", prefix, indent, key, str)
		default:
			fmt.Fprintf(w, "%s%s%s \"%v\"", prefix, indent, key, value)
		}
	}
	return nil
}

func generateSlice(in []interface{}, w io.Writer, prefix, indent string) error {
	return nil
}
