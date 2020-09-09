package main

import (
	"io/ioutil"

	"github.com/mishazawa/heartache/parser"
)

func main () {
	file, err := ioutil.ReadFile("./.test/Twinkle.mid")
	if err != nil {
		panic(err)
	}
	parser.ParseFile(file)
}
