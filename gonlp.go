package main

import (
	"fmt"
	"github.com/jdkato/prose/tokenize"
	"regexp"
	"strings"
)

func main() {

	var mytext string = "Mustafa is learning NLP with Golang"

	// Tokenization

	// Method 1: Split
	tokens := strings.Split(mytext, " ")
	fmt.Println("Text", mytext)
	fmt.Println("Tokens", tokens)

	// Method 2: Rule Based
	r := regexp.MustCompile(`\w+`)
	newtokens := r.FindAllString(mytext, -1)
	fmt.Println(newtokens)

	// Method 3: Regex + Split
	r2 := regexp.MustCompile(" ")
	newtokens2 := r2.Split(mytext, -1)
	fmt.Println(newtokens2)

	// Method 4: Packages
	tokenizer := tokenize.NewTreebankWordTokenizer()
	for _, tok := range tokenizer.Tokenize(mytext) {
		fmt.Println(tok)
	}

	wordtokens := tokenizer.Tokenize(mytext)
	fmt.Println(wordtokens)
}
