package main

import (
	"fmt"
	cregex "github.com/mingrammer/commonregex"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
)

func main() {
	var mystr string = "Hello Golang Developers! my phone number is 055555555 and my mail is xyz@email.com"

	content, err := ioutil.ReadFile("./sample.txt")
	if err != nil {
		log.Fatal(err)
	}

	docx2 := string(content)

	//	Text Preprocessing
	//	Normalizing : case, unicode
	//	Remove noise
	//	Lemma/Stemming
	//  Tokenization

	fmt.Println(strings.ToLower(mystr))
	fmt.Println(strings.ToLower(docx2))

	// Extract Emails with standard library Regexp
	p := regexp.MustCompile(`Golang`)
	fmt.Println(p.FindAllString(mystr, -1))

	fmt.Println(p.ReplaceAllString(mystr, "Rust"))

	fmt.Println(cregex.Emails(docx2))
	fmt.Println(cregex.Phones(docx2))
	fmt.Println(cregex.Links(docx2))

	patternMail := regexp.MustCompile(cregex.EmailPattern)
	patternLink := regexp.MustCompile(cregex.LinkPattern)
	fmt.Println(patternMail.FindAllString(docx2, -1))

	fmt.Println(patternMail.ReplaceAllString(docx2, "---HIDDEN---"))
	fmt.Println(patternLink.ReplaceAllString(docx2, "---FORBIDDEN---"))

}
