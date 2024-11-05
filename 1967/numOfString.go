package main

import "fmt"

func numOfStrings(patterns []string, word string) int {

	count := 0

	// Loop through each pattern in the patterns slice
	for _, pattWord := range patterns {
		// Check if the current pattern exists in the word
		if isSubstring(pattWord, word) {
			count++
		}
	}

	return count
}

// Function to manually check if pattern is a substring of word
func isSubstring(pattern, word string) bool {
	// We iterate through the word to check if the pattern matches at any position
	for i := 0; i <= len(word)-len(pattern); i++ {
		match := true
		for j := 0; j < len(pattern); j++ {
			// If any character doesn't match, mark it as false
			if word[i+j] != pattern[j] {
				match = false
				break
			}
		}
		// If we find a complete match, return true
		if match {
			return true
		}
	}
	return false
}

func main() {

	pattern := []string{"a", "abc", "bc", "d"}
	fmt.Println(numOfStrings(pattern, "abcd"))

}
