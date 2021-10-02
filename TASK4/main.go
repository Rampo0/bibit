package main

import "fmt"

func main() {
	input := []string{"kita", "atik", "tika", "aku", "kia", "makan", "kua"}
	anag_map := make(map[int][]string)

	for _, word := range input {
		count := 0
		for _, ch := range word {
			count = count + int(ch)
		}

		anag_map[count] = append(anag_map[count], word)
	}

	for _, val := range anag_map {
		fmt.Println(val)
	}
}
