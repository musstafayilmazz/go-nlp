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

	fmt.Println(nounchunker(mytext))
}

func nounchunker(text string) []string {
	var nounList []string
	// Tokenize
	mytokens := tokenize.NewTreebankWordTokenizer().Tokenize(text)

	//Tag
	postagger := tag.NewPerceptronTagger()

	// Append to list
	for _, token := range postagger.Tag(mytokens) {
		if token.Tag == "NN" {
			nounList = append(nounList, token.Text)
		}
	}
	return nounList
}
