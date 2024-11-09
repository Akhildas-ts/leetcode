package main

import "fmt"

func rotateString(s string, goal string) bool {

	runes := []rune(s)

	for i := 0; i < len(runes); i++ {
		ex := runes[0]
		for j := 0; j < len(runes)-1; j++ {

			runes[j] = runes[j+1]
		}
		runes[len(runes)-1] = ex

		if string(runes) == goal {
			return true
		}

	}

	return false

}

func main() {

	fmt.Println(rotateString("abced", "cdeab"))

}
