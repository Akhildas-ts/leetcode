package main

import (
	"fmt"
	"strconv"
)

func sumOfTheDigitsOfHarshadNumber(x int) int {

	if x < 10 {
		return x
	}

	arr := []int{}

	stringDigti := strconv.Itoa(x)

	for _, strDigit := range stringDigti {

		digit, _ := strconv.Atoi(string(strDigit))

		arr = append(arr, digit)
	}

	for i := 0; i < len(arr)-1; i++ {
		fmt.Println("arr[i]", arr[i], "arr[j]", arr[i+1])

		if x%(arr[i]+arr[i+1]) == 0 {
			return arr[i] + arr[i+1]
		}
	}

	return -1

}

func main() {

	fmt.Println(sumOfTheDigitsOfHarshadNumber(18))

}
