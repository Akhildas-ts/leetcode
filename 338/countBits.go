package main

import (
	"fmt"
	"strconv"
)

// question
// we got a int parameter, we need to find the 0 to n (number get the paramter) of bits in the slice

//algoritham
// when we got a number so "<< "" these opreater used to copy the bit from left and paste to right
// so when we got a number so that time we range the loop with that number , store each of byte into an slice

func main() {

	fmt.Println(countBits(5))

}
func countBits(n int) []int {

	var result []int
	sum := 0
	for i := 0; i <= n; i++ {

		by := strconv.FormatInt(int64(i), 2)
		sum = 0
		for _, val := range by {
			if val == '1' {
				sum +=1
			}

		}

		result = append(result, sum)

		fmt.Println(by)

	}

	return result

}
