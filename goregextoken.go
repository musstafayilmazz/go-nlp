package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	var mytext string = "Mustafa love golang and all of it's frameworks"

	// Split Based
	tokens := strings.Split(mytext, " ")
	fmt.Println(tokens)

	// Rule Based
	r := regexp.MustCompile(`\w+`)
	mytokens := r.FindAllString(mytext, -1)
	fmt.Println(mytokens)
}
