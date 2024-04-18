package main

func removeElement(nums []int, val int) int {

	var index int

	for _, num := range nums {

		if num != val {
			nums[index] = num
			index++
		}

	}

	return index
}
