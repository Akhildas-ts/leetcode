package main

import (
	"fmt"
)

// question
// we need to find each of the letter asci number then sutract

//algoritham
//  we get a string so convert each letter into aski
// then create a two loop then (i,j) j need to go n and i need to n-1,
// here we need to subtract (i-j) +
// then we got the output

func main() {

	fmt.Println(scoreOfString("hello"))
}

func scoreOfString(s string) int {


	sum := 0
	val := 0
	for i := 0; i < len(s); i++ {


	

		for j := i + 1; j <len(s); j++ {

             fmt.Println("s[i]-",s[i] ,"s[j]",s[j])
			 
			 val = int(s[i]) - int(s[j])
			 if val < 0 {
				val = -val
			 }
			 fmt.Println(val)
			 sum += val
			 break
			 

		}
	}

	return sum
}
