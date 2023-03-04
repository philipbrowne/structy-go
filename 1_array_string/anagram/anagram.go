package main

func main () {
}

// Approach using Map - O(n+m) Runtime where n and m are the lengths of each string
func anagrams(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	count := map[rune]int{}
	for _, char := range s {
		count[char]++
	}
	for _, char := range t {
		if count[char] <= 0 {
			return false
		}
		count[char]--
	}
	return true
}