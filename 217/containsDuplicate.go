package main 


func containsDuplicate(nums []int) bool {

    store := make(map[int]bool)

	for i := 0; i < len(nums); i++ {

		if store[nums[i]] {
			return true
		} else {

			store[nums[i]] = true
		}

	}

	return false


    
    
}