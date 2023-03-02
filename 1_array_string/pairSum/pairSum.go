package main

import "fmt"

func main() {
	pair := pairSum([]int{3,2,5,4,1}, 8)
	fmt.Println("Pairsum for {3,2,5,4,1} with a sum of 8", pair)
}

type Num struct {
	Index int
}

func pairSum(numbers []int, target int) []int {
	previous := map[int]*Num{}
	for i, num := range numbers {
		compliment := target - num
		if previous[compliment] != nil {
			return []int{previous[compliment].Index, i}
		}
		previous[num] = &Num{
			Index: i,
		}
	}
	return []int{0, 0}
}