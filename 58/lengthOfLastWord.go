package main

import "strings"

func lengthOfLastWord(s string) int {
	trimmed := strings.TrimSpace(s)

	words := strings.Split(trimmed, " ")

	lastWord := words[len(words)-1]
	return len(lastWord)

}
