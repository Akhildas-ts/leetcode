package main

import (
	"fmt"
	"strconv"
)

// if  3  divisbile by n "fizz"
// if 5 divisibile by n "buzz"
// if 3 and 5 divisible by n "fizzBuzz"

func main() {

	fmt.Println(fizzBuzz(15))

}

func fizzBuzz(n int) []string {

	outPut := make([]string, 0)

	for i := 1; i <= n; i++ {

		if i%3  == 0 && i%5==0 {
			outPut = append(outPut, "FizzBuzz")
		} else if i%5 == 0 {
			outPut = append(outPut, "Buzz")
		} else if i%3 == 0{
			outPut = append(outPut, "Fizz")
		} else {
			outPut = append(outPut, strconv.Itoa(i))

		}

	}

	return outPut

}
