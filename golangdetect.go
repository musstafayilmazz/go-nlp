package main

import (
	"fmt"
	"github.com/abadojack/whatlanggo"
	"github.com/rylans/getlang"
)

func main() {

	var mystr = "Hello World"
	var mystr2 = "Ankara, Türkiye'nin başkentidir"
	fmt.Println(mystr)

	lang := getlang.FromString(mystr)
	lang2 := getlang.FromString(mystr2)
	fmt.Println(lang.LanguageCode(), lang2.Tag())
	fmt.Println(lang2.LanguageCode(), lang2.Confidence())

	lang3 := whatlanggo.DetectLang(mystr2)
	fmt.Println(lang3.String(), lang3.Iso6391(), lang3.Iso6393())
	lang4 := whatlanggo.Detect(mystr2)
	fmt.Println(lang4.Lang, lang4.Confidence)
	fmt.Println(whatlanggo.Scripts[lang4.Script])
}
