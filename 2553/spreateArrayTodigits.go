package main

import (
	"fmt"
	"strconv"
)

func main() {

	nums := []int{26, 24, 6, 63, 62}

	result := separateDigits(nums)

	fmt.Println(result)
}

func separateDigits(nums []int) []int {

	var digits []int

	for _, num := range nums {

		numString := strconv.Itoa(num)

		for _, char := range numString {

			digit, err := strconv.Atoi(string(char))

			if err != nil {

				fmt.Println("error on converting the digtis")
				return digits
			}

			digits = append(digits, digit)

		}

		
	}
	return digits

}
