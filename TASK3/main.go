package main

import (
	"fmt"
	"strings"
)

func findFirstStringInBracket(str string) string {
	if len(str) > 0 {
		indexFirstBracketFound := strings.Index(str, "(")
		if indexFirstBracketFound >= 0 && indexFirstBracketFound != len(str)-1 {
			runes := []rune(str)
			wordsAfterFirstBracket := string(runes[indexFirstBracketFound:len(str)])
			indexClosingBracketFound := strings.Index(wordsAfterFirstBracket, ")")

			if indexClosingBracketFound >= 0 {
				runes := []rune(wordsAfterFirstBracket)
				return string(runes[1:indexClosingBracketFound])
			}
		}
	}

	return ""
}

func main() {
	var res string
	res = findFirstStringInBracket("Hello World(")
	fmt.Println(res)

	res = findFirstStringInBracket("Hello (World)")
	fmt.Println(res)

	res = findFirstStringInBracket("Hello World()")
	fmt.Println(res)

	res = findFirstStringInBracket("Hell((o W)orld(")
	fmt.Println(res)
}
