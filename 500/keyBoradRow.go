package main

import (
	"fmt"
	"strings"
)

//question

// we got a word into the slice we need to check is the word in keyboard row,
// if is that inside the row then return word in slice .. let's start ..

// algoritham

// frist of all we need to store the key board rows into a slice of something..
//  so we get slice of word ,frist check each letter are in the row
//  each word we need to seach
// 


func main() {

  words:= []string{"akhil","das","mc"}

  fmt.Println(findWords(words))

}



func findWords(words []string) []string {
	row1 := "qwertyuiop"
	row2 := "asdfghjkl"
	row3 := "zxcvbnm"

	validWords := make([]string, 0)

	for _, word := range words {
		lowerWord := strings.ToLower(word)
		inRow1 := true
		inRow2 := true
		inRow3 := true

		for _, letter := range lowerWord {
			if !strings.Contains(row1, string(letter)) {
				inRow1 = false
			}
			if !strings.Contains(row2, string(letter)) {
				inRow2 = false
			}
			if !strings.Contains(row3, string(letter)) {
				inRow3 = false
			}
		}

		if inRow1 || inRow2 || inRow3 {
			validWords = append(validWords, word)
		}
	}

	return validWords
}