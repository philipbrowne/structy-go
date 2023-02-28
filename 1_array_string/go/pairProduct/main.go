package main

import "fmt"

func main() {
	num1, num2 := pairProduct([]int{3, 2, 5, 4, 1}, 10)
	fmt.Println(num1, num2)
}

type Num struct {
	Index int
}

//Write a function, pair_product, that takes in a list and a target product as arguments. The function should return a tuple containing a pair of indices whose elements multiply to the given target. The indices returned must be unique.

// pairProduct({3,2,5,4,1}, 8) --> 1,3
// pairProduct({3, 2, 5, 4, 1}, 10) --> 1,2
func pairProduct(numbers []int, target int) (int, int) {
	previous := map[int]*Num{}
	for i, num := range numbers {
		compliment := target / num
		if previous[compliment] != nil {
			return previous[compliment].Index, i
		}
		previous[num] = &Num{
			Index: i,
		}
	}
	return 0, 0
}