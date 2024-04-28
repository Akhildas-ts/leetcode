package main

import "fmt"

func main() {

	fmt.Println(checkPerfectNumber(7))
}

func checkPerfectNumber(num int) bool {

	sum := 0
	for i := 1; i <= num; i++ {

		if num%i == 0 {

			sum += i

			if sum == num {
				return true
			}

		}
	}
	return false

}