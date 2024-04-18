package  main 


func twoSum(nums []int, target int) []int {

    store := make(map[int]int)

	result := make([]int, 0)

	for i := 0; i < len(nums); i++ {

		if store != nil {

			for j, val := range store {
				if val+nums[i] == target {
					result = append(result, i)
					result = append(result, j)

					return result
				}
			}
		}

		store[i] = nums[i]

	}
	return result
    
}