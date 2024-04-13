package main

import (
	"fmt"
)

// quetion no :  2706


func main() {

  mrp := []int{1,2,4}
val := buyChoco(mrp,4)

fmt.Println(val)
}

// question : we have a array [ prices of choclate we have money ,]we need to must buy 2 chocalate , condtin is money cannot be negative value : 

//algoritham 
//check each element of the array is it okey for the condition




// if is it grater than money skip 

// frist chocalate 
// check the each element of the array if it sutiable for but two chocalate
// second chocalate 

// func buyChoco(prices []int, money int) int {

// l := len(prices)


// for i:= 0; i<l;i++{

// 	if money <= prices[i]{
// 		continue
// 	}

// 	for j:=0;  j<l;j++ {

// 		if i == j {
// 			continue
// 		}

// 		if prices[i] + prices[j] <= money { 

// 			fmt.Println(prices[i],prices[j])
                
// 			return money - (prices[i] + prices[j])
// 		}


// 	}
// }


// fmt.Println("cannot buy two choclate without bargaining ")
// return money

	
    
// }

func buyChoco(prices []int, money int) int {


 seen := make(map[int]bool)


 for _,choco:= range prices{

	if choco < money {

       for seenChoco:= range seen{

		if seenChoco + choco <= money{

			fmt.Println(seenChoco,choco)
			return money - (seenChoco+choco)
		}
	   }
		

		seen[choco] = true
	}


 }


 fmt.Println("cannot buy without barganing")
 return money


}