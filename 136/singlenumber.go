package main

import "fmt"

func singleNumber(nums []int) int {

	store := make(map[int]int)

	for i := 0; i < len(nums); i++ {

		if _, ok := store[nums[i]]; !ok {

			store[nums[i]] = 1
		} else if ok {

			store[nums[i]] += 1
		}

	}

	return uniqueFindWithMap(store)

}

func uniqueFindWithMap(store map[int]int) int {

	for key, value := range store {

		if value == 1 {
			return key
		}
	}

	fmt.Println("No unqiue values are there ")
	return -1
}

func main() {

	nums := []int{2, 2, 1}

	fmt.Println(singleNumber(nums))
}
