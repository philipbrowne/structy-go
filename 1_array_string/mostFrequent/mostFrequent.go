package main

func main () {
}

// Approach using Map - O(n) Runtime
func mostFrequentChar(s string) string {
	count := map[rune]int{}
	for _, char := range s {
		count[char]++
	}
	var best rune
	for _, char := range s {
		if best == 0 || count[char] > count[best] {
			best = char
		}
	}
	return string(best)
}