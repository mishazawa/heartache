package main

import (
	"github.com/mishazawa/heartache/parser"
)

func main () {
	_, err := parser.ParseFile("./.test/parse_me.mid")
	if err != nil {
		panic(err)
	}
}
