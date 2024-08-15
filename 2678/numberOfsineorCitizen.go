package main

import (
	"fmt"
	"strconv"
)

func countSeniors(details []string) int {
	count := 0

	for _, sentence := range details {
		if len(sentence) >= 12 {
			// Extract age substring (assuming age is at positions 10 and 11)
			ageStr := sentence[10:12]
			fmt.Println("Extracted age string:", ageStr) // Debugging output

			age, err := strconv.Atoi(ageStr)
			if err != nil {
				fmt.Println("Error converting age:", err)
				continue
			}

			fmt.Println("Age:", age) // Debugging output

			if age > 60 {
				count++
			}
		}
	}

	fmt.Println("Number of seniors over 60:", count)
	return count
}

func main() {
	details := []string{"7868190130M7522", "5303914400F9211", "9273338290F4010"}
	countSeniors(details)
}
