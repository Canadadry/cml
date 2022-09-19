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
			fmt.Fprintf(w, "%s%s (\n", prefix, key)
			err := Generate(value.(map[string]interface{}), w, prefix+indent, prefix)
			if err != nil {
				return err
			}
			fmt.Fprintf(w, "%s)\n", prefix)
		case []interface{}:
			fmt.Fprintf(w, "%s%s (\n", prefix, key)
			err := generateSlice(value.([]interface{}), w, prefix+indent, prefix)
			if err != nil {
				return err
			}
			fmt.Fprintf(w, "%s)\n", prefix)
		case string:
			str := value.(string)
			str = strings.ReplaceAll(str, "\"", "\\\"")
			fmt.Fprintf(w, "%s%s \"%s\"\n", prefix, key, str)
		default:
			fmt.Fprintf(w, "%s%s %v\n", prefix, key, value)
		}
	}
	return nil
}

func generateSlice(in []interface{}, w io.Writer, prefix, indent string) error {
	for _, value := range in {
		switch value.(type) {
		case map[string]interface{}:
			fmt.Fprintf(w, "%s (\n", prefix)
			err := Generate(value.(map[string]interface{}), w, prefix+indent, prefix)
			if err != nil {
				return err
			}
			fmt.Fprintf(w, "%s)\n", prefix)
		case []interface{}:
			fmt.Fprintf(w, "%s (\n", prefix)
			err := generateSlice(value.([]interface{}), w, prefix+indent, prefix)
			if err != nil {
				return err
			}
			fmt.Fprintf(w, "%s)\n", prefix)
		case string:
			str := value.(string)
			str = strings.ReplaceAll(str, "\"", "\\\"")
			fmt.Fprintf(w, "%s\"%s\"\n", prefix, str)
		default:
			fmt.Fprintf(w, "%s%v\n", prefix, value)
		}
	}
	return nil
}
