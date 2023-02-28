package main

import (
	"fmt"
	"os"
)

func main () {
	isAnagram := anagram(os.Args[1], os.Args[2])
	fmt.Println("Is Anagram: ", isAnagram)
}

func anagram(s string, t string) bool {
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