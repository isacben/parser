package main

import (
	"encoding/json"
	"fmt"
	"parser/pkg/lexer"
	"parser/pkg/parser"
)

// From: https://medium.com/@bradford_hamilton/building-a-json-parser-and-query-tool-with-go-8790beee239a
// Repo https://github.com/bradford-hamilton/dora

func main() {

	var s = `{
		"test": { "num": 1 },
		"name": "isaac"
	}`

	l := lexer.New(s)
	p := parser.New(l)
	tree, err := p.ParseProgram()
	if err != nil {
		fmt.Printf("error: parser error: %v\n", err)
	}

	fmt.Printf("%+v\n", *tree.RootValue)

	jtree, _ := json.MarshalIndent(*tree.RootValue, "  ", "    ")
	fmt.Println(string(jtree))
}
