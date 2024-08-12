package main

import (
	"fmt"
	"github.com/jdkato/prose/v2"
	"log"
)

func main() {
	// NER
	// Entity: Person/People/Org/Location/etc

	var mytext string = "Mustafa Yılmaz works in London as golang developer. He is 28 years old"

	// NLP Document Struct/Object

	doc, err := prose.NewDocument(mytext)
	if err != nil {
		log.Fatal(err)
	}
	for index, token := range doc.Entities() {
		fmt.Println(index, token.Text, token.Label)
	}

	for _, entity := range doc.Entities() {
		fmt.Println(entity.Text, entity.Label)
	}

	for index, tokens := range doc.Tokens() {
		fmt.Println(index, tokens.Text, tokens.Label)
	}
}
