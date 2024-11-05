package main

import "fmt"

func numOfStrings(patterns []string, word string) int {

	count := 0

	for _, pattWord := range patterns {

		for i := 0; i < len(pattWord); i++ {

			if pattWord[i] == word[0] {
				count++
				continue
			}
		}
	}

	return count

}

func main() {

	pattern := []string{"a", "abc", "bc", "d"}
	fmt.Println(numOfStrings(pattern, "abc"))

}
