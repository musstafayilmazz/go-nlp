package main

import (
	"fmt"
	"github.com/JesusIslam/tldr"
	"io/ioutil"
)

func main() {

	content, err := ioutil.ReadFile("sample.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(content))

	summary := tldr.New()
	results, _ := summary.Summarize(string(content), 3)

	fmt.Println("Summarized Text:")
	fmt.Println(results)

}
