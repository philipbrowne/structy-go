package main

func main () {}

// Approach using Two Pointers - O(n) Runtime
func fiveSort(nums []int) []int {
	i := 0
	j := len(nums) - 1
	for i < j {
		if nums[j] == 5 {
			j --
		} else if nums[i] == 5 {
			nums[j], nums[i] = nums[i], nums[j]
			j --
		} else {
			i ++
		}
	}
	return nums
}