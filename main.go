package main

import (
	"fmt"
	"io/ioutil"

	"github.com/mishazawa/heartache/parser"
)

func main () {
	file, err := ioutil.ReadFile("./.test/A Sacred Lot.mid")
	if err != nil {
		panic(err)
	}
	mfile, _ := parser.ParseFile(file)
	fmt.Printf("%+v\n", mfile)
}
