package main

import (
	"fmt"
	"os"
)

func main () {
	fmt.Printf("Most frequent char in %s is: %s\n", os.Args[1], mostFrequent(os.Args[1]))
}

// mostFrequent("bookeeper") -> "e"
// mostFrequent("abby") -> "b"
func mostFrequent(s string) string {
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