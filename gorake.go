package main

import (
	"fmt"
	rake "github.com/Obaied/RAKE.Go"
	"io/ioutil"
	"log"
)

func main() {
	content, err := ioutil.ReadFile("./sample.txt")
	if err != nil {
		log.Fatal(err)
	}

	keywordMap := make(map[string]float64)
	candidates := rake.RunRake(string(content))

	for _, candidate := range candidates {
		keywordMap[candidate.Key] = candidate.Value
	}
	fmt.Println(keywordMap)
}
