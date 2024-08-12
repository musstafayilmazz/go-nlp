package main

import (
	"fmt"
	"github.com/jdkato/prose/v2"
	"io/ioutil"
)

func main() {

	var mystr string = "Hello world this is golang"

	// NLP Document Struct
	docx, err := prose.NewDocument(mystr)
	if err != nil {
		panic(err)
	}

	// Tokenization

	for _, tok := range docx.Tokens() {
		fmt.Println(tok)
	}

	content, err := ioutil.ReadFile("./sample.txt")
	if err != nil {
		panic(err)
	}
	docx2, err := prose.NewDocument(string(content))
	if err != nil {
		panic(err)
	}
	for _, tok := range docx2.Tokens() {
		fmt.Println(tok.Text, tok.Label)
	}

	fmt.Println("----------------")

	for i, sent := range docx2.Sentences() {
		fmt.Println(i, sent.Text)
	}
	fmt.Println("----------------")

	for _, tok := range docx2.Tokens() {
		fmt.Println(tok.Text, tok.Tag)
	}
}
