package main

import "log"

func main () {
	arr1 := []int{12,5,1,5,12,7}
	arr2 := []int{5,2,5,6,5,1,10,2,5,5}
	arr3 := []int{5,5,5,1,1,1,4}
	arr4 := []int{5,5,6,5,5,5,5}
	arr5 := []int{5, 1, 2, 5, 5, 3, 2, 5, 1, 5, 5, 5, 4, 5}

	log.Printf("Output of fiveSort on %v is %v", arr1, fiveSort(arr1))
	log.Printf("Output of fiveSort on %v is %v", arr2, fiveSort(arr2))
	log.Printf("Output of fiveSort on %v is %v", arr3, fiveSort(arr3))
	log.Printf("Output of fiveSort on %v is %v", arr4, fiveSort(arr4))
	log.Printf("Output of fiveSort on %v is %v", arr5, fiveSort(arr5))

}

// Write a function, five_sort, that takes in a list of numbers as an argument. The function should rearrange elements of the list such that all 5s appear at the end. Your function should perform this operation in-place by mutating the original list. The function should return the list.

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