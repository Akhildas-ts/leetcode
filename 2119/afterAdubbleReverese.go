package main

import (
	"fmt"
	"strconv"
)

func main() {

	fmt.Println(isSameAfterReversals(0))
}

func reverseInt(num int) int {

	revNum := strconv.Itoa(num)

	var reversed string

	if revNum[len(revNum)-1] == '0'  && len(revNum)!= 1{
		fmt.Println("last is 0 adda ")
		return -1
	}
	for i := len(revNum) - 1; i >= 0; i-- {

		reversed += string(revNum[i])

	}

	reversseInt, err := strconv.Atoi(reversed)
	

	if err != nil {

		fmt.Println("error on the return case")
		return -1
	}

	

	return reversseInt

}

func isSameAfterReversals(num int) bool{

	first := reverseInt(num)

	if first == -1{
		return false
	}
	if first != -1 {
		fmt.Println("frist :",first)
	}
	second := reverseInt(first)

	fmt.Println(second)


	return true

}
