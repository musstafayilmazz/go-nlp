package main

import (
	"fmt"
	"github.com/jdkato/prose/tag"
	"github.com/jdkato/prose/tokenize"
)

func main() {

	mytext := "I am going to fish a fish at the lake"

	// Tokenize
	mytokens := tokenize.NewTreebankWordTokenizer().Tokenize(mytext)
	// Tag/Annotate
	postagger := tag.NewPerceptronTagger()
	for _, token := range postagger.Tag(mytokens) {
		fmt.Println(token.Text, token.Tag)
	}
}
