package main

import "fmt"

func main() {

	hours := []int{72, 48,  24, 34}

	fmt.Println(countCompleteDayPairs(hours))
}

func countCompleteDayPairs(hours []int) int {

	count := 0
	for i,_ := range hours {
b
		for j := 0; j < i; j++ {

			checkIsDay := hours[i] + hours[j]

			if checkIsDay%24 == 0 {

				count++

			}

		}
	}

	return count
}
