// O(5 + 2n) = 0(n) - Linearly time complexity
func twoSum(nums []int, target int) []int {

	// 1. Convert numbers to hashmap with number and the index.
	numsMap := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		numsMap[nums[i]] = i
	}

	// 2. Loop through numbers and get the difference to reach the target to see if has already exist in numsMap.
	for i := 0; i < len(nums); i++ {

		differenceToReachTheTarget := target - nums[i]

		index, differenceExistsInNumsMap := numsMap[differenceToReachTheTarget]

		if differenceExistsInNumsMap && index != i {
			return []int{i, index}
		}

	}

	return []int{}
}
