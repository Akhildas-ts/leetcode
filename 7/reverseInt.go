package main

import (
	"fmt"
	"strconv"
)

// Example 1:

// Input: x = 123
// Output: 321
// Example 2:

// Input: x = -123
// Output: -321
// Example 3:

// Input: x = 120
// Output: 21


func reverse(x int )int {

	fmt.Println("x ",x)
   var possitive int 
  if x < 0 {
	possitive = -x 
  }else {

	possitive = x
  }

	fmt.Println("possitive",possitive)


      // convert the int to string for traverse all element in the variable 

	


	 strValue := strconv.Itoa(possitive)
 


	 
	 runes := []rune(strValue)
	 

	 if len(runes)>10 {
		fmt.Println("hello",len(runes))
		return 0
	 }
	 j := 0

	for i:=len (runes)-1;i>j;i--{


     if i == len(runes)-1 && runes[i] == 0 {
   
		continue

	 }
    runes[i],runes[j] = runes[j],runes[i]

	j++


	}
	

	strValue = string(runes)

	reverseInt, err := strconv.Atoi(strValue)

	if err != nil {
		
		fmt.Println("error from generate string to int")
		return 0
	}

	fmt.Println(reverseInt)


	
	return reverseInt


}


func main(){


reverse(240)
}