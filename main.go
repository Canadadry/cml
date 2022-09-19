package main

import (
	"app/generator"
	"app/lexer"
	"app/parser"
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

func main() {
	err := run()
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func run() error {
	action := ""
	in := ""

	flag.StringVar(&action, "action", "", "can be : from_json, to_json, pretty")
	flag.StringVar(&in, "in", "", "input file")
	flag.Parse()

	if action == "" || in == "" {
		return fmt.Errorf("you must provide an action with '-action=\"\"' and an input file with '-in=\"\"'")
	}

	f, err := os.Open(in)
	if err != nil {
		return err
	}
	defer f.Close()

	switch action {
	case "from_json":
		out := map[string]interface{}{}
		err := json.NewDecoder(f).Decode(&out)
		if err != nil {
			return err
		}
		err = generator.Generate(out, os.Stdout, "", "\t")
		if err != nil {
			return err
		}
	case "to_json":
		p := parser.New(lexer.New(f))
		out, err := p.Parse()
		if err != nil {
			return err
		}
		err = json.NewEncoder(os.Stdout).Encode(out)
		if err != nil {
			return err
		}
	case "pretty":
		p := parser.New(lexer.New(f))
		out, err := p.Parse()
		if err != nil {
			return err
		}
		err = generator.Generate(out, os.Stdout, "", "\t")
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("unkown action %s valid are from_json, to_json, pretty")
	}
	return nil
}
