package main

import (
	"fmt"
	"unicode"
)

func isPalindrome(s string) bool {

	var result []rune
	for _, char := range s {

		if unicode.IsLetter(char) {

			lowerChar := unicode.ToLower(char)

			result = append(result, lowerChar)
		}
	}

	resultString := string(result)

	fmt.Println("reuslt string", resultString)

	for i, j := 0, len(resultString)-1; i < len(resultString)/2; i, j = i+1, j-1 {

		fmt.Println("i", string(resultString[i]), "j", string(resultString[j]))

		if resultString[i] != resultString[j] {
			return false
		}

	}

	return true
}

func main() {

	fmt.Println(isPalindrome("A man, a plan, a canal: Panamae"))

}
