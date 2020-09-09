package main

import (
	"io/ioutil"

	"github.com/mishazawa/heartache/parser"
)

func main () {
	file, err := ioutil.ReadFile("./.test/parse_me.mid")
	if err != nil {
		panic(err)
	}
	parser.ParseFile(file)
}
