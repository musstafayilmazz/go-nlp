package main

import (
	"fmt"
	"strings"
)

func main() {
	var mystr string = "Hello Golang NLP"
	fmt.Println(mystr)
	fmt.Printf("%T\n", mystr)

	fmt.Printf("%v\n", strings.ToUpper(mystr))
	fmt.Printf("%v\n", strings.ToLower(mystr))
	fmt.Printf("%v\n", strings.ToTitle(mystr))

	fmt.Printf("%v\n", strings.Count(mystr, "ll"))

	fmt.Printf("%v\n", strings.Split(mystr, "l"))

	tok := strings.Split(mystr, " ")
	fmt.Println(tok)

	f1 := []string{"N", "L", "P"}
	fmt.Println(strings.Join(f1, ""))
}
