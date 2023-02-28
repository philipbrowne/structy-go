package main

import "fmt"

func main() {
	num1, num2 := (pairSum([]int{3,2,5,4,1}, 8))
	fmt.Println(num1, num2)
}

type Num struct {
	Index int
}

// Write a function, pair_sum, that takes in a list and a target sum as arguments. The function should return a tuple containing a pair of indices whose elements sum to the given target. The indices returned must be unique.

// pairSum({3,2,5,4,1}, 8) -> (0,2)
func pairSum(numbers []int, target int) (int, int) {
	previous := map[int]*Num{}
	for i, num := range numbers {
		compliment := target - num
		if previous[compliment] != nil {
			return previous[compliment].Index, i
		}
		previous[num] = &Num{
			Index: i,
		}
	}
	return 0, 0
}