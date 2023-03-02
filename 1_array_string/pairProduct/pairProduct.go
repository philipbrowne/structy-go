package main

import "fmt"

func main() {
	pair := pairProduct([]int{3, 2, 5, 4, 1}, 10)
	fmt.Println("Pair Product for {3,2,5,4,1} for a product of 10", pair)
}

type Num struct {
	Index int
}

func pairProduct(numbers []int, target int) []int {
	previous := map[int]*Num{}
	for i, num := range numbers {
		compliment := target / num
		if previous[compliment] != nil {
			return []int{previous[compliment].Index, i}
		}
		previous[num] = &Num{
			Index: i,
		}
	}
	return []int{0, 0}

}