package main

import (
	"fmt"
	"strconv"
)

func main() {

	// isPalindrome(101)
	this := &mystack{}

	this.pop(10)
}

type mystack struct {
	queue1 []int
	queue2 []int
}

func isPalindrome(x int) bool{

	// conver the number to string
	xstr := strconv.Itoa(x)

    var reverseString string

	for i:= len(xstr)-1 ; i>=0; i--{
         
		reverseString += string(xstr[i])


	}

	if xstr == reverseString {
		return true
	}

	
	return false

}

func (this *mystack) pop(x int) {

	xstr := strconv.Itoa(x) 

	for _, digits := range xstr {

		digitsInt, _ := strconv.Atoi(string(digits))

		this.queue2 = append(this.queue2, digitsInt)  

	}

	for len(this.queue1) > 0 {

		this.queue2 = append(this.queue2, this.queue1[0])

		this.queue1 = this.queue1[1:]

	}

	v := this.queue1

	fmt.Println(v)

}
