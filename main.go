package main

import (
	"fmt"
	"io/ioutil"

	"github.com/mishazawa/heartache/parser"
)

func main () {
	file, err := ioutil.ReadFile("./.test/short.mid")
	if err != nil {
		panic(err)
	}
	mfile, _ := parser.ParseFile(file)
	for _, ev := range mfile.Tracks[0].Events {
		fmt.Printf("%+v\n", ev)
	}
}
