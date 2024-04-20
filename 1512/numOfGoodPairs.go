package main

import "fmt"

func numIdenticalPairs(nums []int) int {

	count := 0

	for i := 0; i < len(nums); i++ {

		for j := 0; j < len(nums); j++ {
			if j > i && nums[i] == nums[j] {
				count++

				fmt.Println(count)
			}
		}

	}
	return count

}
