package main

import "fmt"

func main() {
	var char byte = 'A'
	var char2 rune = 'A'

	fmt.Println(char, char2)

	charA := byte('A')
	charB := byte('B')

	fmt.Println(charA, charB)

	str1 := string(charA)
	str2 := string(charB)
	fmt.Println(str1, str2)

	fmt.Printf("%c %c\n", charA, charB)
}
