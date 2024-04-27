package main

import "fmt"

// question is we got a number as parameter
// we need to check is that multiplication  of 3
// if it was multiplication of three return true otherwise false

// algoritham

// we got a number  that number % 3 == o
// that time we now is that divied by 3 or not
// let me check these alogritham is working or not ?

func poweOfThree(n int )bool{ 

	if n % 3 == 0 {

		if n == 0 {
			return false
		}
		return true
	}

	return false
}


func main() {

	fmt.Println(poweOfThree(-1))
}