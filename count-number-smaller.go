// First and easiest solution - Time complexity equals quadratic O(n²)
func smallerNumbersThanCurrent(nums []int) []int {
    counter := make([]int, len(nums))

    for i := 0; i < len(nums); i++ {
        for j := 0; j < len(nums); j++ {
            if i != j && nums[i] > nums[j] {
                counter[i]++
            }
        }
    }

    return counter
}
