package main

func main() {}

type Num struct {
	Index int
}

// Approach using a Map - O(n) Runtime
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