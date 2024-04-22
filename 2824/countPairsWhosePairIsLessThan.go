package main

import "fmt"

func main() {

	nums := []int{-6,2,5,-2,-7,-1,3}

	fmt.Println(countPairs(nums,-2))

}

func countPairs(nums []int, target int) int {

	count := 0

	for i:= 0;i<len(nums);i++ {


		for j:=i+1 ; j<len(nums);j++ {


			if nums[i] + nums[j]  < target {
				count ++
			}
		}
	}

	return count

}