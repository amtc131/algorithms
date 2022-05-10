package module

// Given an array of zise n, can you find the minimum valu in the array, implement your solution.
// see your output matches the expected output
func FindMinimum(nums []int) int {
	k := nums[0]

	for i := 0; i < len(nums); i++ {
		if nums[i] < k {
			k = nums[i]
		}
	}
	return k
}
